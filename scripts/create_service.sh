#!/usr/bin/env bash

#needs to run only once per project and per environment

export ENV="dev"

export PROJECT_ID="things-running"
export SERVICE_ACCOUNT_NAME="agritech-lora-ingress-${USER}"
export SERVICE_ACCOUNT_EMAIL="${SERVICE_ACCOUNT_NAME}@${PROJECT_ID}.iam.gserviceaccount.com"

#gcloud iam service-accounts create \
#    ${SERVICE_ACCOUNT_NAME} \
#    --display-name="Cloud Run LoRa ingress Service Account - for ${USER}@${ENV}" \
#    --description="Used to run the LoRa ingress services in cloud-run"

#gcloud iam service-accounts keys create \
#    ${PROJECT_ID}-${ENV}-4b06ebd5e307.json \
#    --iam-account=${SERVICE_ACCOUNT_EMAIL}

gcloud projects add-iam-policy-binding ${PROJECT_ID} \
    --member="serviceAccount:${SERVICE_ACCOUNT_EMAIL}" \
    --role="roles/pubsub.publisher"