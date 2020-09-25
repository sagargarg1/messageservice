package handlers

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/sagargarg1/messageservice/pkg/data"
	"github.com/sagargarg1/messageservice/pkg/utils"
	"github.com/sagargarg1/messageservice/pkg/model"
)

type MessageHandlerInterface interface {
        AddRoutes(router *mux.Router)
	AddMessage(rw http.ResponseWriter, r *http.Request)
        UpdateMessage(rw http.ResponseWriter, r *http.Request)
        GetAllMessages(rw http.ResponseWriter, r *http.Request)
        GetMessage(rw http.ResponseWriter, r *http.Request)
        DeleteMessage(rw http.ResponseWriter, r *http.Request)
}


type MessageHandler struct {
	Logging	hclog.Logger
        DBInterface data.MessageDBInterface         
}

func NewMessageHandler(Logging hclog.Logger, DBInterface data.MessageDBInterface) MessageHandlerInterface {
	return &MessageHandler{
		Logging:        Logging,
		DBInterface: 	DBInterface,
	}
}

func (m *MessageHandler) AddRoutes(router *mux.Router) {
	routes := router.PathPrefix("").Subrouter()
	routes.HandleFunc("", m.AddMessage).Methods(http.MethodPost)
        routes.HandleFunc("", m.UpdateMessage).Methods(http.MethodPut)
        routes.HandleFunc("/{id}", m.GetMessage).Methods(http.MethodGet)
        routes.HandleFunc("/{id}", m.DeleteMessage).Methods(http.MethodDelete)
        routes.HandleFunc("/all", m.GetAllMessages).Methods(http.MethodGet)
}

// responses:
//	201: Created
//	400: BadRequest
//	500: InternalError
func (m *MessageHandler) AddMessage(rw http.ResponseWriter, r *http.Request) {

	messageBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		m.Logging.Error("Invalid request body", "error", err)
		rw.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericError{Message: "Invalid request body"}, rw)
		return
	}
	defer r.Body.Close()

	var message model.Message
	if err := json.Unmarshal(messageBody, &message); err != nil {
		m.Logging.Error("Invalid json body", "error", err)
		rw.WriteHeader(http.StatusBadRequest)
                utils.ToJSON(&model.GenericError{Message: "Invalid json body"}, rw)
                return
	}
	m.Logging.Debug("Adding message: %#v\n", message)

        id := m.DBInterface.AddMessage(message)
	err = utils.ToJSON(id, rw)
        if err != nil {
		m.Logging.Error("Failed to encode response", "error", err)
                rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        }
	rw.WriteHeader(http.StatusCreated)
	m.Logging.Debug("Successfully added message: %#v\n", message)
}

// responses:
//      200: Success
//      400: BadRequest
//	404: NotFound
func (m *MessageHandler) UpdateMessage(rw http.ResponseWriter, r *http.Request) {

	messageBody, err := ioutil.ReadAll(r.Body)
        if err != nil {
		m.Logging.Error("Invalid request body", "error", err)
                rw.WriteHeader(http.StatusBadRequest)
                utils.ToJSON(&model.GenericError{Message: "Invalid request body"}, rw)
                return
        }
        defer r.Body.Close()

        var message model.Message
        if err := json.Unmarshal(messageBody, &message); err != nil {
		m.Logging.Error("Invalid json body", "error", err)
                rw.WriteHeader(http.StatusBadRequest)
                utils.ToJSON(&model.GenericError{Message: "Invalid json body"}, rw)
                return
        }

	m.Logging.Debug("Updating message: %#v\n", message)
        error := m.DBInterface.UpdateMessage(message)
        if error == model.ErrMessageNotFound {
		m.Logging.Error("Message ID not found", "error", error)
                rw.WriteHeader(http.StatusNotFound)
                utils.ToJSON(&model.GenericError{Message: "Message not found in database"}, rw)
                return
        }

        rw.WriteHeader(http.StatusOK)
	m.Logging.Debug("Successfully updates message: %#v\n", message)
}

// responses:
//      200: Success
//      404: NotFound
//      500: InternalError
func (m *MessageHandler) GetAllMessages(rw http.ResponseWriter, r *http.Request) {

	m.Logging.Debug("Get all messages \n")
        messages := m.DBInterface.GetAllMessages()

        err := utils.ToJSON(messages, rw)
        if err != nil {
		m.Logging.Error("Failed to encode response", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        }
	rw.WriteHeader(http.StatusOK)
        m.Logging.Debug("Successfully retrieved messages \n")
}

// responses:
//      200: Success
//      404: NotFound
//      500: InternalError
func (m *MessageHandler) GetMessage(rw http.ResponseWriter, r *http.Request) {

        id := getMessageID(r)

	m.Logging.Debug("Get message with ID: %#d\n", id)
        message, err := m.DBInterface.GetMessage(id)

        switch err {
        case nil:

        case model.ErrMessageNotFound:
		m.Logging.Error("Failed to find message with given id", "error", err)
                rw.WriteHeader(http.StatusNotFound)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        default:
		m.Logging.Error("Failed to retrieve message", "error", err)
                rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        }

        err = utils.ToJSON(message, rw)
        if err != nil {
		m.Logging.Error("Failed to encode response", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        }
	rw.WriteHeader(http.StatusOK)
	m.Logging.Debug("Successfully retrieved message with ID: %#d\n", id)	
}

// responses:
//      204: NoContent
//      404: NotFound
//      500: InternalError
func (m *MessageHandler) DeleteMessage(rw http.ResponseWriter, r *http.Request) {

        id := getMessageID(r)

	m.Logging.Debug("Delete message with ID: %#d\n", id)

        err := m.DBInterface.DeleteMessage(id)
	
	switch err {
	case nil:

	case model.ErrMessageNotFound:
		m.Logging.Error("Failed to find message with given id", "error", err)
                rw.WriteHeader(http.StatusNotFound)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return

	default:
		m.Logging.Error("Failed to delete message with given id", "error", err)
                rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
	}

        rw.WriteHeader(http.StatusNoContent)
	m.Logging.Debug("Successfully deleted message with ID: %#d\n", id)
}

func getMessageID(r *http.Request) int {
        // parse the product id from the url
        vars := mux.Vars(r)

        // convert the id into an integer and return
        id, err := strconv.Atoi(vars["id"])
        if err != nil {
                // should never happen
                panic(err)
        }

        return id
}
