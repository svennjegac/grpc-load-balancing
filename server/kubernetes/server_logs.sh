#!/bin/bash

# logs for each of the server pods
kubectl get pods --selector=app=ip-service-server-pod-selector | grep ip-service-server | while read line
do
  POD="$(cut -d' ' -f1 <<< $line)"
  echo "Pod: $POD"
  kubectl logs $POD
done
