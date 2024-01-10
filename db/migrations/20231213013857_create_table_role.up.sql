CREATE TABLE
    ms_role (
        id SERIAL PRIMARY KEY,
        role_name VARCHAR(255) NOT NULL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    );