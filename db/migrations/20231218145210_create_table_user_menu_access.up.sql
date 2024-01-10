CREATE TABLE
    ms_user_menu_access (
        id SERIAL PRIMARY KEY,
        menu_id INT NOT NULL,
        user_id INT NOT NULL,
        role_access_id INT NOT NULL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    );