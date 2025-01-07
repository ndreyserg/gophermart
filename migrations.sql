create table users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    login VARCHAR(256) NOT NULL,
    pass_hash VARCHAR(256) NOT NULL,
    UNIQUE(login)
)