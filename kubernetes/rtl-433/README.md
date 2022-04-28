# rtl-433

rtl_433 is a generic data receiver, mainly for the 433.92 MHz, 868 MHz (SRD), 315 MHz, 345 MHz, and 915 MHz ISM bands.

## Configuration

1. Update the configuration environment variables in deployment.yaml

2. Create the MQTT server username and password Kubernetes secret
```bash
kubectl -n rtl-433 create secret generic mqtt --from-literal=username='user' --from-literal=password='password'
```

## Installation

```bash
kubectl apply -k kubernetes/rtl-433
```