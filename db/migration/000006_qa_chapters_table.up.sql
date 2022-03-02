CREATE TABLE "qa_chapters" (
    "id" bigserial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 4),
    "number" INTEGER NOT NULL CONSTRAINT numberchk CHECK (number >= 1),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ,
    "qa_subject_id" bigserial,
    CONSTRAINT fk_qa_subject FOREIGN Key(qa_subject_id) REFERENCES qa_subjects(id) 
)