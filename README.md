# NodeAgent Prototype

## Instructions

Start Writer

Start Reader

Verify File Visible for Reader

## Further Work

### Which user to run

Without specifying, container runs as root user,
https://docs.docker.com/engine/reference/builder/#user.

All workloads, nodeagent able to modify the mounted dir. Seems fine though,
because the scope is limited by mount option. Also on the other hand, we can't
restrict whathever user the workload is running, that's up to istio customer's
choice.

https://kubernetes.io/docs/concepts/storage/volumes/#hostpath
