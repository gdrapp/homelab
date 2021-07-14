# MetalLB
MetalLB hooks into your Kubernetes cluster, and provides a network load-balancer implementation. In short, it allows you to create Kubernetes services of type “LoadBalancer” in clusters that don’t run on a cloud provider, and thus cannot simply hook into paid products to provide load-balancers.

## Configuration
Set IP address range in configmap.yaml

## Installation
```bash
sudo kubectl apply -k kubernetes/metallb/
```

