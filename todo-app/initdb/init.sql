CREATE TABLE todos (
    id SERIAL NOT NULL,
    todo VARCHAR(128),
    done BOOLEAN,
    PRIMARY KEY (id)
);

INSERT INTO todos (todo, done)
VALUES ('Hello!', FALSE);
