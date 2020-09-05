package controller

import (
	"net/http"
)

type WellKnownController struct {
	AppVersion    string
	AppCommitHash string
}

func (c *WellKnownController) Root(w http.ResponseWriter, _ *http.Request) {
	WriteJSONResponse(w, http.StatusOK, Response{
		Data: map[string]string{
			"service":     "grail-participant-registry",
			"version":     c.AppVersion,
			"commit_hash": c.AppCommitHash,
		},
	})
}

func (c *WellKnownController) Health(w http.ResponseWriter, _ *http.Request) {
	WriteJSONResponse(w, http.StatusOK, Response{
		Data: map[string]string{
			"status": "OK",
		},
	})
}

func (c *WellKnownController) Version(w http.ResponseWriter, _ *http.Request) {
	WriteJSONResponse(w, http.StatusOK, map[string]string{
		"version":     c.AppVersion,
		"commit_hash": c.AppCommitHash,
	})
}

func (c *WellKnownController) ServiceInformation(w http.ResponseWriter, r *http.Request) {
	WriteJSONResponse(w, http.StatusOK, Response{
		Data: map[string]string{
			"service_name":      "Grail - participant registry",
			"slos":              "",
			"api_documentation": "",
			"repository":        "https://gitlab.com/ftamas88/grail",
			"description": "This is a participant registry microservice which supports adding, updating, " +
				"removing and retrieving personal information about participants in the study.",
			"develop_url": "http://localhost:3000",
		},
	})
}
