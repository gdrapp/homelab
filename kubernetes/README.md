# Building the K3S cluster

The cluster will consist of three nodes all running both the server and the agent processes.

## Installing the first node 

```bash
curl -sfL https://get.k3s.io | INSTALL_K3S_VERSION=v1.21.2+k3s1 sh -s - server --disable=traefik --disable=servicelb --cluster-init
```

## Installing subsequent nodes

Get token on first node at /var/lib/rancher/k3s/server/node-token

```bash
curl -sfL https://get.k3s.io | INSTALL_K3S_VERSION=v1.21.2+k3s1 sh -s - server --disable=traefik --disable=servicelb --server https://<ip or hostname of node1>:6443 --token <TOKEN> 
```