-- +goose Up
CREATE TABLE "categories"(
  "id" serial PRIMARY KEY,
  "text" text NOT NULL UNIQUE,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMPTZ
);

INSERT INTO
  "categories"(text)
VALUES
  ('Categoria 1'),
  ('Categoria 2');

-- +goose Down
DROP TABLE IF EXISTS "categories";