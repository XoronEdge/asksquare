CREATE TABLE "qa_answer_laters" (
    "id" bigserial PRIMARY KEY,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ,
    "qa_question_id" bigserial,
    "user_id" bigserial,
    CONSTRAINT fk_qa_question FOREIGN Key(qa_question_id) REFERENCES qa_questions(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN Key(user_id) REFERENCES users(id) ON DELETE CASCADE
);