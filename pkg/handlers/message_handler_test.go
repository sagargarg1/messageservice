package handlers

import (
        "net/http"
	"net/http/httptest"
        "testing"
	"encoding/json"
	"bytes"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
        "github.com/hashicorp/go-hclog"
        "github.com/sagargarg1/messageservice/pkg/data"
        "github.com/sagargarg1/messageservice/pkg/model"
)

func TestGetAllMessages(t *testing.T) {

	Logging := hclog.Default()
        DB := data.NewMessageDB()
        MessageHandler := NewMessageHandler(Logging, DB)

	req := httptest.NewRequest("GET", "/messageservice/v1/message/all", nil)
	resp := httptest.NewRecorder()
	MessageHandler.GetAllMessages(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetMessage(t *testing.T) {

        Logging := hclog.Default()
        DB := data.NewMessageDB()
        MessageHandler := NewMessageHandler(Logging, DB)

        req := httptest.NewRequest("GET", "/messageservice/v1/message/1", nil)
	req = mux.SetURLVars(req, map[string]string{
		"id": "1",
	})
        resp := httptest.NewRecorder()
        MessageHandler.GetMessage(resp, req)
	message := &model.Message{}
	_ = json.Unmarshal(resp.Body.Bytes(), &message)
        assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "I love India", message.Text)
}

func TestAddMessage(t *testing.T) {

        Logging := hclog.Default()
        DB := data.NewMessageDB()
        MessageHandler := NewMessageHandler(Logging, DB)

	message := model.Message{
		Text: "I am sorry",
	}
	body, _ := json.Marshal(message)
        req := httptest.NewRequest("POST", "/messageservice/v1/message", bytes.NewBuffer(body))
        resp := httptest.NewRecorder()
        MessageHandler.AddMessage(resp, req)
	var id int
	_ = json.Unmarshal(resp.Body.Bytes(), &id)
        assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, 3, id)	
}

func TestUpdateMessage(t *testing.T) {

        Logging := hclog.Default()
        DB := data.NewMessageDB()
        MessageHandler := NewMessageHandler(Logging, DB)

	message := model.Message{
                1,
                "I am sorry",
        }
	body, _ := json.Marshal(message)
        req := httptest.NewRequest("PUT", "/messageservice/v1/message", bytes.NewBuffer(body))
        resp := httptest.NewRecorder()
        MessageHandler.UpdateMessage(resp, req)
        assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteMessage(t *testing.T) {

        Logging := hclog.Default()
        DB := data.NewMessageDB()
        MessageHandler := NewMessageHandler(Logging, DB)

        req := httptest.NewRequest("DELETE", "/messageservice/v1/message/4", nil)
	req = mux.SetURLVars(req, map[string]string{
                "id": "4",
        })
        resp := httptest.NewRecorder()
        MessageHandler.DeleteMessage(resp, req)
        assert.Equal(t, http.StatusNotFound, resp.Code)
}
