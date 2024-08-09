# cert-manager

## Deploying

```bash
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.12.7/cert-manager.yaml
```

## Verifying the install

```bash
kubectl get pods --namespace cert-manager

NAME                                       READY   STATUS    RESTARTS   AGE
cert-manager-5d7f97b46d-259h7              1/1     Running   0          98s
cert-manager-cainjector-69d885bf55-m79lr   1/1     Running   0          98s
cert-manager-webhook-54754dcdfd-dl498      1/1     Running   0          98s
```

## Configure

1. Create Kubernetes secrect for the Cloudflare API token

```
kubectl create secret generic cloudflare-credentials -n cert-manager --from-literal=api-token="abc123"
```

2. Deploy the ClusterIssuer

```
kubectl apply -f clusterissuer.yaml
```

3. Patch the cert-manager deployment to use an external nameserver for DNS01 propagation check

```
kubectl -n cert-manager patch deployment cert-manager --type='json' -p='[{"op":"add","path":"/spec/template/spec/containers/0/args/-","value":"--dns01-recursive-nameservers=1.1.1.1:53"}]'
```