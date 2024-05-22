# Eclipse Mosquitto

Eclipse Mosquitto is an open source message broker which implements MQTT version 5, 3.1.1 and 3.1

## Installation

Apply the Kubernetes config
```bash
kubectl apply -k kubernetes/mosquitto
```

Create the Mosquitto password file
```bash
mosquitto_passwd passwd [username]
```
or
```bash
kubectl -n mosquitto exec -it [mosquitto pod name] -- mosquitto_passwd
```

Create the Mosquitto password Kubernetes secret from the password file
```bash
kubectl create secret generic passwd -n mosquitto --from-file=passwd
```

Get existing Mosquitto password file secret contents
```bash
kubectl -n mosquitto get secret passwd --template={{.data.passwd}} | base64 -d
```

Update existing Mosquitto password file secret
```bash
echo "[secret contents]" | base64 -w0
```