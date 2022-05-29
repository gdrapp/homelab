# MySQL

MySQL is one of the most popular database servers in the world.

## Installation

1. Apply the Kubernetes config
```bash
sudo kubectl apply -k kubernetes/mysql
```

2. Create the K8S secret for the MySQL root password
```bash
sudo kubectl -n mysql create secret generic mysql-root --from-literal=password=[password]
```