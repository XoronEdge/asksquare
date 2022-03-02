CREATE TABLE "qa_terms" (
    "id" bigserial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 4),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ,
    "qa_categorie_id" bigserial,
    CONSTRAINT fk_qa_categorie FOREIGN Key(qa_categorie_id) REFERENCES qa_categories(id) 
)