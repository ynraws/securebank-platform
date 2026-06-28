from kafka import KafkaConsumer
from app.services.fraud_detector import FraudDetector
import json
import os

def start_kafka_consumer():
    bootstrap_servers = os.getenv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092")
    
    try:
        consumer = KafkaConsumer(
            "payment-events",
            bootstrap_servers=bootstrap_servers,
            auto_offset_reset="earliest",
            enable_auto_commit=True,
            group_id="fraud-detection-group",
            value_deserializer=lambda m: json.loads(m.decode("utf-8")),
        )
        
        detector = FraudDetector()
        print(f"Kafka consumer connected to {bootstrap_servers}")
        
        for message in consumer:
            try:
                payment = message.value
                print(f"Received payment event: {payment.get('paymentId')}")
                
                result = detector.analyze(payment)
                
                if result["is_fraudulent"]:
                    print(f"🚨 FRAUD DETECTED! Payment {result['payment_id']} - Score: {result['risk_score']} - Reason: {result['reason']}")
                else:
                    print(f"✅ Payment {result['payment_id']} cleared - Score: {result['risk_score']}")
                    
            except Exception as e:
                print(f"Error processing message: {e}")
                
    except Exception as e:
        print(f"Kafka consumer error: {e}")
