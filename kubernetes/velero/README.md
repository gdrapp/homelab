# velero

Velero is an open source tool to safely backup and restore, perform disaster recovery, and migrate Kubernetes cluster resources and persistent volumes.

# Installation

1. Download Velero from velero.io

2. Untar Velero and copy binary to /usr/local/bin
```bash
tar zxfv velero-v1.6.3-linux-amd64.tar.gz
sudo cp velero /usr/local/bin
```

3. Create file containing AWS secrets to access S3 bucket storing backups
```bash
cat << EOF > credentials-velero
[default]
aws_access_key_id=xxx
aws_secret_access_key=xxx
```

4. Install Velero into Kubernetes
```bash
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
velero install \
--provider aws \
--plugins velero/velero-plugin-for-aws:v1.2.1 \
--bucket [s3-bucket-name] \
--backup-location-config region=us-west-2 \
--snapshot-location-config region=us-west-2 \
--secret-file ./credentials-velero
```

# Configuration

## Create a backup schedule
```bash
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
velero create schedule daily-no-volumes --schedule="0 4 * * *" --snapshot-volumes=false
```

## View schedules
```bash
export KUBECONFIG=/etc/rancher/k3s/k3s.yaml
velero get schedule
```