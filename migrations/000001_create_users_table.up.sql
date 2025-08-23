-- SQLITE Syntax
CREATE TABLE users (
    id TEXT PRIMARY KEY, -- UUID disimpan sebagai TEXT
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Manually create unique index for users.email
CREATE UNIQUE INDEX idx_users_email ON users(email);

