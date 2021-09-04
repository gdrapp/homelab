# Wireguard

## Deploying

1. Generate the Wireguard server private and public keys
```bash
sudo ctr run --rm ghcr.io/gdrapp/wireguard-container:master wg bash -c 'wg genkey | tee privatekey | wg pubkey && cat privatekey'
```

2. Add the public key to the wg0-publickey configmap generator in kustomization.yaml

3. Apply the K8S Kustomization from the wireguard  directory
```bash
sudo kubectl apply -k kubernetes/wireguard/
```

Apply the scripts ConfigMap
```bash
kubectl apply -n wireguard -f kubernetes/wireguard/configmap.yaml
```

4. Create the K8S secret for the Wireguard private key 
```bash
sudo kubectl -n wireguard create secret generic wg0-privatekey --from-file=privatekey=privatekey
```

## Add a New Client

1. Generate the client and server config
```bash
sudo kubectl exec -n wireguard daemonset/wireguard -it -- /usr/local/bin/wg-client-gen.sh
```

2. Add the generated server config to wg0.conf

2. Apply the K8S Kustomization to redeploy the configmap
```bash
sudo kubectl apply -k kubernetes/wireguard/
```

## Troubleshooting

- Run the "wg" command in the wireguard container to see who's connected
```bash
sudo kubectl exec -n wireguard daemonset/wireguard -it -- wg
```