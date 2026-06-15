-- 003_add_profit_access.sql

CREATE TABLE IF NOT EXISTS businesses (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id),
    name            VARCHAR(255) NOT NULL,
    description     TEXT,
    wallet_addr     VARCHAR(255),
    profit_sats     BIGINT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_businesses_user ON businesses(user_id);

CREATE TABLE IF NOT EXISTS pool_settings (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    admin_fee_pct   NUMERIC(5,4) NOT NULL DEFAULT 0.02,
    kes_per_sat     NUMERIC(14,8) NOT NULL DEFAULT 0.0028,
    target_yield    NUMERIC(6,4) NOT NULL DEFAULT 0.05,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_by      UUID REFERENCES users(id)
);

INSERT INTO pool_settings (admin_fee_pct, kes_per_sat, target_yield)
VALUES (0.02, 0.0028, 0.05)
ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS assets (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ticker          VARCHAR(20) NOT NULL,
    name            VARCHAR(255) NOT NULL,
    category        VARCHAR(50) NOT NULL,
    allocation_pct  NUMERIC(5,2) NOT NULL DEFAULT 0,
    current_value   BIGINT NOT NULL DEFAULT 0,
    yield_pct       NUMERIC(6,4) NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO assets (ticker, name, category, allocation_pct, yield_pct) VALUES
('VOO',  'Vanguard S&P 500 ETF',         'equities',     20.00, 0.0180),
('BND',  'Vanguard Total Bond Market',    'bonds',        40.00, 0.0430),
('VMFXX','Vanguard Money Market Fund',    'money_market', 30.00, 0.0530),
('BTC',  'Bitcoin On-chain Reserve',      'cash',         10.00, 0.0000)
ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS profit_logs (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    amount_sats     BIGINT NOT NULL,
    note            TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    created_by      UUID REFERENCES users(id)
);
