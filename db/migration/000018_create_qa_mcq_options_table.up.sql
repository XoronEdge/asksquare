CREATE TABLE qa_mcq_options (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ,

  qa_question_id BIGSERIAL REFERENCES qa_questions(id)
);