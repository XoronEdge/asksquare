CREATE TABLE "qa_questions" (
    "id" bigserial PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL CONSTRAINT titlechk CHECK (char_length(title) >= 4),
    "question" VARCHAR(255) NOT NULL CONSTRAINT questionchk CHECK (char_length(question) >= 5),
    "description" VARCHAR(255) CONSTRAINT descriptionchk CHECK (char_length(description) >= 5),
    "total_answers_count" INTEGER,
    "fib_blank_count" INTEGER,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ,
    "qa_question_type_id" bigserial,
    "qa_categorie_id" bigserial,
    "qa_subject_id" bigserial,
    "qa_term_id" bigserial,
    "qa_chapter_id" bigserial,
    "qa_topic_id" bigserial,
    CONSTRAINT fk_qa_question_type FOREIGN Key(qa_question_type_id) REFERENCES qa_question_type(id) ON DELETE CASCADE,
    CONSTRAINT fk_qa_categories FOREIGN Key(qa_categorie_id) REFERENCES qa_categories(id) ON DELETE CASCADE,
    CONSTRAINT fk_qa_subjects FOREIGN Key(qa_subject_id) REFERENCES qa_subjects(id) ON DELETE CASCADE,
    CONSTRAINT fk_qa_terms FOREIGN Key(qa_term_id) REFERENCES qa_terms(id) ON DELETE CASCADE,
    CONSTRAINT fk_qa_chapters FOREIGN Key(qa_chapter_id) REFERENCES qa_chapters(id) ON DELETE CASCADE,
    CONSTRAINT fk_qa_topics FOREIGN Key(qa_topic_id) REFERENCES qa_topics(id) ON DELETE CASCADE
);