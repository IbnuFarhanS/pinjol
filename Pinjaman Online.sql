CREATE TABLE "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "username" VARCHAR(255) UNIQUE NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "nik" VARCHAR(255) NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "alamat" VARCHAR(255) NOT NULL,
  "phone_number" VARCHAR(20) NOT NULL,
  "limit" decimal(10,2) NOT NULL,
  "id_role" INT,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "products" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "amount" decimal(10,2) NOT NULL,
  "installment" INT NOT NULL,
  "bunga" decimal(10,2) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "roles" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "transactions" (
  "id" BIGSERIAL PRIMARY KEY,
  "id_user" INT,
  "id_requirement" INT,
  "id_product" INT,
  "status" bool,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()',
  "due_date" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "accept_statuss" (
  "id" BIGSERIAL PRIMARY KEY,
  "id_transaction" INT,
  "status" bool NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

CREATE TABLE "payments" (
  "id" BIGSERIAL PRIMARY KEY,
  "id_transaction" INT,
  "payment_amount" decimal(10,2) NOT NULL,
  "payment_date" timestamptz NOT NULL DEFAULT 'NOW()',
  "id_payment_method" INT
);

CREATE TABLE "payment_methods" (
  "id" BIGSERIAL PRIMARY KEY,
  "name" VARCHAR(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'NOW()'
);

ALTER TABLE "users" ADD FOREIGN KEY ("id_role") REFERENCES "roles" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("id_product") REFERENCES "products" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("id_user") REFERENCES "users" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("id_transaction") REFERENCES "transactions" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("id_payment_method") REFERENCES "payment_methods" ("id");

ALTER TABLE "accept_statuss" ADD FOREIGN KEY ("id_transaction") REFERENCES "transactions" ("id");
