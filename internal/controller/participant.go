package controller

import (
	"grail-participant-registry/internal/service"
	"net/http"

	"github.com/sirupsen/logrus"
)

type ParticipantController struct {
	Service *service.Participant
}

func (p *ParticipantController) Index(w http.ResponseWriter, _ *http.Request) {
	WriteJSONResponse(w, http.StatusOK, Response{
		OperationID: "participant_get",
		Data: map[string]string{
			"status": "OK",
		},
	})
}

func (p *ParticipantController) Fetch(w http.ResponseWriter, _ *http.Request) {
	ps, err := p.Service.Repository().FindAll()
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
