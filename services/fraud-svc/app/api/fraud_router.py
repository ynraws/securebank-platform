from fastapi import APIRouter
from app.services.fraud_detector import FraudDetector
from pydantic import BaseModel
from typing import Optional

router = APIRouter()
detector = FraudDetector()

class PaymentCheckRequest(BaseModel):
    paymentId: str
    fromAccountId: str
    toAccountId: str
    amount: float
    currency: Optional[str] = "INR"

@router.post("/fraud/check")
def check_fraud(request: PaymentCheckRequest):
    result = detector.analyze(request.dict())
    return {
        "paymentId": request.paymentId,
        "fromAccountId": request.fromAccountId,
        "amount": request.amount,
        "riskScore": result["risk_score"],
        "isFraudulent": result["is_fraudulent"],
        "reason": result["reason"]
    }

@router.get("/fraud/health")
def health():
    return {"status": "healthy", "service": "fraud-svc"}
