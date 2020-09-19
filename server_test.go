package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePaireDevice(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/pair-device", nil)
	rec := httptest.NewRecorder()

	PairDeviceHandler(rec, req)

	if http.StatusOK != rec.Code {
		t.Error("expecred 200 OK")
	}

	expected := `{"status":"active"}`

	if rec.Body.String() != expected {
		t.Errorf("expected %q but got %q\n", expected, rec.Body)
	}

}
