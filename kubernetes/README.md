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

### Remove an extended resource from a node

Note: ~1 is the encoding for the character / in the patch path

```bash
curl --header "Content-Type: application/json-patch+json" \
--request PATCH \
--data '[{"op": "remove", "path": "/status/capacity/smarter-devices~1rtlsdr2"}]' \
http://localhost:8001/api/v1/nodes/<your-node-name>/status
```

## Node Maintenance

Drain pods from a node

```bash
kubectl drain server2 --ignore-daemonsets=true
```

Uncordon a node (allow pods to return)

```bash
kubectl uncordon server2
```

## Resize a persistent volume

1. Delete the Statefulset so the pods don't get auto-recreated
```bash
kubectl delete sts --cascade=orphan [statefulset name] -n [namespace]
```

2. Delete pods so that the volume can resize
```bash
kubectl delete pod --cascade=orphan [pod name] -n [namespace]
```

3. Update the PersistentVolumeClaim with new size and apply it
```bash
kubectl apply -f persistentvolumeclaim.yaml -n [namespace]
```

4. Redeploy the Statefulset
```bash
kubectl apply -f statefulset.yaml -n [namespace]
```

## K3S cluster upgrade

Perform steps on each node in the cluster.

1. Drain the node to be upgraded
```bash
kubectl drain server2 --ignore-daemonsets=true --delete-emptydir-data
```

2. Replace the k3s binary at /usr/local/bin with new version
```bash
cd /usr/local/bin
mv k3s k3s.old
mv /tmp/k3s /usr/local/bin
chmod +x k3s
chown root.root k3s
```

3. Restart server or start K3S

4. Uncordon node (allow pods to return)
```bash
kubectl uncordon server2
```