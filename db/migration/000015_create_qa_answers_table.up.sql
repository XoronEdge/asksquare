CREATE TABLE qa_answers (
  id BIGSERIAL PRIMARY KEY,
  description VARCHAR(255) NOT NULL,
  explaination TEXT,
  total_comments_count INT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ,

  qa_question_id BIGSERIAL REFERENCES qa_questions(id)
);