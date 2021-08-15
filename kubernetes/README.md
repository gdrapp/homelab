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

## Node extended resources

Start a proxy, so that you can easily send requests to the Kubernetes API server:

```bash
kubectl proxy
```

Note: ~1 in the URLs below is the encoding for the character / in the patch path

### Add an extended resource to a node 

```bash
curl --header "Content-Type: application/json-patch+json" \
--request PATCH \
--data '[{"op": "add", "path": "/status/capacity/gregrapp.net~1rtl2832u", "value": "1"}]' \
http://localhost:8001/api/v1/nodes/<your-node-name>/status
```

## Remove an extended resource from a node

Note: ~1 is the encoding for the character / in the patch path

```bash
curl --header "Content-Type: application/json-patch+json" \
--request PATCH \
--data '[{"op": "remove", "path": "/status/capacity/smarter-devices~1rtlsdr2"}]' \
http://localhost:8001/api/v1/nodes/<your-node-name>/status
```