CREATE TABLE IF NOT EXISTS songs (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       groups VARCHAR(255) NOT NULL,
                       song VARCHAR(255) NOT NULL,
                       release_date DATE NOT NULL,
                       text TEXT NOT NULL,
                       link VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT NOW() NOT NULL,
                       updated_at TIMESTAMP DEFAULT NOW() NOT NULL
);