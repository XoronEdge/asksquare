CREATE TABLE qa_mcq_answer_options (
  id BIGSERIAL PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ,

  qa_mcq_option_id BIGSERIAL REFERENCES qa_mcq_options(id),
  qa_answer_id BIGSERIAL REFERENCES qa_answers(id)
);