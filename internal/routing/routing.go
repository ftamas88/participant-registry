package routing

import (
	"net/http"

	"grail-participant-registry/internal/controller"

	"github.com/gorilla/mux"
)

type RouterConfig struct {
	WellKnown *controller.WellKnownController
}

func NewRouter(conf *RouterConfig) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", conf.WellKnown.Root).Methods(http.MethodGet)

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/health", conf.WellKnown.Health).Methods(http.MethodGet)
	api.HandleFunc("/service", conf.WellKnown.ServiceInformation).Methods(http.MethodGet)
	api.HandleFunc("/version", conf.WellKnown.Version).Methods(http.MethodGet)

	return r
}
