# Apache Airflow

Apache Airflow (or simply Airflow) is a platform to programmatically author, schedule, and monitor workflows.

## Installation

Apply the Kubernetes config
```bash
kubectl apply -k kubernetes/airflow
```

Create the fernet key
```bash
kubectl create secret generic airflow-fernet-key -n airflow --from-literal fernet-key=<RANDOM STRING>
```

Create the web server secret key
```bash
kubectl create secret generic airflow-webserver-secret-key -n airflow --from-literal webserver-secret-key=<RANDOM STRING>
```

Create the Airflow database connection string
```bash
kubectl create secret generic airflow-metadata -n airflow --from-literal connection='mysql+mysqldb://<DB_USER>:<DB_PASSWORD>@<DB_HOST>/airflow'
```