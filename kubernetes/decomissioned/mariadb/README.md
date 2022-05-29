# MariaDB

MariaDB Server is one of the most popular database servers in the world. It’s made by the original developers of MySQL and guaranteed to stay open source.

## Installation

1. Apply the Kubernetes config
```bash
sudo kubectl apply -k kubernetes/mariadb
```

2. Create the K8S secret for the MariaDB root password
```bash
sudo kubectl -n mariadb create secret generic mariadb-root --from-literal=password=[password]
```
