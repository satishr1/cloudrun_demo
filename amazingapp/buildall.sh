#!/bin/bash

gcloud services enable cloudbuild.googleapis.com
gcloud services enable run.googleapis.com
gcloud builds submit --tag gcr.io/${DEVSHELL_PROJECT_ID}/amazingapp:blue blue/
gcloud builds submit --tag gcr.io/${DEVSHELL_PROJECT_ID}/amazingapp:green green/
gcloud builds submit --tag gcr.io/${DEVSHELL_PROJECT_ID}/amazingapp:purple purple/
