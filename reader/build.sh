#!/bin/bash

docker build -t hostpath-workload .

echo "Pushing to gcr..."
docker tag hostpath-workload gcr.io/jianfeih-test/hostpath-workload
gcloud docker -- push gcr.io/jianfeih-test/hostpath-workload

