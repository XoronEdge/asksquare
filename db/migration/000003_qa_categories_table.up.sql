CREATE TABLE "qa_categories" (
    "id" bigserial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 4),
    "description" VARCHAR(255) NOT NULL CONSTRAINT descriptionchk CHECK (char_length(description) >= 8),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ
)