from fastapi import FastAPI
from datetime import datetime
import os
import requests

app = FastAPI(title="Cost Reporter Service", version="1.0.0")

# Simulated cost data (in real world this comes from AWS Cost Explorer)
SERVICE_COSTS = {
    "account-svc": {"daily": 2.50, "monthly": 75.00, "currency": "USD"},
    "payment-svc": {"daily": 4.20, "monthly": 126.00, "currency": "USD"},
    "fraud-svc":   {"daily": 3.10, "monthly": 93.00, "currency": "USD"},
    "notify-svc":  {"daily": 1.80, "monthly": 54.00, "currency": "USD"},
    "api-gateway": {"daily": 1.20, "monthly": 36.00, "currency": "USD"},
    "kafka":       {"daily": 5.00, "monthly": 150.00, "currency": "USD"},
    "postgres":    {"daily": 3.50, "monthly": 105.00, "currency": "USD"},
}

@app.get("/health")
def health():
    return {"status": "healthy", "service": "cost-reporter"}

@app.get("/api/v1/costs/summary")
def get_cost_summary():
    total_daily = sum(s["daily"] for s in SERVICE_COSTS.values())
    total_monthly = sum(s["monthly"] for s in SERVICE_COSTS.values())
    
    return {
        "report_date": datetime.utcnow().isoformat(),
        "environment": os.getenv("ENVIRONMENT", "local"),
        "summary": {
            "total_daily_cost_usd": round(total_daily, 2),
            "total_monthly_cost_usd": round(total_monthly, 2),
            "projected_annual_usd": round(total_monthly * 12, 2)
        },
        "services": SERVICE_COSTS,
        "finops_recommendations": [
            "Use Spot instances for fraud-svc — save ~65%",
            "Enable KEDA scale-to-zero for notify-svc at night",
            "Move Kafka to MSK Serverless — save ~30%",
            "RDS Aurora Serverless for postgres — save ~40% on idle",
            "Karpenter bin-packing can reduce node count by 2"
        ]
    }

@app.get("/api/v1/costs/services/{service_name}")
def get_service_cost(service_name: str):
    if service_name not in SERVICE_COSTS:
        return {"error": f"Service {service_name} not found"}
    
    cost = SERVICE_COSTS[service_name]
    return {
        "service": service_name,
        "daily_cost_usd": cost["daily"],
        "monthly_cost_usd": cost["monthly"],
        "annual_cost_usd": round(cost["monthly"] * 12, 2),
        "cost_center": "banking-platform",
        "team": "securebank-devops"
    }

@app.get("/api/v1/costs/showback")
def get_showback_report():
    return {
        "report_date": datetime.utcnow().isoformat(),
        "showback": {
            "banking-namespace": {
                "services": ["account-svc", "payment-svc", "fraud-svc", "notify-svc"],
                "monthly_cost_usd": 348.00,
                "percentage": 56.3
            },
            "platform-namespace": {
                "services": ["api-gateway", "kafka", "postgres"],
                "monthly_cost_usd": 291.00,
                "percentage": 43.7
            }
        },
        "total_monthly_usd": 639.00,
        "savings_opportunities_usd": 187.50
    }
