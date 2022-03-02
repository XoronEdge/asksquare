CREATE TABLE "qa_hides" (
    "id" bigserial PRIMARY KEY,
    "hide_by_user" INTEGER NOT NULL CONSTRAINT hide_by_userchk CHECK (
        hide_by_user >= 0
        AND hide_by_user <= 1
    ),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ,
    "qa_question_id" bigserial,
    "user_id" bigserial,
    CONSTRAINT fk_qa_question FOREIGN Key(qa_question_id) REFERENCES qa_questions(id) ,
    CONSTRAINT fk_user FOREIGN Key(user_id) REFERENCES users(id) 
);