#!/bin/bash

# set script directory as current directory (for relative paths)
cd $(dirname $0)

# delete server
kubectl delete -f ip-service-server.yaml
