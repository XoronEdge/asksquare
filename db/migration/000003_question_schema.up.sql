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
    "qa_question_type_id"  bigserial, 
    CONSTRAINT  fk_qa_question_type FOREIGN Key(qa_question_type_id) REFERENCES qa_question_type(id) ON DELETE CASCADE
);