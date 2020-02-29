#!/bin/bash

# set script directory as current directory (for relative paths)
cd $(dirname $0)

# build docker image
docker build -t ip-service-server .

# tag the image for the gcr.io registry
docker tag ip-service-server gcr.io/sandbox-269511/ip-service-server

# push the image to the gcr.io registry
docker push gcr.io/sandbox-269511/ip-service-server