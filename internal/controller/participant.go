package controller

import (
	"encoding/json"
	"grail-participant-registry/internal/domain"
	"grail-participant-registry/internal/service"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type ParticipantController struct {
	Service *service.Participant
}

func (pc *ParticipantController) Index(w http.ResponseWriter, _ *http.Request) {
	WriteJSONResponse(w, http.StatusOK, Response{
		Data: map[string]string{
			"status": "OK",
		},
	})
}

func (pc *ParticipantController) Get(w http.ResponseWriter, r *http.Request) {
	p, err := pc.Service.GetParticipant(mux.Vars(r)["ref"])
	if err != nil {
		handleParticipantError(w, err, http.StatusNotFound)

		return
	}

	WriteJSONResponse(w, http.StatusOK, p)
}

func (pc *ParticipantController) Create(w http.ResponseWriter, r *http.Request) {
	p := domain.Participant{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		handleParticipantError(w, err, http.StatusBadRequest)

		return
	}

	err := pc.Service.CreateParticipant(p)
	if err != nil {
		handleParticipantError(w, err, http.StatusInternalServerError)

		return
	}

	WriteJSONResponse(w, http.StatusCreated, nil)
}

func (pc *ParticipantController) Update(w http.ResponseWriter, r *http.Request) {
	p := domain.Participant{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		handleParticipantError(w, err, http.StatusBadRequest)

		return
	}

	err := pc.Service.UpdateParticipant(mux.Vars(r)["ref"], p)
	if err != nil {
		handleParticipantError(w, err, http.StatusBadRequest)

		return
	}

	WriteJSONResponse(w, http.StatusOK, nil)
}

func (pc *ParticipantController) Delete(w http.ResponseWriter, r *http.Request) {
	err := pc.Service.RemoveParticipant(mux.Vars(r)["ref"])
	if err != nil {
		handleParticipantError(w, err, http.StatusBadRequest)

		return
	}

	WriteJSONResponse(w, http.StatusOK, nil)
}

func handleParticipantError(w http.ResponseWriter, err error, code int) {
	logrus.Warnf(err.Error())
	WriteJSONResponse(w, code, ErrorResponse{
		Error: map[string]string{
			"message": err.Error(),
		},
	})
}
