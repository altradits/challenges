-- 002_add_lock_tables.sql
-- Investment lock tables

CREATE TABLE IF NOT EXISTS investment_locks (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id),
    amount_sats     BIGINT NOT NULL,
    lock_years      INT NOT NULL DEFAULT 5,
    status          VARCHAR(20) NOT NULL DEFAULT 'active',
    profit_sats     BIGINT NOT NULL DEFAULT 0,
    locked_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    matures_at      TIMESTAMPTZ NOT NULL,
    withdrawn_at    TIMESTAMPTZ,
    milestone_1yr   BOOLEAN NOT NULL DEFAULT FALSE,
    milestone_2yr   BOOLEAN NOT NULL DEFAULT FALSE,
    milestone_3yr   BOOLEAN NOT NULL DEFAULT FALSE,
    milestone_5yr   BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS idx_locks_user ON investment_locks(user_id);
CREATE INDEX IF NOT EXISTS idx_locks_status ON investment_locks(status);
CREATE INDEX IF NOT EXISTS idx_locks_matures ON investment_locks(matures_at);

CREATE TABLE IF NOT EXISTS profit_access (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lock_id         UUID NOT NULL REFERENCES investment_locks(id),
    user_id         UUID NOT NULL REFERENCES users(id),
    schedule        VARCHAR(20) NOT NULL DEFAULT 'monthly',
    last_accessed   TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_profit_access_lock ON profit_access(lock_id);
