#!/bin/bash

# set script directory as current directory (for relative paths)
cd $(dirname $0)

# build docker image
docker build -t ip-service-client .

# tag the image for the gcr.io registry
docker tag ip-service-client gcr.io/sandbox-269511/ip-service-client

# push the image to the gcr.io registry
docker push gcr.io/sandbox-269511/ip-service-client
