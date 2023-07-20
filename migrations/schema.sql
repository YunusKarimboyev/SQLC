CREATE TABLE authors (
  id   SERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW(),
  deleted_at timestamp
);