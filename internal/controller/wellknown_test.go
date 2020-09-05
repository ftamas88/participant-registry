package controller

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWellKnownController_Health(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://example.com/api/health", nil)
	wk := &WellKnownController{}
	wk.Health(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.JSONEq(t, `
		{
			"data": {
				"status": "OK"
			},
			"metadata": null
		}
	`, string(body))
}

func TestWellKnownController_ServiceInformation(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://example.com/api/service", nil)
	wk := &WellKnownController{}
	wk.ServiceInformation(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.JSONEq(t, `
		{
			"data": {
				"api_documentation": "",
				"description": "This is a participant registry microservice which supports adding, updating, removing and retrieving personal information about participants in the study.",
				"develop_url": "http://localhost:3000",
				"repository": "https://gitlab.com/ftamas88/grail",
				"service_name": "Grail - participant registry",
				"slos":""
			}, 
			"metadata": null
		}
	`, string(body))
}

func TestWellKnownController_Root(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://example.com/", nil)
	wk := &WellKnownController{
		AppVersion:    "1.0.0",
		AppCommitHash: "123abc",
	}
	wk.Root(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.JSONEq(t, `
	{
		"data": {
			"commit_hash": "123abc",
			"service": "grail-participant-registry",
			"version": "1.0.0"
		},
		"metadata": null
	}
	`, string(body))
}

func TestWellKnownController_Version(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "http://example.com/", nil)
	wk := &WellKnownController{
		AppVersion:    "1.0.0",
		AppCommitHash: "123abc",
	}
	wk.Version(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.JSONEq(t, `
	{
		"commit_hash": "123abc",
		"version": "1.0.0"
	}
	`, string(body))
}
