package handlers

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sagargarg1/messageservice/pkg/data"
	"github.com/sagargarg1/messageservice/pkg/utils"
	"github.com/sagargarg1/messageservice/pkg/model"
)

var (
	HandlerInterface MessageHandlerInterface = &MessagesHandler{}
)

type MessagesHandler struct {}

type MessageHandlerInterface interface {
	AddMessage(rw http.ResponseWriter, r *http.Request)
	UpdateMessage(rw http.ResponseWriter, r *http.Request)
	GetAllMessages(rw http.ResponseWriter, r *http.Request)
	GetMessage(rw http.ResponseWriter, r *http.Request)
	DeleteMessage(rw http.ResponseWriter, r *http.Request)
}

func (m *MessagesHandler) AddMessage(rw http.ResponseWriter, r *http.Request) {

	messageBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		utils.ToJSON(&model.GenericError{Message: "Invalid request body"}, rw)
		return
	}
	defer r.Body.Close()

	var message model.Message
	if err := json.Unmarshal(messageBody, &message); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
                utils.ToJSON(&model.GenericError{Message: "Invalid json body"}, rw)
                return
	}

        data.MessageInterface.AddMessage(message)
}

func (m *MessagesHandler) UpdateMessage(rw http.ResponseWriter, r *http.Request) {

	messageBody, err := ioutil.ReadAll(r.Body)
        if err != nil {
                rw.WriteHeader(http.StatusBadRequest)
                utils.ToJSON(&model.GenericError{Message: "Invalid request body"}, rw)
                return
        }
        defer r.Body.Close()

        var message model.Message
        if err := json.Unmarshal(messageBody, &message); err != nil {
                rw.WriteHeader(http.StatusBadRequest)
                utils.ToJSON(&model.GenericError{Message: "Invalid json body"}, rw)
                return
        }

        error := data.MessageInterface.UpdateMessage(message)
        if error == model.ErrMessageNotFound {

                rw.WriteHeader(http.StatusNotFound)
                utils.ToJSON(&model.GenericError{Message: "Message not found in database"}, rw)
                return
        }

        rw.WriteHeader(http.StatusNoContent)
}

func (m *MessagesHandler) GetAllMessages(rw http.ResponseWriter, r *http.Request) {

        messages := data.MessageInterface.GetAllMessages()

        err := utils.ToJSON(messages, rw)
        if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        }
}

func (m *MessagesHandler) GetMessage(rw http.ResponseWriter, r *http.Request) {

        id := getProductID(r)

        message, err := data.MessageInterface.GetMessage(id)

        switch err {
        case nil:

        case model.ErrMessageNotFound:
                rw.WriteHeader(http.StatusNotFound)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        default:
                rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        }

        err = utils.ToJSON(message, rw)
        if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        }
}

func (m *MessagesHandler) DeleteMessage(rw http.ResponseWriter, r *http.Request) {

        id := getProductID(r)

        err := data.MessageInterface.DeleteMessage(id)
        if err == model.ErrMessageNotFound {
                rw.WriteHeader(http.StatusNotFound)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        }

        if err != nil {
                rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        }

        rw.WriteHeader(http.StatusNoContent)
}

func getProductID(r *http.Request) int {
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
