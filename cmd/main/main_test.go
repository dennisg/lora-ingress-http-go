package main_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIngressSuccess(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	t.Error("implement me")
}

func TestIngressFail(t *testing.T) {
	t.Error("implement me")
}



