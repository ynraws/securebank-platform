package com.securebank.payment.controller;

import com.securebank.payment.model.Payment;
import com.securebank.payment.service.PaymentService;
import jakarta.validation.Valid;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api/v1/payments")
public class PaymentController {

    @Autowired
    private PaymentService paymentService;

    @GetMapping("/health")
    public ResponseEntity<?> health() {
        return ResponseEntity.ok().body(
            java.util.Map.of("status", "healthy", "service", "payment-svc")
        );
    }

    @PostMapping
    public ResponseEntity<Payment> createPayment(@Valid @RequestBody Payment payment) {
        Payment created = paymentService.createPayment(payment);
        return ResponseEntity.status(HttpStatus.CREATED).body(created);
    }

    @GetMapping("/{id}")
    public ResponseEntity<Payment> getPayment(@PathVariable UUID id) {
        return ResponseEntity.ok(paymentService.getPayment(id));
    }

    @GetMapping
    public ResponseEntity<List<Payment>> getAllPayments() {
        return ResponseEntity.ok(paymentService.getAllPayments());
    }

    @GetMapping("/account/{accountId}")
    public ResponseEntity<List<Payment>> getPaymentsByAccount(@PathVariable UUID accountId) {
        return ResponseEntity.ok(paymentService.getPaymentsByAccount(accountId));
    }

    @PatchMapping("/{id}/status")
    public ResponseEntity<Payment> updateStatus(
            @PathVariable UUID id,
            @RequestParam String status) {
        return ResponseEntity.ok(paymentService.updateStatus(id, status));
    }
}
