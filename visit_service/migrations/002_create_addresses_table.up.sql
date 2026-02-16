CREATE TABLE IF NOT EXISTS addresses (
    id SERIAL PRIMARY KEY,
    city VARCHAR(100) NOT NULL,
    locality VARCHAR(100),
    region VARCHAR(100) NOT NULL,
    street VARCHAR(100) NOT NULL,
    house_number INTEGER NOT NULL,
    letter VARCHAR(10),
    building INTEGER,
    appartment_number INTEGER,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);