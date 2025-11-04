-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "ucodes" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMPTZ,
    "status" SMALLINT DEFAULT 1,
    "user_id" UUID,
    "token" TEXT,
    "email" TEXT,
    "expired_at" TIMESTAMPTZ,
    CONSTRAINT "ucodes_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "ucodes";
-- +goose StatementEnd
