# Home Assistant

Open source home automation that puts local control and privacy first.

## Installation

Apply the Kubernetes config
```bash
kubectl apply -k kubernetes/homeassistant
```

Modify the Home Assistant config to allow access through Kubernetes ingress
```bash
kubectl exec -it homeassistant-0 -n homeassistant -- vi /config/configuration.yaml
```

Add the following to the configuration, substitute 10.x.x.x for the ingress IP address
```yaml
http:
  use_x_forwarded_for: true
  trusted_proxies:
    - 10.x.x.x
```