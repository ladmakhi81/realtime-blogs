CREATE TABLE IF NOT EXISTS "_tokens" (
    "id"            BIGSERIAL      PRIMARY KEY,
    "created_at"    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    "access_token"  TEXT           NOT NULL,
    "refresh_token" TEXT           NOT NULL,
    "user_id"       BIGINT         NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY ("user_id") REFERENCES "_users"("id") ON DELETE NO ACTION
);