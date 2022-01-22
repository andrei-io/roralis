-- +goose Up
CREATE TABLE "posts"(
  "id" serial PRIMARY KEY,
  "user_id" integer NOT NULL REFERENCES "users"("id"),
  "latitude" NUMERIC(10, 7),
  "longitude" NUMERIC(10, 7),
  "title" VARCHAR(50) NOT NULL,
  "description" VARCHAR(750),
  "address" VARCHAR(750),
  "expiry" TIMESTAMPTZ,
  "photo" VARCHAR(2048),
  "region_id" integer NOT NULL REFERENCES "regions"("id"),
  "category_id" integer NOT NULL REFERENCES "categories"("id"),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS "posts";