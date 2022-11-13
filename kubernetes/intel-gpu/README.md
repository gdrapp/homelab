# Intel GPU device plugin for Kubernetes

Intel GPU plugin facilitates Kubernetes workload offloading by providing access to Intel discrete (Xe) and integrated GPU HW device files.

## Installation

```bash
kubectl apply -k https://github.com/intel/intel-device-plugins-for-kubernetes/deployments/gpu_plugin?ref=v0.25.0
```