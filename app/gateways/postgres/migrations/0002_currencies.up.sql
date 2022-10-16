BEGIN;

CREATE TABLE IF NOT EXISTS currencies (
    id UUID PRIMARY KEY,
    code TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
    usd_rate DECIMAL NOT NULL
);

INSERT INTO currencies (
    id,
    code,
    created_at,
    updated_at,
    usd_rate
) VALUES (
    'ef55fb9b-f80f-4f2c-8212-9b57d7619f30',
    'USD',
    now(),
    now(),
    1
);

COMMIT;