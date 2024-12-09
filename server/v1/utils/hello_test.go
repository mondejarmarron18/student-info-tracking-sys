package utils_test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	mux := http.NewServeMux()

	req := httptest.NewRequest(http.MethodGet, "http://localhost:3000/v1/users", nil)

	statusCpde := req.Response.StatusCode

	if statusCpde != http.StatusOK {
		log.Fatalf("Error creating request")
	}

	recorder := httptest.NewRecorder()

	mux.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		log.Fatalf("Error creating request")
	}

	expected := "testing"

	if recorder.Body.String() != expected {
		log.Fatalf("Error creating request")
	}
}
