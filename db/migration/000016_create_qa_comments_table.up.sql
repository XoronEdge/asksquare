CREATE TABLE qa_comments (
  id BIGSERIAL PRIMARY KEY,
  description VARCHAR(255) NOT NULL,
  total_replies_count INT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ,

  qa_answer_id BIGSERIAL REFERENCES qa_answers(id)
);