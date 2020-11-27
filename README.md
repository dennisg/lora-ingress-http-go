## LoRa Ingress via HTTP written in Go

This project is able to ingest LoRa (TTN) uplink messages and forward them to
a PubSub topic.

## Prerequisites

- A Google project, with enough credentials to have some fun 
- A Linux PC
- Go 1.15 installed
- gcloud installed
- A PubSub topic, create it and note the name.

## Getting started

First create a .env file in the root of the project with the following contents:

```
PROJECT_ID="<your google project id>"
ROOT="agritech"
ACCESS_TOKEN="<an access token>"
TARGET_TOPIC="projects/{projectId}/topics/<the topic you created>"
```

To access that pubsub topic inspect and run from the root of the project you need to create a service account:

```./scripts/create_service.sh```

If all is well, a json file is generated that provides your local development environment
access to the PubSub topic.

## Running the Project

From the root of the project, run:

```./scripts/debug.sh```

The Docker image is built and run, output should be alike:

```
{"level":"info","msg":"listening on port :8080","time":"2020-11-27T22:00:10+01:00"}
```


## Deploying to the Cloud using Google Cloud Build

```
➜  lora-ingress-http-go git:(master) ✗ ./scripts/submit.sh
Deploying...
Creating temporary tarball archive of 1090 file(s) totalling 13.0 MiB before compression.
Some files were not included in the source upload.

Check the gcloud log [/home/dennisg/.config/gcloud/logs/2020.11.27/22.08.09.343407.log] to see which files and the contents of the
default gcloudignore file used (see `$ gcloud topic gcloudignore` to learn
more).

Uploading tarball of [.] to [gs://things-running_cloudbuild/source/1606511289.47-fdf8ccf8dad74566922e73e55cbd62d0.tgz]
Created [https://cloudbuild.googleapis.com/v1/projects/things-running/builds/a0819c67-d8e0-42aa-8df1-e9eb40c59e90].
Logs are available at [https://console.cloud.google.com/cloud-build/builds/a0819c67-d8e0-42aa-8df1-e9eb40c59e90?project=566942708108].
--------------------------------------------------------------------------------------------------------- REMOTE BUILD OUTPUT ----------------------------------------------------------------------------------------------------------
starting build "a0819c67-d8e0-42aa-8df1-e9eb40c59e90"

FETCHSOURCE
Fetching storage object: gs://things-running_cloudbuild/source/1606511289.47-fdf8ccf8dad74566922e73e55cbd62d0.tgz#1606511294344204
Copying gs://things-running_cloudbuild/source/1606511289.47-fdf8ccf8dad74566922e73e55cbd62d0.tgz#1606511294344204...
/ [1 files][  2.8 MiB/  2.8 MiB]                                                
Operation completed over 1 objects/2.8 MiB.                                      
BUILD
Starting Step #0
Step #0: Already have image (with digest): gcr.io/cloud-builders/docker
Step #0: Sending build context to Docker daemon   14.6MB
Step #0: Step 1/18 : FROM golang:1.15 as builder
Step #0: 1.15: Pulling from library/golang
Step #0: 756975cb9c7e: Pulling fs layer
Step #0: d77915b4e630: Pulling fs layer
Step #0: 5f37a0a41b6b: Pulling fs layer
Step #0: 96b2c1e36db5: Pulling fs layer
Step #0: 145393847161: Pulling fs layer
Step #0: 71dfa979a65c: Pulling fs layer
Step #0: 88a83f11b30a: Pulling fs layer
Step #0: 96b2c1e36db5: Waiting
Step #0: 145393847161: Waiting
Step #0: 71dfa979a65c: Waiting
Step #0: 88a83f11b30a: Waiting
Step #0: d77915b4e630: Verifying Checksum
Step #0: d77915b4e630: Download complete
Step #0: 5f37a0a41b6b: Verifying Checksum
Step #0: 5f37a0a41b6b: Download complete
Step #0: 756975cb9c7e: Verifying Checksum
Step #0: 756975cb9c7e: Download complete
Step #0: 145393847161: Verifying Checksum
Step #0: 145393847161: Download complete
Step #0: 96b2c1e36db5: Verifying Checksum
Step #0: 96b2c1e36db5: Download complete
Step #0: 88a83f11b30a: Verifying Checksum
Step #0: 88a83f11b30a: Download complete
Step #0: 71dfa979a65c: Verifying Checksum
Step #0: 71dfa979a65c: Download complete
Step #0: 756975cb9c7e: Pull complete
Step #0: d77915b4e630: Pull complete
Step #0: 5f37a0a41b6b: Pull complete
Step #0: 96b2c1e36db5: Pull complete
Step #0: 145393847161: Pull complete
Step #0: 71dfa979a65c: Pull complete
Step #0: 88a83f11b30a: Pull complete
Step #0: Digest: sha256:cf46c759511d0376c706a923f2800762948d4ea1a9290360720d5124a730ed63
Step #0: Status: Downloaded newer image for golang:1.15
Step #0:  ---> 6d8772fbd285
Step #0: Step 2/18 : ADD go.mod /src/
Step #0:  ---> de7c894a9d2a
Step #0: Step 3/18 : ADD go.sum /src/
Step #0:  ---> d95b2a6ed083
Step #0: Step 4/18 : ADD pkg /src/pkg
Step #0:  ---> cd29e2498b62
Step #0: Step 5/18 : ADD cmd /src/cmd
Step #0:  ---> 71af1005ecd0
Step #0: Step 6/18 : ADD vendor /src/vendor
Step #0:  ---> dc83bb3e12ac
Step #0: Step 7/18 : WORKDIR /src/cmd/main
Step #0:  ---> Running in ca2f4a614a27
Step #0: Removing intermediate container ca2f4a614a27
Step #0:  ---> 89707dced89a
Step #0: Step 8/18 : RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main .
Step #0:  ---> Running in 8b0b9cd27bda
Step #0: Removing intermediate container 8b0b9cd27bda
Step #0:  ---> 0eb71f2dd736
Step #0: Step 9/18 : RUN go test
Step #0:  ---> Running in d4033a7b8802
Step #0: ?      lora-ingress-http-go/cmd/main   [no test files]
Step #0: Removing intermediate container d4033a7b8802
Step #0:  ---> 93c37410086e
Step #0: Step 10/18 : FROM scratch as container
Step #0:  ---> 
Step #0: Step 11/18 : COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /lib/time/zoneinfo.zip
Step #0:  ---> a700647aec5b
Step #0: Step 12/18 : COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
Step #0:  ---> 64985d0e5ee1
Step #0: Step 13/18 : ENV PORT=8080
Step #0:  ---> Running in cd3616728e7b
Step #0: Removing intermediate container cd3616728e7b
Step #0:  ---> 55c9460da359
Step #0: Step 14/18 : ENV GOROOT=/
Step #0:  ---> Running in 0243d1e0221b
Step #0: Removing intermediate container 0243d1e0221b
Step #0:  ---> a3923988b921
Step #0: Step 15/18 : ENV TZ=Europe/Amsterdam
Step #0:  ---> Running in dbb2467d6073
Step #0: Removing intermediate container dbb2467d6073
Step #0:  ---> 6cec2cd9e7ad
Step #0: Step 16/18 : WORKDIR /
Step #0:  ---> Running in a7291a6f4978
Step #0: Removing intermediate container a7291a6f4978
Step #0:  ---> 9c75de37a24d
Step #0: Step 17/18 : ENTRYPOINT ["/main"]
Step #0:  ---> Running in 81777b2d457c
Step #0: Removing intermediate container 81777b2d457c
Step #0:  ---> 22f302c10b56
Step #0: Step 18/18 : COPY --from=builder /main /main
Step #0:  ---> 8acb91af9c6e
Step #0: Successfully built 8acb91af9c6e
Step #0: Successfully tagged eu.gcr.io/things-running/agritech/agritech-lora-ingress-http-go:latest
Finished Step #0
Starting Step #1
Step #1: Already have image (with digest): gcr.io/cloud-builders/docker
Step #1: The push refers to repository [eu.gcr.io/things-running/agritech/agritech-lora-ingress-http-go]
Step #1: ae8fdb1b505a: Preparing
Step #1: e494b592a9ac: Preparing
Step #1: 75b71d07f6f2: Preparing
Step #1: 75b71d07f6f2: Pushed
Step #1: e494b592a9ac: Pushed
Step #1: ae8fdb1b505a: Pushed
Step #1: latest: digest: sha256:f0be1da8d79ff955153e7060af537a886595be33a0eb1e135877ed13677dc705 size: 948
Finished Step #1
Starting Step #2
Step #2: Pulling image: gcr.io/google.com/cloudsdktool/cloud-sdk
Step #2: Using default tag: latest
Step #2: latest: Pulling from google.com/cloudsdktool/cloud-sdk
Step #2: e4c3d3e4f7b0: Already exists
Step #2: fabd81c9382a: Pulling fs layer
Step #2: 3b82b1e4ddb7: Pulling fs layer
Step #2: b8c205461bdc: Pulling fs layer
Step #2: 47936139ff4d: Pulling fs layer
Step #2: 47936139ff4d: Waiting
Step #2: fabd81c9382a: Verifying Checksum
Step #2: fabd81c9382a: Download complete
Step #2: 47936139ff4d: Verifying Checksum
Step #2: 47936139ff4d: Download complete
Step #2: b8c205461bdc: Verifying Checksum
Step #2: b8c205461bdc: Download complete
Step #2: 3b82b1e4ddb7: Verifying Checksum
Step #2: 3b82b1e4ddb7: Download complete
Step #2: fabd81c9382a: Pull complete
Step #2: 3b82b1e4ddb7: Pull complete
Step #2: b8c205461bdc: Pull complete
Step #2: 47936139ff4d: Pull complete
Step #2: Digest: sha256:12e3701513753324519dcca0cd64c7fd23b00621d08aa91cdbcb2b67c7a27f36
Step #2: Status: Downloaded newer image for gcr.io/google.com/cloudsdktool/cloud-sdk:latest
Step #2: gcr.io/google.com/cloudsdktool/cloud-sdk:latest
Step #2: Deploying container to Cloud Run service [agritech-lora-ingress-http-go-dev] in project [things-running] region [europe-west3]
Step #2: Deploying new service...
Step #2: Setting IAM Policy........................done
Step #2: Creating Revision..............................................done
Step #2: Routing traffic.....done
Step #2: Done.
Step #2: Service [agritech-lora-ingress-http-go-dev] revision [agritech-lora-ingress-http-go-dev-00001-pin] has been deployed and is serving 100 percent of traffic.
Step #2: Service URL: https://agritech-lora-ingress-http-go-dev-67u7iac2ua-ey.a.run.app
Finished Step #2
PUSH
Pushing eu.gcr.io/things-running/agritech/agritech-lora-ingress-http-go
The push refers to repository [eu.gcr.io/things-running/agritech/agritech-lora-ingress-http-go]
ae8fdb1b505a: Preparing
e494b592a9ac: Preparing
75b71d07f6f2: Preparing
ae8fdb1b505a: Layer already exists
e494b592a9ac: Layer already exists
75b71d07f6f2: Layer already exists
latest: digest: sha256:f0be1da8d79ff955153e7060af537a886595be33a0eb1e135877ed13677dc705 size: 948
DONE
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

ID                                    CREATE_TIME                DURATION  SOURCE                                                                                    IMAGES                                                                     STATUS
a0819c67-d8e0-42aa-8df1-e9eb40c59e90  2020-11-27T21:08:14+00:00  2M55S     gs://things-running_cloudbuild/source/1606511289.47-fdf8ccf8dad74566922e73e55cbd62d0.tgz  eu.gcr.io/things-running/agritech/agritech-lora-ingress-http-go (+1 more)  SUCCESS

```