package main

//go:generate docker image ls lora-ingress-http-go

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"lora-ingress-http-go/pkg/domain"
	"net/http"
	"os"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

var accessToken = os.Getenv("ACCESS_TOKEN")
var topic = os.Getenv("TARGET_TOPIC")
var port = ":8080"

func init() {
	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Infof("listening on port %s", port)
}

//main entrypoint
func main() {
	messageBus, err := domain.NewMessageBus(topic)
	if err != nil {
		logrus.Fatalf("unable to create message bus: %v", err)
	}
	handler := ingress(messageBus)
	http.HandleFunc("/api/ttn/ingress", middleware(handler))
	logrus.Fatal(http.ListenAndServe(port, nil))
}

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

//create the ingress endpoint
func ingress(bus domain.MessageBus) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var ttnMessage = domain.TtnMessage{}
		err := json.NewDecoder(r.Body).Decode(&ttnMessage)
		if err != nil {
			logrus.Warnf("unable to decode ingress message: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			err = publish(ttnMessage, bus)
			if err != nil {
				logrus.Errorf("unable to forward TTN message to the Message Bus: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusCreated)
			}
		}
	}
}

func publish(message domain.TtnMessage, bus domain.MessageBus) error {
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Errorf("unable to generate uuid: %v", err)
		return err
	}

	loraEvent := cloudevents.NewEvent()
	loraEvent.SetID(id.String())
	loraEvent.SetSource(fmt.Sprintf("lora-ingress-http/%s", message.AppID))
	loraEvent.SetType("lora.ttn.message")
	loraEvent.SetTime(message.Metadata.Time)

	err = loraEvent.SetData(cloudevents.ApplicationJSON, message)
	if err != nil {
		logrus.Errorf("unable to set CloudEvent data: %v", err)
		return err
	}
	data, err := loraEvent.MarshalJSON()
	if err != nil {
		logrus.Errorf("unable to encode cloud event: %v", err)
		return err
	}
	return bus.Send(loraEvent.ID(), data)
}
