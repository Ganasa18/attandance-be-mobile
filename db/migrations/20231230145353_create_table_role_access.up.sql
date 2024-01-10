CREATE TABLE
    ms_role_access (
        id SERIAL PRIMARY KEY,
        role_id INTEGER NOT NULL,
        "create" BOOLEAN DEFAULT false,
        "read" BOOLEAN DEFAULT false,
        "update" BOOLEAN DEFAULT false,
        "delete" BOOLEAN DEFAULT false,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    );