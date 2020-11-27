FROM golang:1.15 as builder

ADD go.mod /src/
ADD go.sum /src/
ADD pkg /src/pkg
ADD cmd /src/cmd

#all dependencies
ADD vendor /src/vendor

WORKDIR /src/cmd/main
RUN go test
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main .

FROM scratch as container

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /lib/time/zoneinfo.zip
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENV PORT=8080
ENV GOROOT=/
ENV TZ=Europe/Amsterdam

#ENV ACCESS_TOKEN=""
#ENV TARGET_TOPIC="projects/things-running/topics/agritech-lora-ttn-ingress"

WORKDIR /
ENTRYPOINT ["/main"]
COPY --from=builder /main /main

