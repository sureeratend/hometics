package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePaireDevice(t *testing.T) {
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(Pair{DeviceID: 1234, UserID: 4433})
	req := httptest.NewRequest(http.MethodPost, "/pair-device", payload)
	rec := httptest.NewRecorder()

	handler := &PairDeviceHandler{createPairDevice: func(p Pair) error {
		return nil
	}}

	handler.ServeHTTP(rec, req)

	if http.StatusOK != rec.Code {
		t.Error("expecred 200 OK")
	}

	expected := `{"status":"active"}`

	if rec.Body.String() != expected {
		t.Errorf("expected %q but got %q\n", expected, rec.Body)
	}

}
