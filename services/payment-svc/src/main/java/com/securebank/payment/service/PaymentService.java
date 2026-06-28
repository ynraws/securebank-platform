package com.securebank.payment.service;

import com.securebank.payment.kafka.PaymentEventProducer;
import com.securebank.payment.model.Payment;
import com.securebank.payment.repository.PaymentRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.UUID;

@Service
public class PaymentService {

    @Autowired
    private PaymentRepository paymentRepository;

    @Autowired
    private PaymentEventProducer eventProducer;

    @Transactional
    public Payment createPayment(Payment payment) {
        payment.setStatus("PENDING");
        Payment saved = paymentRepository.save(payment);
        // Send event to Kafka
        eventProducer.sendPaymentEvent("PAYMENT_INITIATED", saved);
        return saved;
    }

    public Payment getPayment(UUID id) {
        return paymentRepository.findById(id)
            .orElseThrow(() -> new RuntimeException("Payment not found: " + id));
    }

    public List<Payment> getPaymentsByAccount(UUID accountId) {
        return paymentRepository.findByFromAccountIdOrderByCreatedAtDesc(accountId);
    }

    @Transactional
    public Payment updateStatus(UUID id, String status) {
        Payment payment = getPayment(id);
        payment.setStatus(status);
        Payment updated = paymentRepository.save(payment);
        eventProducer.sendPaymentEvent("PAYMENT_" + status, updated);
        return updated;
    }

    public List<Payment> getAllPayments() {
        return paymentRepository.findAll();
    }
}
