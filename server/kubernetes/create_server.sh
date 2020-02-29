#!/bin/bash

# set script directory as current directory (for relative paths)
cd $(dirname $0)

# create server
kubectl create -f ip-service-server.yaml
