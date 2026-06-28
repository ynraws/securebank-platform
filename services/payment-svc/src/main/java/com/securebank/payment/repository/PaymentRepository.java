package com.securebank.payment.repository;

import com.securebank.payment.model.Payment;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;

@Repository
public interface PaymentRepository extends JpaRepository<Payment, UUID> {
    List<Payment> findByFromAccountIdOrderByCreatedAtDesc(UUID fromAccountId);
    List<Payment> findByToAccountIdOrderByCreatedAtDesc(UUID toAccountId);
    List<Payment> findByStatus(String status);
}
