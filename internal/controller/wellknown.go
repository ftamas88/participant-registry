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
		OperationID: "root_get",
		Data: map[string]string{
			"service":     "grail-participant-registry",
			"version":     c.AppVersion,
			"commit_hash": c.AppCommitHash,
		},
	})
}

func (c *WellKnownController) Health(w http.ResponseWriter, _ *http.Request) {
	WriteJSONResponse(w, http.StatusOK, Response{
		OperationID: "health_get",
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
		OperationID: "service_get",
		Data: map[string]string{
			"service_name":                 "",
			"original_build_specification": "",
			"slos":                         "",
			"api_documentation":            "",
			"repository":                   "",
			"description":                  "",
			"develop_url":                  "",
		},
	})
}
