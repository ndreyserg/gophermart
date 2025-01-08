create table users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    pass_hash VARCHAR(255) NOT NULL,
    UNIQUE(login)
)

create TABLE orders (
	id BIGSERIAL NOT NULL PRIMARY KEY,
   number VARCHAR(255) NOT NULL,
   status VARCHAR(255) NOT NULL,
   accrual DECIMAL(20, 2) NOT NULL,
   user_id BIGINT NOT NULL,
   uploaded_at timestamptz NOT NULL,
   UNIQUE(number)
)


create TABLE accounts (
	id BIGSERIAL NOT NULL PRIMARY KEY,
   balance DECIMAL(20, 2) NOT NULL,
   user_id BIGINT NOT NULL,
   UNIQUE(user_id)
)


CREATE TABLE withdrawals (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	sum DECIMAL(20, 2) NOT NULL,
   account_id BIGINT NOT NULL,
   order_number VARCHAR(256) NOT NULL,
   processed_at timestamptz NOT NULL
)