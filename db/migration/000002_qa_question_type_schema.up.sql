CREATE TABLE "qa_question_type" (
    "id" bigserial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 3),
    "concise_name" VARCHAR(255) NOT NULL CONSTRAINT concise_namechk CHECK (char_length(concise_name) >= 3),
    "length" VARCHAR(255),
    "is_mcqs" BOOLEAN,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ
);