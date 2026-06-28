package com.securebank.payment.kafka;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.securebank.payment.model.Payment;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.Map;

@Component
public class PaymentEventProducer {

    private static final String TOPIC = "payment-events";

    @Autowired
    private KafkaTemplate<String, String> kafkaTemplate;

    @Autowired
    private ObjectMapper objectMapper;

    public void sendPaymentEvent(String eventType, Payment payment) {
        try {
            Map<String, Object> event = new HashMap<>();
            event.put("eventType", eventType);
            event.put("paymentId", payment.getId().toString());
            event.put("fromAccountId", payment.getFromAccountId().toString());
            event.put("toAccountId", payment.getToAccountId().toString());
            event.put("amount", payment.getAmount());
            event.put("currency", payment.getCurrency());
            event.put("status", payment.getStatus());
            event.put("timestamp", System.currentTimeMillis());

            String message = objectMapper.writeValueAsString(event);
            kafkaTemplate.send(TOPIC, payment.getId().toString(), message);
            System.out.println("Payment event sent: " + eventType + " for payment " + payment.getId());
        } catch (Exception e) {
            System.err.println("Failed to send payment event: " + e.getMessage());
        }
    }
}
