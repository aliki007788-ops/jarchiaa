#!/bin/bash

echo "?? Deploying backend..."
docker-compose -f docker-compose.prod.yml build
docker-compose -f docker-compose.prod.yml push
kubectl apply -f infrastructure/kubernetes/
echo "? Backend deployed"


