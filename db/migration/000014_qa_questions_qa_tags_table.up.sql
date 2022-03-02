CREATE TABLE qa_questions_qa_tags(
    "id" bigserial PRIMARY KEY,
    "qa_tags_id" bigserial,
    "qa_questions_id" bigserial,
    CONSTRAINT fk_qa_tags_id FOREIGN Key(qa_tags_id) REFERENCES qa_tags(id) ,
    CONSTRAINT fk_qa_questions_id FOREIGN Key(qa_questions_id) REFERENCES qa_questions(id) 
)