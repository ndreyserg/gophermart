create table users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    pass_hash VARCHAR(255) NOT NULL,
    UNIQUE(login)
);