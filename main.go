package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/sagargarg1/messageservice/pkg/handlers"
	"github.com/sagargarg1/messageservice/pkg/data"
	"github.com/sagargarg1/messageservice/pkg/middleware"
)

var (
	router = mux.NewRouter()
)

func main() {

	setRoutes()
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

func setRoutes() {

	Logging := hclog.Default()
	DB := data.NewMessageDB()
	MessageHandler := handlers.NewMessageHandler(Logging, DB)
	MetricHandler := handlers.NewMetricsHandler(Logging)

	v1 := router.PathPrefix("/messageservice/v1/message").Subrouter()
	router.Use(middleware.MetricsMiddleware(Logging))
	
	MessageHandler.AddRoutes(v1)
	MetricHandler.AddRoutes(v1)
}
