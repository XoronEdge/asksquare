CREATE TABLE "qa_subjects" (
    "id" bigserial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 4),
    "concise_name" VARCHAR(255) NOT NULL CONSTRAINT concise_namechk CHECK (char_length(concise_name) >= 1),
    "version" VARCHAR(255) NOT NULL CONSTRAINT versionchk CHECK (char_length(version) >= 1),    
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ,
    "qa_term_id" bigserial,
    CONSTRAINT  fk_qa_term FOREIGN Key(qa_term_id) REFERENCES qa_terms(id) 
)