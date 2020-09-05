package routing

import (
	"net/http"

	"grail-participant-registry/internal/controller"

	"github.com/gorilla/mux"
)

type RouterConfig struct {
	WellKnown   *controller.WellKnownController
	Participant *controller.ParticipantController
}

func NewRouter(conf *RouterConfig) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", conf.WellKnown.Root).Methods(http.MethodGet)

	api := r.PathPrefix("/api").Subrouter()

	v1 := api.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/participants", conf.Participant.Index).Methods(http.MethodGet)

	api.HandleFunc("/health", conf.WellKnown.Health).Methods(http.MethodGet)
	api.HandleFunc("/service", conf.WellKnown.ServiceInformation).Methods(http.MethodGet)
	api.HandleFunc("/version", conf.WellKnown.Version).Methods(http.MethodGet)

	return r
}
