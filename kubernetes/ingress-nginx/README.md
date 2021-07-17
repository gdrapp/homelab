# ingress-nginx

## Deploying

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.48.1/deploy/static/provider/baremetal/deploy.yaml
```

To check if the ingress controller pods have started, run the following command:

```bash
kubectl get pods -n ingress-nginx -l app.kubernetes.io/name=ingress-nginx --watch
```

## Tips

Get current ingress-nginx version:

```bash
POD_NAMESPACE=ingress-nginx
POD_NAME=$(kubectl get pods -n $POD_NAMESPACE -l app.kubernetes.io/name=ingress-nginx --field-selector=status.phase=Running -o jsonpath='{.items[0].metadata.name}')

kubectl exec -it $POD_NAME -n $POD_NAMESPACE -- /nginx-ingress-controller --version
```