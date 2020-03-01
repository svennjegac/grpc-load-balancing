#!/bin/bash

# scale client
kubectl scale deployment ip-service-client-deployment --replicas=2
