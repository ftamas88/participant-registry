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
			"operation_id": "health_get",
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
			"operation_id": "service_get",
			"data": {
				"api_documentation": "",
				"description": "",
				"develop_url": "",
				"original_build_specification": "",
				"repository": "",
				"service_name": "",
				"slos": ""
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
		"operation_id": "root_get",
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
