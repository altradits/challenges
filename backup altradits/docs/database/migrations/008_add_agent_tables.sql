-- 008_add_agent_tables.sql
-- Agent banking: cash-in / cash-out via local agents, with commission tracking

ALTER TABLE pool_settings ADD COLUMN IF NOT EXISTS agent_fee_pct NUMERIC(5,4) NOT NULL DEFAULT 0.005;

UPDATE pool_settings SET agent_fee_pct = 0.005 WHERE agent_fee_pct IS NULL;

-- Tracks which agent processed a cash-in/cash-out transaction on a customer's wallet
ALTER TABLE transactions ADD COLUMN IF NOT EXISTS agent_id UUID REFERENCES users(id);

CREATE INDEX IF NOT EXISTS idx_transactions_agent ON transactions(agent_id);
