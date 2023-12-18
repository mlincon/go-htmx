CREATE SCHEMA IF NOT EXISTS todo;

CREATE TABLE todo.todos (
    id SERIAL NOT NULL,
    todo VARCHAR(128),
    done BOOLEAN,
    PRIMARY KEY (id)
);

INSERT INTO todo.todos (todo, done)
VALUES ('Hello!', FALSE);
