BEGIN;

CREATE TABLE IF NOT EXISTS exchanges (
    id UUID PRIMARY KEY,
    "from" TEXT NOT NULL,
    "to" TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    rate DECIMAL NOT NULL
);

ALTER TABLE exchanges DROP CONSTRAINT IF EXISTS exchanges_from_to_key;
ALTER TABLE exchanges ADD CONSTRAINT exchanges_from_to_key UNIQUE ("from", "to");

COMMIT;