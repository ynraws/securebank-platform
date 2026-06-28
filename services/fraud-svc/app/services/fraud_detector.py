from datetime import datetime
from typing import Dict, Any

class FraudDetector:
    
    # Rules for fraud detection
    HIGH_AMOUNT_THRESHOLD = 100000.0   # 1 Lakh INR
    MEDIUM_AMOUNT_THRESHOLD = 50000.0  # 50K INR
    
    def analyze(self, payment_data: Dict[str, Any]) -> Dict[str, Any]:
        risk_score = 0.0
        reasons = []
        
        amount = float(payment_data.get("amount", 0))
        from_account = payment_data.get("fromAccountId", "")
        
        # Rule 1 — High amount check
        if amount >= self.HIGH_AMOUNT_THRESHOLD:
            risk_score += 60.0
            reasons.append(f"High amount: ₹{amount}")
        elif amount >= self.MEDIUM_AMOUNT_THRESHOLD:
            risk_score += 30.0
            reasons.append(f"Medium-high amount: ₹{amount}")

        # Rule 2 — Odd hours check (midnight to 5am IST = 18:30 to 23:30 UTC)
        hour = datetime.utcnow().hour
        if 18 <= hour <= 23:
            risk_score += 20.0
            reasons.append("Transaction during odd hours")

        # Rule 3 — Round number check (suspicious)
        if amount % 10000 == 0 and amount > 0:
            risk_score += 10.0
            reasons.append("Suspiciously round amount")

        # Determine fraud
        is_fraudulent = risk_score >= 70.0
        
        return {
            "payment_id": payment_data.get("paymentId"),
            "account_id": from_account,
            "amount": amount,
            "risk_score": risk_score,
            "is_fraudulent": is_fraudulent,
            "reason": ", ".join(reasons) if reasons else "No issues detected"
        }
