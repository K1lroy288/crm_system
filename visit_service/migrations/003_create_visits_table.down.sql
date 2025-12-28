CREATE TABLE IF NOT EXISTS visits (
    id SERIAL PRIMARY KEY,
    client_id INTEGER NOT NULL REFERENCES clients(id) ON DELETE RESTRICT,
    master_id INTEGER NOT NULL,
    address_id INTEGER NOT NULL REFERENCES addresses(id) ON DELETE RESTRICT,
    contract_number VARCHAR(100),
    contract_date TIMESTAMPTZ,
    scheduled_date TIMESTAMPTZ,
    equipment_description TEXT,
    assigned_month VARCHAR(20),
    amount NUMERIC(10, 2)
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);