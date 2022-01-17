-- +goose Up
CREATE TABLE "users"(
    "id" serial PRIMARY KEY,
    "email" text NOT NULL UNIQUE,
    "name" text NOT NULL,
    "password" text NOT NULL,
	"created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"deleted_at" timestamptz 
);

-- +goose Down
DROP TABLE "users";
