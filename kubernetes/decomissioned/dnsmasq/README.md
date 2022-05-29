# dnsmasq

dnsmasq DNS and DHCP server

## Installation

Apply the dnsmasq Kubernetes config
```bash
kubectl apply -k kubernetes/dnsmasq
```

Apply the dnsmasq ConfigMap
```bash
kubectl apply -n dnsmasq -f configmap.yaml
```

## Configuration

Edit the dnsmasq config
```bash
kubectl -n dnsmasq edit configmap dnsmasq
```

Restart dnsmasq after editing the ConfigMap
```bash
kubectl -n dnsmasq rollout restart statefulsets
```

Verify pod restarted successfully
```bash
kubectl -n dnsmasq get pod --watch
```