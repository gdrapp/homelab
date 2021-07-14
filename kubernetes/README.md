# Building the K3S cluster

The cluster will consist of three nodes all running both the server and the agent processes

## Installing the first node 

```bash
curl -sfL https://get.k3s.io | sh -s - server --disable=traefik --disable=servicelb --cluster-init
```