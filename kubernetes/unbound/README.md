# unbound

Unbound DNS server

## Installation

Apply the Unbound Kubernetes config
```bash
kubectl apply -k kubernetes/unbound
```

Apply the unbound ConfigMap
```bash
kubectl apply -n unbound -f kubernetes/unbound/configmap.yaml
```

## Configuration

Edit the Unbound config
```bash
kubectl -n unbound edit configmap unbound
```

Restart unbound after editing the ConfigMap
```bash
kubectl -n unbound rollout restart deployments
```

Verify pod restarted successfully
```bash
kubectl -n unbound get pod --watch
```