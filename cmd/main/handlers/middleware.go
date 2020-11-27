package handlers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

//provide security, simple for now
//TODO: replace with JWT validation
func middleware(ingress func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			logrus.Warnf("Illegal access from: %s", r.RemoteAddr)
			w.WriteHeader(http.StatusForbidden)
			return
		}
		var bearer = r.Header.Get("Authorization")
		if bearer == fmt.Sprintf("Bearer %s", accessToken) {
			ingress(w, r)
		} else {
			logrus.Warnf("Illegal access from: %s", r.RemoteAddr)
			w.WriteHeader(http.StatusForbidden)
		}
	}
}
