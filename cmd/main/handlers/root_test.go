package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"lora-ingress-http-go/pkg/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type testBus struct{}

func (t *testBus) Send(id string, data []byte) error {
	if strings.Contains(string(data), "error") {
		return errors.New(id)
	}
	return nil
}

func newBus() domain.MessageBus {
	return &testBus{}
}

func TestIngressSuccess(t *testing.T) {
	handler := ingress(newBus())

	message := domain.TtnMessage{AppID: "test", Metadata: domain.Metadata{Time: time.Now()}}
	payload, _ := json.Marshal(message)

	req := httptest.NewRequest("POST", "http://example.com/api/ttn/ingress", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	handler(w, req)

	if w.Result().StatusCode != http.StatusCreated {
		t.Errorf("invalid return status code: %d", w.Result().StatusCode)
	}
}

func TestIngressFail(t *testing.T) {
	handler := ingress(newBus())

	message := domain.TtnMessage{AppID: "error", Metadata: domain.Metadata{Time: time.Now()}}
	payload, _ := json.Marshal(message)

	req := httptest.NewRequest("POST", "http://example.com/api/ttn/ingress", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	handler(w, req)

	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Errorf("invalid return status code: %d", w.Result().StatusCode)
	}
}
