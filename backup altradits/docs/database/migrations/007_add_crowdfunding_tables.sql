-- 007_add_crowdfunding_tables.sql

CREATE TABLE IF NOT EXISTS campaigns (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    creator_id      UUID NOT NULL REFERENCES users(id),
    title           VARCHAR(255) NOT NULL,
    description     TEXT,
    goal_sats       BIGINT NOT NULL,
    raised_sats     BIGINT NOT NULL DEFAULT 0,
    category        VARCHAR(50) NOT NULL DEFAULT 'general',
    status          VARCHAR(20) NOT NULL DEFAULT 'active',
    deadline        TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_campaigns_status ON campaigns(status);

CREATE TABLE IF NOT EXISTS donations (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    campaign_id     UUID NOT NULL REFERENCES campaigns(id),
    donor_id        UUID REFERENCES users(id),
    amount_sats     BIGINT NOT NULL,
    is_anonymous    BOOLEAN NOT NULL DEFAULT FALSE,
    message         TEXT,
    donated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_donations_campaign ON donations(campaign_id);

CREATE TABLE IF NOT EXISTS audit_logs (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    actor_id    UUID REFERENCES users(id),
    action      VARCHAR(100) NOT NULL,
    target_id   UUID,
    detail      TEXT,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_audit_logs_actor ON audit_logs(actor_id);
CREATE INDEX IF NOT EXISTS idx_audit_logs_created ON audit_logs(created_at DESC);
