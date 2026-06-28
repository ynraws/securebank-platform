# ELK Stack - SecureBank Platform

## Components
- **Elasticsearch** - Log storage and search
- **Kibana** - Log visualization and dashboards
- **Filebeat** - Log shipping from all pods

## Installation

```bash
# Add elastic helm repo
helm repo add elastic https://helm.elastic.co
helm repo update

# Create namespace
kubectl create namespace logging

# Install Elasticsearch
helm install elasticsearch elastic/elasticsearch \
  --namespace logging \
  --set replicas=1 \
  --set minimumMasterNodes=1 \
  --set resources.requests.cpu=500m \
  --set resources.requests.memory=1Gi \
  --set persistence.enabled=false \
  --set antiAffinity=soft

# Install Kibana
helm install kibana elastic/kibana \
  --namespace logging \
  --set resources.requests.cpu=500m \
  --set resources.requests.memory=1Gi \
  --timeout 10m

# Install Filebeat
helm install filebeat elastic/filebeat \
  --namespace logging \
  --set resources.requests.cpu=100m \
  --set resources.requests.memory=200Mi
```

## Access Kibana

```bash
# Get password
kubectl get secrets --namespace=logging elasticsearch-master-credentials \
  -ojsonpath='{.data.password}' | base64 -d && echo

# Port forward
kubectl port-forward svc/kibana-kibana -n logging 5601:5601
```

Open: http://localhost:5601
- Username: `elastic`
- Password: (from above command)

## Dashboards
- SecureBank Platform - Security & Observability Dashboard
  - Log Volume by Namespace
  - Log Volume by Service
  - Log Count by Container
