CREATE TABLE IF NOT EXISTS "_users" (
    "id"            BIGSERIAL      PRIMARY KEY,
    "created_at"    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    "email"         VARCHAR(255)   UNIQUE,
    "password"      VARCHAR(255)   NOT NULL,
    "profile_url"   VARCHAR(255)   DEFAULT '',
    "first_name"    VARCHAR(255)   DEFAULT '',
    "last_name"     VARCHAR(255)   DEFAULT ''
);