package handlers

import (
	"encoding/json"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"lora-ingress-http-go/pkg/actuator"
	"lora-ingress-http-go/pkg/domain"
	"net/http"
	"os"
)

const cloudEventSource = "lora://ingress-http/%s"

var topic = os.Getenv("TARGET_TOPIC")
var accessToken = os.Getenv("ACCESS_TOKEN")

func Initialize() {
	messageBus, err := domain.NewMessageBus(topic)
	if err != nil {
		logrus.Fatalf("unable to create message bus: %v", err)
	}
	actuator.Initialize()

	handler := ingress(messageBus)
	http.HandleFunc("/api/ttn/ingress", middleware(handler))
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
			logrus.Infof("LoRa uplink message received for %s:%s", ttnMessage.AppID, ttnMessage.DevID)
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



//publish the message, converting it first into a CloudEvent
func publish(message domain.TtnMessage, bus domain.MessageBus) error {
	id, err := uuid.NewUUID()
	if err != nil {
		logrus.Errorf("unable to generate uuid: %v", err)
		return err
	}

	loraEvent := cloudevents.NewEvent()
	loraEvent.SetID(id.String())
	loraEvent.SetSource(fmt.Sprintf(cloudEventSource, message.AppID))
	loraEvent.SetType("lora.ttn.message")
	loraEvent.SetTime(message.Metadata.Time)

	err = loraEvent.SetData(cloudevents.ApplicationJSON, message)
	if err != nil {
		logrus.Errorf("unable to set CloudEvent data: %v", err)
		return err
	}

	data, err := loraEvent.MarshalJSON()
	if err != nil {
		logrus.Errorf("unable to marshal the cloud event: %v", err)
		return err
	}
	return bus.Send(loraEvent.ID(), data)
}
