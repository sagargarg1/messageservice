package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/sagargarg1/messageservice/pkg/utils"
	"github.com/sagargarg1/messageservice/pkg/model"
)

type MetricsHandlerInterface interface {
	AddRoutes(router *mux.Router)
        GetMetrics(rw http.ResponseWriter, r *http.Request)
}


type MetricsHandler struct {
	Logging	hclog.Logger
}

func NewMetricsHandler(Logging hclog.Logger) MetricsHandlerInterface {
	return &MetricsHandler{
		Logging:        Logging,
	}
}

func (m *MetricsHandler) AddRoutes(router *mux.Router) {
	routes := router.PathPrefix("").Subrouter()
        routes.HandleFunc("/metrics", m.GetMetrics).Methods(http.MethodGet)
}

// responses:
//      200: Success
//      500: InternalError
func (m *MetricsHandler) GetMetrics(rw http.ResponseWriter, r *http.Request) {

	m.Logging.Debug("Get metrics \n")
        err := utils.ToJSON(model.Metrics, rw)
        if err != nil {
		m.Logging.Error("Failed to encode response", "error", err)
		rw.WriteHeader(http.StatusInternalServerError)
                utils.ToJSON(&model.GenericError{Message: err.Error()}, rw)
                return
        }
	rw.WriteHeader(http.StatusOK)
        m.Logging.Debug("Successfully retrieved metrics \n")
}
