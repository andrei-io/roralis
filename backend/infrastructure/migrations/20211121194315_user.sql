-- +goose Up
CREATE TABLE "users"(
  "id" serial PRIMARY KEY,
  "email" VARCHAR(320) NOT NULL UNIQUE,
  "password" VARCHAR(200) NOT NULL,
  -- role is an integer, the higher it is the more permission one has
  "role" SMALLINT NOT NULL,
  "name" VARCHAR(50) NOT NULL,
  "phone" VARCHAR(15),
  "profile" VARCHAR(2048),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" TIMESTAMPTZ
);

-- Passwords are hashed, but will leave it as is for now
-- For production this file has to be modified before running 
INSERT INTO
  "users"("email", "password", "role", "name")
VALUES
  (
    'countryroad.app@gmail.com',
    'asdfasdf',
    100,
    'CountryRoads Team'
  );

-- +goose Down
DROP TABLE IF EXISTS "users";