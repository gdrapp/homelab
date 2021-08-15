# rtlamr

rtlamr is an rtl-sdr receiver for Itron ERT compatible smart meters operating in the 900MHz ISM band.

## Configuration

1. Update the configuration environment variables in deployment.yaml

2. Create the MQTT server username and password Kubernetes secret
```bash
kubectl -n rtlamr create secret generic mqtt --from-literal=username='user' --from-literal=password='password'
```

## Installation

```bash
kubectl apply -k kubernetes/rtlamr
```