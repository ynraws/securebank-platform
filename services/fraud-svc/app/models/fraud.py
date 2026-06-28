from sqlalchemy import Column, String, Float, DateTime, Boolean
from sqlalchemy.dialects.postgresql import UUID
from sqlalchemy.ext.declarative import declarative_base
from datetime import datetime
import uuid

Base = declarative_base()

class FraudCheck(Base):
    __tablename__ = "fraud_checks"

    id = Column(UUID(as_uuid=True), primary_key=True, default=uuid.uuid4)
    payment_id = Column(String, nullable=False)
    account_id = Column(String, nullable=False)
    amount = Column(Float, nullable=False)
    is_fraudulent = Column(Boolean, default=False)
    risk_score = Column(Float, default=0.0)
    reason = Column(String, nullable=True)
    created_at = Column(DateTime, default=datetime.utcnow)
