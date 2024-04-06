# longhorn
Persistent Kubernetes pod storage

# Installation
1. Install dependencies
```bash
apt install nfs-common
```

2. Install Longhorn
```bash
kubectl apply -f https://raw.githubusercontent.com/longhorn/longhorn/v1.2.6/deploy/longhorn.yaml
```

3. Wait for the pods to enter a running state
```bash
kubectl -n longhorn-system get pod
```

4. Create a secret containing the basic auth credentials
```bash
USER=<USERNAME_HERE>; PASSWORD=<PASSWORD_HERE>; echo "${USER}:$(openssl passwd -stdin -apr1 <<< ${PASSWORD})" >> auth
kubectl -n longhorn-system create secret generic basic-auth --from-file=auth
rm auth
```

5. Create the ingress and storageclass
```bash
kubectl apply -k kubernetes/longhorn
```