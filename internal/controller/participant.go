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
		OperationID: "participant_get",
		Data: map[string]string{
			"status": "OK",
		},
	})
}

func (pc *ParticipantController) Create(w http.ResponseWriter, r *http.Request) {
	p := domain.Participant{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		logrus.Warnf("invalid participant request body: %s", err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, ErrorResponse{
			OperationID: "participant_add_post",
			Error: map[string]string{
				"message": err.Error(),
			},
		})

		return
	}

	err := pc.Service.Repository().Add(p)
	if err != nil {
		logrus.Warnf("unable to create participant: %s", err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, ErrorResponse{
			OperationID: "participant_add_post",
			Error: map[string]string{
				"message": err.Error(),
			},
		})

		return
	}

	WriteJSONResponse(w, http.StatusCreated, nil)
}

func (pc *ParticipantController) Update(w http.ResponseWriter, r *http.Request) {
	p, err := pc.Service.Repository().FindByID(
		domain.ParticipantReference(
			mux.Vars(r)["ref"],
		),
	)
	if err != nil {
		logrus.Warnf("unable to find participant: %s", err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, ErrorResponse{
			OperationID: "participant_update_post",
			Error: map[string]string{
				"message": err.Error(),
			},
		})
	}

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		logrus.Warnf("invalid participant request body: %s", err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, ErrorResponse{
			OperationID: "participant_update_post",
			Error: map[string]string{
				"message": err.Error(),
			},
		})

		return
	}

	err = pc.Service.Repository().Update(p)
	if err != nil {
		logrus.Warnf("unable to update participant: %s", err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, ErrorResponse{
			OperationID: "participant_update_post",
			Error: map[string]string{
				"message": err.Error(),
			},
		})

		return
	}

	WriteJSONResponse(w, http.StatusOK, nil)
}

func (pc *ParticipantController) Fetch(w http.ResponseWriter, _ *http.Request) {
	ps, err := pc.Service.Repository().FindAll()
	if err != nil {
		logrus.Warnf("unable to fetch participants: %s", err.Error())
		WriteJSONResponse(w, http.StatusBadRequest, ErrorResponse{
			OperationID: "participant_fetch_get",
			Error: map[string]string{
				"message": err.Error(),
			},
		})

		return
	}

	WriteJSONResponse(w, http.StatusOK, Response{
		OperationID: "participant_get",
		Data:        ps,
	})
}
