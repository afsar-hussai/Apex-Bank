# Apex Bank — Requirements Document

## Functional Requirements
- User Registration with KYC
- Secure Login — Email/Password + 2FA
- Role-Based Access — Customer, Teller, Admin
- Multiple Account Types — Savings, Current, FD
- Real-Time Balance Updates
- Account Statement — PDF Download
- Fund Transfer — NEFT/RTGS/IMPS Simulation
- Idempotent Transactions — Double Payment Prevention
- Transaction History — Searchable, Paginated
- Real-Time Notifications — Email + In-App
- Basic Fraud Detection
- Account Freeze/Unfreeze
- Admin Dashboard
- Audit Logs

## Non-Functional Requirements
- Transaction Latency Under 500ms (P99)
- Scale — 10M Users, 100k Concurrent
- Throughput — 10,000 TPS
- 99.99% Uptime
- ACID Compliant Transactions
- Disaster Recovery — RTO 15min, RPO 1min
- AES-256 Encryption + TLS 1.3
- Rate Limiting + DDoS Protection
- Monitoring — Prometheus + Grafana
- Distributed Tracing — Jaeger

---
Inspired by: Monzo, N26, Razorpay
