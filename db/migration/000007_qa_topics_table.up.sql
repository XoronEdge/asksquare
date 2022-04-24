CREATE TABLE "qa_topics" (
    "id" bigserial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 4),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ,
    "qa_chapter_id" bigserial,
    CONSTRAINT fk_qa_chapter FOREIGN Key(qa_chapter_id) REFERENCES qa_chapters(id) 
)