-- +goose Up
CREATE TABLE "one_time_codes" (
  "id" serial PRIMARY KEY,
  "code" VARCHAR(10) NOT NULL,
  "active" BOOLEAN NOT NULL,
  "expire" TIMESTAMPTZ NOT NULL,
  "user_id" integer NOT NULL REFERENCES "users"("id")
);

-- +goose Down
DROP TABLE IF EXISTS "one_time_codes";