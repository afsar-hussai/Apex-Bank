
# Apex Bank

A high-performance, scalable, and secure digital banking backend inspired by modern neobanks like Monzo and Razorpay. This project focuses on high throughput, low latency, and ACID-compliant financial transactions.

## 🚀 Overview

Apex Bank is designed to handle the complexities of modern digital finance, featuring a microservices architecture capable of supporting millions of users and high-concurrency transaction environments.

## ✨ Core Features

### Functional
- **Identity & Access:** Secure User Registration with KYC, 2FA, and Role-Based Access Control (Customer, Teller, Admin).
- **Banking Operations:** Multiple account types (Savings, Current, Fixed Deposit) with real-time balance updates.
- **Transactions:** Idempotent fund transfers (NEFT/RTGS/IMPS simulations) and searchable, paginated transaction history.
- **Reporting:** PDF Account Statements and real-time Email/In-app notifications.
- **Security & Control:** Account freeze/unfreeze capabilities, basic fraud detection, and comprehensive Audit Logs.
- **Management:** Centralized Admin Dashboard for system oversight.

### Non-Functional (Performance & Reliability)
- **Performance:** < 500ms P99 latency and 10,000 Transactions Per Second (TPS).
- **Scalability:** Built for 10M total users and 100k concurrent connections.
- **Availability:** 99.99% Uptime with ACID-compliant transaction integrity.
- **Security:** AES-256 at-rest encryption and TLS 1.3 in-transit.
- **Resilience:** Disaster recovery with 15min RTO and 1min RPO.
- **Observability:** Prometheus/Grafana monitoring and Jaeger for distributed tracing.

## 🛠 Tech Stack

- **Backend:** Go (Golang)
- **Database:** [E.g., PostgreSQL] (ACID compliance)
- **Cache:** [E.g., Redis]
- **Infrastructure:** Docker, Kubernetes, Terraform
- **Observability:** Prometheus, Grafana, Jaeger

## 🏗 System Design & Architecture

![HLD Diagram](/Images/Architecture.png)

Apex Bank follows a layered architecture with clear separation of concerns:

**Presentation Layer**
- Customer Web App (Next.js) — account management, transfers, statements
- Admin Dashboard (Next.js) — user management, transaction oversight, fraud alerts

**Business Logic Layer**
- API Gateway — single entry point, handles rate limiting, auth routing, and load balancing
- Auth Service — JWT-based authentication, 2FA, RBAC enforcement
- Account Service — account lifecycle, balance management, ledger operations
- Payment Service — idempotent fund transfers, NEFT/RTGS/IMPS simulation
- Notification Service — async email and in-app alerts via event queue

**Data Layer**
- PostgreSQL — primary store for all financial data (ACID compliant)
- Redis — session store, balance caching, rate-limit counters

> The system is designed so that no request reaches a service without passing
> through the API Gateway. Auth is enforced at the gateway level on every call.

## 📊 Data Flow — Fund Transfer Request

![Data Flow Diagram](/Images/Data_Flow.png)

The following describes the complete lifecycle of a fund transfer request:

1. **Customer Web App** initiates a transfer request with JWT token
2. **API Gateway** receives the request, applies rate limiting
3. **Auth Service** validates the JWT token and user permissions
4. **Account Service** checks sender balance (Redis cache → PostgreSQL fallback)
5. **PostgreSQL** debits the sender account atomically
6. **PostgreSQL** credits the receiver account in the same transaction
7. **Payment Service** records the transaction with idempotency key
8. **Notification Service** dispatches transfer confirmation asynchronously
9. **Customer Web App** receives success response with transaction ID

> All steps 5 and 6 occur within a single PostgreSQL transaction —
> if any step fails, the entire operation is rolled back (ACID guarantee).

## 📁 Project Structure

```
apex-bank/
├── cmd/
│   └── server/          # Application entrypoint
├── internal/
│   ├── auth/            # Authentication & authorization
│   ├── account/         # Account management & ledger
│   ├── payment/         # Fund transfers & idempotency
│   └── notification/    # Async alerts & emails
├── pkg/
│   ├── errors/          # Custom error types
│   ├── logger/          # Structured logging (zap)
│   └── validator/       # Request validation
├── migrations/          # PostgreSQL schema migrations
├── docker-compose.yml
├── Makefile
└── README.md
```



## 📈 Roadmap

- [ ] Core Banking Engine & Ledger
- [ ] Authentication & KYC Service
- [ ] Transaction Processing Pipeline
- [ ] Notification & Reporting Service
- [ ] Fraud Detection & Audit Logging
- [ ] Monitoring & Tracing Integration

---

## ⚙️ Getting Started

### Prerequisites
- Go 1.22+
- Docker & Docker Compose
- PostgreSQL 15+
- Redis 7+

### Local Setup

```bash
# Clone the repository
git clone https://github.com/afsar-hussai/Apex-Bank.git
cd Apex-Bank

# Start infrastructure
docker-compose up -d

# Run database migrations
make migrate-up

# Start the server
make run
```

> Full setup guide and environment variables documented in `docs/setup.md`
> (coming soon as development progresses)


*Note: This is a work-in-progress project inspired by industry leaders in the FinTech space.*
README.md
Displaying README.md.