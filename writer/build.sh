#!/bin/bash

echo "Building hostpath-writer image..."
docker build -t hostpath-writer .

echo "Pushing to gcr..."
docker tag hostpath-writer gcr.io/jianfeih-test/hostpath-writer
gcloud docker -- push gcr.io/jianfeih-test/hostpath-writer
