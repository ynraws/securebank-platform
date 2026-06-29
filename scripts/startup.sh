#!/bin/bash
echo "=========================================="
echo "🏦 SecureBank Platform - Startup Script"
echo "=========================================="

GREEN='\033[0;32m'; YELLOW='\033[1;33m'; RED='\033[0;31m'; NC='\033[0m'
log() { echo -e "${GREEN}[✅ OK]${NC} $1"; }
warn() { echo -e "${YELLOW}[⚠️  WAIT]${NC} $1"; }

echo "🔷 Step 1: Checking Kind cluster..."
if kind get clusters 2>/dev/null | grep -q "securebank"; then
  log "Kind cluster already running"
else
  warn "Creating Kind cluster..."
  kind create cluster --name securebank --config ~/securebank/kind-config.yaml
fi

echo "🔷 Step 2: Setting context..."
kubectl cluster-info --context kind-securebank &>/dev/null && log "Context OK"

echo "🔷 Step 3: Waiting for nodes..."
kubectl wait --for=condition=Ready nodes --all --timeout=120s && log "Nodes ready"

echo "🔷 Step 4: Waiting for ArgoCD..."
kubectl wait --for=condition=available deployment/argocd-server -n argocd --timeout=120s && log "ArgoCD ready"

echo "🔷 Step 5: Waiting for Vault..."
kubectl wait --for=condition=ready pod/vault-0 -n vault --timeout=120s && log "Vault ready"

echo "🔷 Step 6: Waiting for SecureBank services..."
kubectl wait --for=condition=available deployments --all -n securebank --timeout=120s && log "Services ready"

echo "🔷 Step 7: Waiting for Monitoring..."
kubectl wait --for=condition=available deployment/prometheus-grafana -n monitoring --timeout=60s && log "Monitoring ready"

echo "🔷 Step 8: Starting port-forwards..."
pkill -f "port-forward" 2>/dev/null
sleep 2
kubectl port-forward svc/argocd-server -n argocd 8080:443 &>/dev/null &
kubectl port-forward svc/prometheus-grafana -n monitoring 3000:80 &>/dev/null &
kubectl port-forward svc/vault -n vault 8200:8200 &>/dev/null &
kubectl port-forward svc/sonarqube-sonarqube -n sonarqube 9000:9000 &>/dev/null &
log "Port forwards started"

echo ""
echo "=========================================="
echo "🎉 SecureBank Platform is UP!"
echo "=========================================="
echo "  ArgoCD:     https://localhost:8080"
echo "  Grafana:    http://localhost:3000"
echo "  Vault:      http://localhost:8200"
echo "  SonarQube:  http://localhost:9000"
echo "=========================================="
