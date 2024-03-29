CREATE TABLE
    ms_users (
        id SERIAL PRIMARY KEY,
        user_unique_id VARCHAR(255) UNIQUE NOT NULL,
        username VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL,
        token VARCHAR(255),
        is_active BOOLEAN DEFAULT false,
        user_role INT NULL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP,
        deleted_at TIMESTAMP NULL
    );