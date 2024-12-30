CREATE TABLE IF NOT EXISTS "_categories" (
    "id"            BIGSERIAL      PRIMARY KEY,
    "created_at"    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    "title"         VARCHAR(255)   NOT NULL UNIQUE,
    "created_by_id" BIGINT         NOT NULL,
    CONSTRAINT fk_created_by FOREIGN KEY ("created_by_id") REFERENCES "_users"("id") ON DELETE NO ACTION
);