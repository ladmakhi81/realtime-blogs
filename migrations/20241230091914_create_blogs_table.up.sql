CREATE TABLE IF NOT EXISTS "_blogs" (
    "id"            BIGSERIAL      PRIMARY KEY,
    "created_at"    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    "title"         VARCHAR(255)   NOT NULL,
    "content"       TEXT           NOT NULL,
    "created_by_id" BIGINT         NOT NULL,
    "category_id"   BIGINT         NOT NULL,
    "tags"          TEXT           NOT NULL,
    CONSTRAINT "fk_created_by_id" FOREIGN KEY ("created_by_id") REFERENCES "_users"("id"), 
    CONSTRAINT "fk_category_id"   FOREIGN KEY ("category_id") REFERENCES "_categories"("id") 
);