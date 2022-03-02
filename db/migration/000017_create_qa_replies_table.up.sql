CREATE TABLE qa_replies (
  id BIGSERIAL PRIMARY KEY,
  description VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ,

  qa_comment_id BIGSERIAL REFERENCES qa_comments(id)
);