# kea

Kea DHCP server

## Installation

Apply the Kea Kubernetes config
```bash
kubectl apply -k kubernetes/kea
```

Apply the Kea ConfigMap
```bash
kubectl apply -n kea -f kubernetes/kea/configmap.yaml
```

## Configuration

Edit the Kea config
```bash
kubectl -n kea edit configmap kea
```

Restart Kea after editing the ConfigMap
```bash
kubectl -n kea rollout restart statefulset
```

Verify pod restarted successfully
```bash
kubectl -n kea get pod --watch
```