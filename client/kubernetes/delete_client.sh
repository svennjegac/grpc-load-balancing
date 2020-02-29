#!/bin/bash

# set script directory as current directory (for relative paths)
cd $(dirname $0)

# delete client
kubectl delete -f ip-service-client.yaml
