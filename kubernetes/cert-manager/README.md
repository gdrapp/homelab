# cert-manager

## Deploying

```bash
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml
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

1. Create Kubernetes secrect for the AWS secret access key

```
kubectl create secret generic route53-credentials -n cert-manager --from-literal=secret-access-key="abc123"
```

2. Deploy the ClusterIssuer

```
kubectl apply -f clusterissuer.yaml
```