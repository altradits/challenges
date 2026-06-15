-- 006_add_travel_tables.sql

CREATE TABLE IF NOT EXISTS travel_packages (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title           VARCHAR(255) NOT NULL,
    description     TEXT,
    destination     VARCHAR(255) NOT NULL,
    duration_days   INT NOT NULL,
    price_sats      BIGINT NOT NULL,
    capacity        INT NOT NULL DEFAULT 20,
    departure_date  TIMESTAMPTZ,
    provider        VARCHAR(255) NOT NULL DEFAULT 'Gorilla Sats',
    image_url       TEXT,
    status          VARCHAR(20) NOT NULL DEFAULT 'available',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

INSERT INTO travel_packages (title, description, destination, duration_days, price_sats, capacity, provider, status) VALUES
('Bwindi Gorilla Trek', '3-day mountain gorilla experience in Uganda', 'Bwindi, Uganda', 3, 35000000, 8, 'Gorilla Sats', 'available'),
('Masai Mara Safari', '5-day wildebeest migration safari', 'Masai Mara, Kenya', 5, 28000000, 12, 'Gorilla Sats', 'available'),
('Rwanda Volcano Trek', '4-day Volcanoes National Park adventure', 'Musanze, Rwanda', 4, 42000000, 6, 'Gorilla Sats', 'available')
ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS travel_bookings (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    package_id      UUID NOT NULL REFERENCES travel_packages(id),
    user_id         UUID NOT NULL REFERENCES users(id),
    traveler_count  INT NOT NULL DEFAULT 1,
    total_sats      BIGINT NOT NULL,
    status          VARCHAR(20) NOT NULL DEFAULT 'pending',
    notes           TEXT,
    booked_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    confirmed_at    TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_bookings_user ON travel_bookings(user_id);
CREATE INDEX IF NOT EXISTS idx_bookings_status ON travel_bookings(status);
