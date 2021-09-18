# Kubernetes Dashboard

Kubernetes Dashboard is a general purpose, web-based UI for Kubernetes clusters. It allows users to manage applications running in the cluster and troubleshoot them, as well as manage the cluster itself.

## Installation

Apply the Kubernetes config
```bash
kubectl apply -k kubernetes/kubernetes-dashboard
```

## Usage

Get the authentication token for the service account
```bash
kubectl -n kube-system describe secret $(kubectl -n kube-system get secret | grep admin | awk '{print $1}')
```