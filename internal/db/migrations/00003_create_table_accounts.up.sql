create TABLE accounts (
	id BIGSERIAL NOT NULL PRIMARY KEY,
   balance DECIMAL(20, 2) NOT NULL,
   user_id BIGINT NOT NULL,
   UNIQUE(user_id)
);