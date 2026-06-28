# 🏦 SecureBank Platform

> Production-grade **DevSecOps Banking Platform** built with microservices architecture, containerized with Docker, orchestrated on Kubernetes (EKS), and deployed via GitOps with ArgoCD.

![Platform](https://img.shields.io/badge/Platform-AWS%20EKS-orange?logo=amazon-aws)
![GitOps](https://img.shields.io/badge/GitOps-ArgoCD-blue?logo=argo)
![CI/CD](https://img.shields.io/badge/CI%2FCD-GitHub%20Actions-black?logo=github-actions)
![IaC](https://img.shields.io/badge/IaC-Terraform-purple?logo=terraform)
![Monitoring](https://img.shields.io/badge/Monitoring-Prometheus%20%7C%20Grafana-orange?logo=grafana)
![Security](https://img.shields.io/badge/Security-Vault%20%7C%20Trivy-red?logo=vault)

---

## 🏗️ Architecture Overview

```
Client ──▶ [API Gateway :8080]
                │
    ┌───────────┼───────────┐
    ▼           ▼           ▼
[account-svc] [payment-svc] [fraud-svc]
  :8081 (Go)  :8082 (Java) :8083 (Py)
                │
          [notify-svc]
          :8084 (Node)

       [cost-reporter :8085]
```

---

## 🚀 Microservices

| Service | Language | Port | Responsibility |
|---|---|---|---|
| `api-gateway` | Go | 8080 | Route & authenticate all incoming requests |
| `account-svc` | Go | 8081 | Account management & balance operations |
| `payment-svc` | Java Spring Boot | 8082 | Payment processing & transaction handling |
| `fraud-svc` | Python FastAPI | 8083 | Real-time fraud detection & risk scoring |
| `notify-svc` | Node.js | 8084 | Email/SMS/push notification delivery |
| `cost-reporter` | Python | 8085 | FinOps — AWS cost reporting & optimization |

---

## 🛠️ Tech Stack

| Layer | Technology |
|---|---|
| **Containerization** | Docker, Docker Compose |
| **Orchestration** | Kubernetes (AWS EKS) |
| **IaC** | Terraform |
| **CI/CD** | GitHub Actions |
| **GitOps** | ArgoCD |
| **Monitoring** | Prometheus + Grafana |
| **Logging** | ELK Stack |
| **Secrets** | HashiCorp Vault |
| **Security Scanning** | Trivy |
| **FinOps** | Kubecost, AWS Cost Explorer |

---

## 📁 Project Structure

```
securebank-platform/
├── services/
│   ├── api-gateway/        # Go — reverse proxy & JWT auth
│   ├── account-svc/        # Go — account & balance APIs
│   ├── payment-svc/        # Java Spring Boot — payments
│   ├── fraud-svc/          # Python FastAPI — fraud detection
│   ├── notify-svc/         # Node.js — notifications
│   └── cost-reporter/      # Python — FinOps cost reporting
├── infra/
│   ├── terraform/          # AWS EKS, VPC, IAM
│   ├── helm/               # Helm charts per service
│   └── argocd/             # ArgoCD Application manifests
├── .github/
│   └── workflows/          # GitHub Actions CI/CD pipelines
├── monitoring/
│   ├── prometheus/
│   └── grafana/
├── docker-compose.yml
└── README.md
```

---

## ⚡ Quick Start (Local)

```bash
git clone https://github.com/ynraws/securebank-platform.git
cd securebank-platform
docker compose up -d
docker compose ps
```

### Verify Services

```bash
curl http://localhost:8080/health
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
curl http://localhost:8084/health
curl http://localhost:8085/health
```

---

## 🔄 CI/CD Pipeline

1. **Lint & Test** — per-service unit tests
2. **Security Scan** — Trivy container image scanning
3. **Docker Build & Push** — images pushed to ECR
4. **ArgoCD Sync** — GitOps deployment to EKS

---

## 📊 Observability

- **Prometheus + Grafana** — metrics & dashboards
- **ELK Stack** — centralized logging
- **Jaeger** — distributed tracing

---

## 💰 FinOps

- AWS Cost Explorer API integration
- Kubecost for per-namespace cost allocation
- Daily cost reports per microservice team

---

## 🔐 Security

- HashiCorp Vault — runtime secrets injection
- Trivy — vulnerability scanning in CI
- RBAC — least-privilege service accounts
- Network Policies — zero-trust pod communication
- OPA Gatekeeper — policy enforcement

---

## 🗺️ Roadmap

- [x] Microservices implementation (6 services)
- [x] Docker Compose local setup
- [x] GitHub Actions CI/CD pipeline
- [x] Kind cluster local deployment
- [x] ArgoCD GitOps setup
- [x] Prometheus + Grafana monitoring
- [x] Terraform EKS Infrastructure as Code
- [ ] AWS EKS live deployment (requires AWS credentials)
- [ ] Vault secrets integration
- [ ] Istio service mesh
- [ ] FinOps dashboard (Kubecost)

---

## 👨‍💻 Author

**Narayan Reddy** — Senior DevOps/Platform Engineer
[GitHub](https://github.com/ynraws)

> Built to demonstrate HSBC-grade production DevSecOps practices.
