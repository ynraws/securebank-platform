from fastapi import FastAPI
from app.api import fraud_router
from app.kafka.consumer import start_kafka_consumer
import asyncio
import threading

app = FastAPI(title="Fraud Detection Service", version="1.0.0")

app.include_router(fraud_router.router, prefix="/api/v1")

@app.get("/health")
def health():
    return {"status": "healthy", "service": "fraud-svc"}

@app.on_event("startup")
async def startup_event():
    # Start Kafka consumer in background thread
    thread = threading.Thread(target=start_kafka_consumer, daemon=True)
    thread.start()
    print("Fraud service started - Kafka consumer running")
