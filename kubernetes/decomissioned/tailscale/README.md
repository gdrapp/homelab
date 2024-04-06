# tailscale

tailscale subnet router

## Installation

Apply the tailscale Kubernetes config
```bash
kubectl apply -k kubernetes/tailscale
```

Create the tailscale auth secret key
```bash
kubectl create secret generic tailscale-auth -n tailscale --from-literal TS_AUTH_KEY=<TOKEN>
```