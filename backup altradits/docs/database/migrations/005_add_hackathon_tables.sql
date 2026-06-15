-- 005_add_hackathon_tables.sql

CREATE TABLE IF NOT EXISTS hackathons (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id        UUID REFERENCES events(id),
    title           VARCHAR(255) NOT NULL,
    description     TEXT,
    start_date      TIMESTAMPTZ NOT NULL,
    end_date        TIMESTAMPTZ NOT NULL,
    status          VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS hackathon_students (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hackathon_id    UUID NOT NULL REFERENCES hackathons(id),
    user_id         UUID NOT NULL REFERENCES users(id),
    points          INT NOT NULL DEFAULT 0,
    bio             TEXT,
    skills          TEXT,
    enrolled_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    graduated_at    TIMESTAMPTZ,
    UNIQUE(hackathon_id, user_id)
);

CREATE TABLE IF NOT EXISTS hackathon_submissions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hackathon_id    UUID NOT NULL REFERENCES hackathons(id),
    student_id      UUID NOT NULL REFERENCES users(id),
    title           VARCHAR(255) NOT NULL,
    description     TEXT,
    repo_url        TEXT,
    demo_url        TEXT,
    avg_rating      NUMERIC(3,2) NOT NULL DEFAULT 0,
    review_count    INT NOT NULL DEFAULT 0,
    submitted_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS submission_reviews (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    submission_id   UUID NOT NULL REFERENCES hackathon_submissions(id),
    reviewer_id     UUID NOT NULL REFERENCES users(id),
    rating          INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    comment         TEXT,
    points_awarded  INT NOT NULL DEFAULT 0,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(submission_id, reviewer_id)
);

CREATE TABLE IF NOT EXISTS quiz_questions (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id        UUID REFERENCES events(id),
    question        TEXT NOT NULL,
    option_a        VARCHAR(255) NOT NULL,
    option_b        VARCHAR(255) NOT NULL,
    option_c        VARCHAR(255) NOT NULL,
    option_d        VARCHAR(255) NOT NULL,
    correct_option  CHAR(1) NOT NULL,
    points          INT NOT NULL DEFAULT 10,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS certificates (
    id              UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id         UUID NOT NULL REFERENCES users(id),
    hackathon_id    UUID REFERENCES hackathons(id),
    cert_hash       VARCHAR(64) UNIQUE NOT NULL,
    issued_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
