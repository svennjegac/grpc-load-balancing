#!/bin/bash

# scale server
kubectl scale deployment ip-service-server-deployment --replicas=2
