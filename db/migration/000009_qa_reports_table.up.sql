CREATE TABLE "qa_reports" (
    "id" bigserial PRIMARY KEY,
    "reason" VARCHAR(255) NOT NULL CONSTRAINT reasonchk CHECK (char_length(reason) >= 4),
    "description" VARCHAR(255) NOT NULL CONSTRAINT descriptionchk CHECK (char_length(description) >= 5),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ,
    "qa_question_id" bigserial,
    "user_id" bigserial,
    CONSTRAINT fk_qa_question FOREIGN Key(qa_question_id) REFERENCES qa_questions(id) ,
    CONSTRAINT fk_user FOREIGN Key(user_id) REFERENCES users(id) 
);