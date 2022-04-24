CREATE TABLE qa_cm_answers (
  id BIGSERIAL PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ,

  qa_cm_prompt_id BIGSERIAL REFERENCES qa_cm_prompts(id),
  qa_cm_option_id BIGSERIAL REFERENCES qa_cm_options(id),
  qa_answer_id BIGSERIAL REFERENCES qa_answers(id)
);