CREATE TABLE IF NOT EXISTS visits (
    id SERIAL PRIMARY KEY,
    client_id INTEGER NOT NULL REFERENCES clients(id) ON DELETE RESTRICT,
    master_id INTEGER,
    address_id INTEGER NOT NULL REFERENCES addresses(id) ON DELETE RESTRICT,
    contract_number VARCHAR(100) NOT NULL,
    contract_date DATE NOT NULL,
    scheduled_date DATE,
    scheduled_time DATE,
    equipment_description TEXT,
    assigned_month VARCHAR(20),
    amount NUMERIC(10, 2),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);