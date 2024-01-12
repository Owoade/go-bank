CREATE TYPE "transaction_status" AS ENUM (
  'credit',
  'debit'
);

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "email" varchar,
  "password" text,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "user_id" int,
  "balance" decimal,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE TABLE "account_numbers" (
  "id" bigserial,
  "account_id" int,
  "account_name" text,
  "account_number" bigint,
  "bank_name" text,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "account_id" int,
  "amount" decimal,
  "type" transaction_status,
  "created_at" timestampz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "transactions" ("account_id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "account_numbers" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
