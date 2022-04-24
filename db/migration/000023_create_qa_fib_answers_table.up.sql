CREATE TABLE qa_fib_answers (
  id BIGSERIAL PRIMARY KEY,
  answer VARCHAR(50) NOT NULL,
  option_no INT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ,

  qa_answer_id BIGSERIAL REFERENCES qa_answers(id)
);