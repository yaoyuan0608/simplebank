CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account" (
  "id" BIGSERIAL PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "transfer" (
  "id" BIGSERIAL PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "create_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "account" ("owner");

CREATE UNIQUE INDEX ON "account" ("owner", "currency");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfer" ("from_account_id");

CREATE INDEX ON "transfer" ("to_account_id");

CREATE INDEX ON "transfer" ("from_account_id", "to_account_id");

COMMENT ON COLUMN "entries"."amount" IS 'can be +/-';

COMMENT ON COLUMN "transfer"."amount" IS 'must be +';

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("from_account_id") REFERENCES "account" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_account_id") REFERENCES "account" ("id");

ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");
