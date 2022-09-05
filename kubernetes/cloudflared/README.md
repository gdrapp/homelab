# cloudflared

cloudflared tunnel server

## Installation

Apply the cloudflared Kubernetes config
```bash
kubectl apply -k kubernetes/cloudflared
```

Create the tunnel token secret key
```bash
kubectl create secret generic cloudflared -n cloudflared --from-literal tunnel-token=<TOKEN>
```