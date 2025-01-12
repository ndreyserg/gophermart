CREATE TABLE withdrawals (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	sum DECIMAL(20, 2) NOT NULL,
   account_id BIGINT NOT NULL,
   order_number VARCHAR(256) NOT NULL,
   processed_at timestamptz NOT NULL
)