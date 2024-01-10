CREATE TABLE
    ms_menu (
        id SERIAL PRIMARY KEY,
        menu_name VARCHAR(255) NOT NULL,
        title_menu VARCHAR(255) NOT NULL,
        path VARCHAR(255) NOT NULL,
        is_submenu BOOLEAN DEFAULT false,
        parent_name VARCHAR(255) NULL,
        created_at TIMESTAMP,
        updated_at TIMESTAMP
    );