create database gotask owner postgres;
\connect gotask

CREATE TABLE users
(
    id SERIAL,
    name TEXT NOT NULL,
    role smallint NOT NULL DEFAULT 0,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

CREATE TABLE tasks
(
    id SERIAL,
    title TEXT,
    description TEXT,
    done boolean,
    owner_id bigint,
    CONSTRAINT tasks_pkey PRIMARY KEY (id),
    FOREIGN KEY(owner_id) REFERENCES users (id) ON DELETE CASCADE
);


INSERT INTO users(name, role) VALUES
 ('test user', 1);

INSERT INTO tasks(title, description, done, owner_id) VALUES
 ('test task', 'test description', 'false', 1);

INSERT INTO tasks(title, description, done, owner_id) VALUES
 ('test task2', 'test description2', 'false', 1);
