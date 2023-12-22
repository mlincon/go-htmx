CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; --noqa: RF05

CREATE SCHEMA IF NOT EXISTS form;

DROP TABLE IF EXISTS form.users;
CREATE TABLE form.users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(), --noqa: CP03
    name TEXT,
    created_ts TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS form.comments;
CREATE TABLE form.comments (
    comment_id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id UUID, -- CHAR(36)
    product TEXT,
    category TEXT,
    comment TEXT,
    created_ts TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
ALTER TABLE form.comments
ADD CONSTRAINT fk_comment_user FOREIGN KEY (
    user_id
) REFERENCES form.users (user_id);
