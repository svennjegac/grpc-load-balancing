#!/bin/bash

# set script directory as current directory (for relative paths)
cd $(dirname $0)

# create client
kubectl create -f ip-service-client.yaml
