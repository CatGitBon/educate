CREATE SCHEMA IF NOT EXISTS auth;

CREATE TABLE IF NOT EXISTS auth.users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO auth.users (username, email, password) VALUES
    ('testuser1', 'test1@example.com', 'password1'),
    ('testuser2', 'test2@example.com', 'password2');

CREATE INDEX IF NOT EXISTS idx_users_username ON auth.users (username);