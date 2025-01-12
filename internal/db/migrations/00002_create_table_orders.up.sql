create TABLE orders (
	id BIGSERIAL NOT NULL PRIMARY KEY,
   number VARCHAR(255) NOT NULL,
   status VARCHAR(255) NOT NULL,
   accrual DECIMAL(20, 2) NOT NULL,
   user_id BIGINT NOT NULL,
   uploaded_at timestamptz NOT NULL,
   UNIQUE(number)
);