-- +goose Up
CREATE TABLE "regions"(
  "id" serial PRIMARY KEY,
  "text" text NOT NULL UNIQUE,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS "regions";