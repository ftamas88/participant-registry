package controller

import "net/http"

type ParticipantController struct{}

func (p *ParticipantController) Index(w http.ResponseWriter, _ *http.Request) {
	WriteJSONResponse(w, http.StatusOK, Response{
		OperationID: "participant_get",
		Data: map[string]string{
			"status": "OK",
		},
	})
}
