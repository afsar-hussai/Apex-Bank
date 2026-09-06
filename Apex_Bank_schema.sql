CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "full_name" varchar(100) NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "phone" varchar(20) UNIQUE NOT NULL,
  "password_hash" varchar(255) NOT NULL,
  "role" varchar(20) NOT NULL DEFAULT 'customer',
  "kyc_status" varchar(20) NOT NULL DEFAULT 'pending',
  "is_active" boolean NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "user_id" uuid NOT NULL,
  "balance" bigint NOT NULL DEFAULT 0,
  "currency" varchar(3) NOT NULL,
  "account_type" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "sender_account_id" bigint,
  "receiver_account_id" bigint,
  "amount" bigint NOT NULL,
  "currency" varchar(3) NOT NULL,
  "status" varchar NOT NULL,
  "transaction_type" varchar NOT NULL,
  "idempotency_key" varchar(255) UNIQUE,
  "description" text,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "ledger_entries" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "account_id" bigint NOT NULL,
  "transaction_id" uuid NOT NULL,
  "credit_debit" varchar NOT NULL,
  "amount" bigint NOT NULL,
  "balance_after" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "audit_logs" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "actor_id" uuid,
  "action" varchar NOT NULL,
  "table_name" varchar NOT NULL,
  "old_values" jsonb,
  "new_values" jsonb,
  "ip_address" varchar,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "transactions" ADD FOREIGN KEY ("sender_account_id") REFERENCES "accounts" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "transactions" ADD FOREIGN KEY ("receiver_account_id") REFERENCES "accounts" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "ledger_entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "ledger_entries" ADD FOREIGN KEY ("transaction_id") REFERENCES "transactions" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "audit_logs" ADD FOREIGN KEY ("actor_id") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;
