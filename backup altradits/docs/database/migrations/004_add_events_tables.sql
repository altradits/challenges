-- 004_add_events_tables.sql

CREATE TABLE IF NOT EXISTS events (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    organizer_id    UUID NOT NULL REFERENCES users(id),
    title           VARCHAR(255) NOT NULL,
    description     TEXT,
    venue           VARCHAR(255),
    event_date      TIMESTAMPTZ NOT NULL,
    capacity        INT NOT NULL DEFAULT 100,
    ticket_sats     BIGINT NOT NULL DEFAULT 0,
    status          VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_events_organizer ON events(organizer_id);
CREATE INDEX IF NOT EXISTS idx_events_status ON events(status);
CREATE INDEX IF NOT EXISTS idx_events_date ON events(event_date);

CREATE TABLE IF NOT EXISTS event_registrations (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id        UUID NOT NULL REFERENCES events(id),
    user_id         UUID NOT NULL REFERENCES users(id),
    paid_sats       BIGINT NOT NULL DEFAULT 0,
    status          VARCHAR(20) NOT NULL DEFAULT 'registered',
    registered_at   TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(event_id, user_id)
);

CREATE TABLE IF NOT EXISTS event_materials (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id        UUID NOT NULL REFERENCES events(id),
    title           VARCHAR(255) NOT NULL,
    url             TEXT,
    material_type   VARCHAR(30) NOT NULL DEFAULT 'link',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS event_checkins (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id        UUID NOT NULL REFERENCES events(id),
    user_id         UUID NOT NULL REFERENCES users(id),
    check_in_date   DATE NOT NULL DEFAULT CURRENT_DATE,
    scanned_by      UUID REFERENCES users(id),
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(event_id, user_id, check_in_date)
);

CREATE TABLE IF NOT EXISTS event_messages (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id        UUID NOT NULL REFERENCES events(id),
    sender_id       UUID NOT NULL REFERENCES users(id),
    message         TEXT NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
