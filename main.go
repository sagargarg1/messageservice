package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"github.com/sagargarg1/messageservice/pkg/handlers"
)

var (
	router = mux.NewRouter()
)

func main() {

	mapUrls()
	srv := &http.Server{
		Addr: ":8081",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	//logger.Info("about to start the application...")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func mapUrls() {
	router.HandleFunc("/messageservice/v1/message", handlers.HandlerInterface.AddMessage).Methods(http.MethodPost)
	router.HandleFunc("/messageservice/v1/message", handlers.HandlerInterface.UpdateMessage).Methods(http.MethodPut)
        router.HandleFunc("/messageservice/v1/message/{messageID}", handlers.HandlerInterface.GetMessage).Methods(http.MethodGet)
        router.HandleFunc("/messageservice/v1/message/{messageID}", handlers.HandlerInterface.DeleteMessage).Methods(http.MethodDelete)
        router.HandleFunc("/messageservice/v1/message/all", handlers.HandlerInterface.GetAllMessages).Methods(http.MethodGet)
}
