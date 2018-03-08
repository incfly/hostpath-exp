#!/bin/bash

echo "Building hostpath-writer image..."
docker build -t hostpath-writer .

echo "Pushing to gcr..."
docker tag hostpath-writer gcr.io/jianfeih-test/hostpath-writer
gcloud docker -- push gcr.io/jianfeih-test/hostpath-writer

# then login into the pod, write some content
kc exec -it nginx-deployment-789f59656-6j62d    -- bash

# inside the container
cd /path
echo "something" > content.txt

# TODO: verify this.
# when workload is created, could be prior to the nodeagent create the
# serviceaccount directory. so the dir is created by mount/k8s, then able to be
# seen by the nodeagent?

# ssh into the cluster node from patheon, confirm the
# /home/kubernetes/hostpath-driver/sa1, exists.

