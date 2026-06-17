#!/bin/bash

kubectl apply -f deployments/kubernetes/base/namespace.yaml

kubectl apply -f deployments/kubernetes/base/edge-deployment.yaml

kubectl apply -f deployments/kubernetes/base/service.yaml
