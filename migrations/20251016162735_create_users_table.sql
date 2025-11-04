-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    status SMALLINT DEFAULT 1,
    approved_at TIMESTAMPTZ,
    availability TEXT,
    email TEXT UNIQUE,
    username TEXT UNIQUE,
    name VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    password VARCHAR(255),
    domain TEXT UNIQUE,
    avatar TEXT,
    phone_number TEXT,
    country TEXT,
    state TEXT,
    city TEXT,
    address TEXT,
    zip_code TEXT,
    gender TEXT,
    date_of_birth DATE,
    billing_id TEXT,
    type TEXT DEFAULT 'user',
    email_verified_at TIMESTAMPTZ,
    is_two_factor_enabled SMALLINT DEFAULT 0,
    two_factor_secret TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
