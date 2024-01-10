CREATE TABLE
    otp (
        id SERIAL PRIMARY KEY,
        otp VARCHAR(8) NOT NULL,
        is_used BOOLEAN DEFAULT false,
        created_at TIMESTAMP
    );