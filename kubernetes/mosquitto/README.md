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

Create the Mosquitto password Kubernetes secret from the password file
```bash
kubectl create secret generic passwd -n mosquitto --from-file=passwd
```