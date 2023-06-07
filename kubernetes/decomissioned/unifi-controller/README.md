# unifi-controller

Ubiquity UniFi Controller

# Installation

1. Install the UniFi Controller
```bash
kubectl apply -k kubernetes/unifi-controller
```

3. Wait for the pods to enter a running state
```bash
kubectl -n unifi-controller get pod
```
