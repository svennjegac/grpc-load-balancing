#!/bin/bash

# logs for each of the client pods
kubectl get pods --selector=app=ip-service-client | grep ip-service-client | while read line
do
  POD="$(cut -d' ' -f1 <<< $line)"
  echo "Pod: $POD"
  kubectl logs $POD
done
