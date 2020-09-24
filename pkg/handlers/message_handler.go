package handlers

import (
	"net/http"
	"encoding/json"
        "fmt"
	"strings"

	"github.com/gorilla/mux"
	"github.com/saggarg/messageservice/pkg/data"
	"github.com/saggarg/messageservice/pkg/utils"
	"github.com/saggarg/messageservice/pkg/model"
)

type MessagesHandler struct {
}

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
		utils.ToJSON(&GenericError{Message: "Invalid request body"}, rw)
		return
	}
	defer r.Body.Close()

	var message model.Message
	if err := json.Unmarshal(messageBody, &message); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
                utils.ToJSON(&GenericError{Message: "Invalid json body"}, rw)
                return
	}

        data.messageDB.AddMessage(message)
}

func (m *MessagesHandler) UpdateMessage(rw http.ResponseWriter, r *http.Request) {

	messageBody, err := ioutil.ReadAll(r.Body)
        if err != nil {
                rw.WriteHeader(http.StatusBadRequest)
                utils.ToJSON(&GenericError{Message: "Invalid request body"}, rw)
                return
        }
        defer r.Body.Close()

        var message model.Message
        if err := json.Unmarshal(messageBody, &message); err != nil {
                rw.WriteHeader(http.StatusBadRequest)
                utils.ToJSON(&GenericError{Message: "Invalid json body"}, rw)
                return
        }

        err := data.messageDB.UpdateMessage(message)
        if err == model.ErrProductNotFound {

                rw.WriteHeader(http.StatusNotFound)
                utils.ToJSON(&GenericError{Message: "Message not found in database"}, rw)
                return
        }

        rw.WriteHeader(http.StatusSuccess)
}

func (m *MessagesHandler) GetAllMessages(rw http.ResponseWriter, r *http.Request) {

        messages, err := data.messageDB.GetAllMessages()

        err = data.ToJSON(messages, rw)
        if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&GenericError{Message: err.Error()}, rw)
                return
        }
}

func (m *MessagesHandler) GetMessage(rw http.ResponseWriter, r *http.Request) {

        id := getProductID(r)

        message, err := data.messageDB.GetMessage(id)

        switch err {
        case nil:

        case model.ErrProductNotFound:
                rw.WriteHeader(http.StatusNotFound)
                utils.ToJSON(&GenericError{Message: err.Error()}, rw)
                return
        default:
                rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&GenericError{Message: err.Error()}, rw)
                return
        }

        err = utils.ToJSON(message, rw)
        if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&GenericError{Message: err.Error()}, rw)
                return
        }
}

func (m *MessagesHandler) DeleteMessage(rw http.ResponseWriter, r *http.Request) {

        id := getProductID(r)

        err := data.messageDB.DeleteMessage(id)
        if err == model.ErrProductNotFound {
                rw.WriteHeader(http.StatusNotFound)
                utils.ToJSON(&GenericError{Message: err.Error()}, rw)
                return
        }

        if err != nil {
                rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&GenericError{Message: err.Error()}, rw)
                return
        }

        rw.WriteHeader(http.StatusNoContent)
}

func (m *MessagesHandler) getProductID(r *http.Request) int {
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
