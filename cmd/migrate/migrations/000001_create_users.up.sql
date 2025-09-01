CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE users (
    id bigserial PRIMARY KEY,
    email citext NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    password bytea NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
