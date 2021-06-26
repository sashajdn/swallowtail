CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $$ 
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'pager') THEN
		CREATE TYPE pager AS ENUM ('DISCORD', 'EMAIL', 'SMS', 'PHONE', 'UNKNOWN');
	END IF;
END
$$;

DO $$ 
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'exchange') THEN
		CREATE TYPE exchange AS ENUM ('BINANCE', 'FTX');
	END IF;
END
$$;

CREATE TABLE IF NOT EXISTS s_account_accounts (
	account_id uuid DEFAULT uuid_generate_v4(),
	username VARCHAR(50) NOT NULL UNIQUE,
	password VARCHAR(50) NOT NULL,
	discord_id VARCHAR(20) NOT NULL UNIQUE,

	email VARCHAR(50),
	phone_number VARCHAR(20),

	high_priority_pager pager NOT NULL DEFAULT 'UNKNOWN',
	low_priority_pager pager NOT NULL DEFAULT 'UNKNOWN',

	created TIME NOT NULL DEFAULT now(),
	updated TIME NOT NULL DEFAULT now(),
	last_payment_timestamp TIME NOT NULL DEFAULT now(),

	is_admin BOOLEAN DEFAULT FALSE,
	is_futures_member BOOLEAN DEFAULT FALSE,

	PRIMARY KEY(account_id)
);

CREATE INDEX IF NOT EXISTS idx_s_account_accounts_discord_id
ON s_account_accounts(discord_id);

CREATE TABLE IF NOT EXISTS s_account_googlesheets (
	googlesheets_id uuid DEFAULT uuid_generate_v4(),
	spreadsheet_id VARCHAR(200) NOT NULL UNIQUE,
	sheet_id VARCHAR(200) NOT NULL UNIQUE,
	account_id uuid,

	created TIME NOT NULL DEFAULT now(),
	updated TIME NOT NULL DEFAULT now(),

	PRIMARY KEY(googlesheets_id),
	CONSTRAINT fk_account
		FOREIGN KEY(account_id)
			REFERENCES s_account_accounts(account_id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS s_account_exchanges (
	exchange_id uuid DEFAULT uuid_generate_v4(),
	exchange exchange,
	api_key VARCHAR(200),
	secret_key VARCHAR(200),
	account_id uuid,

	created TIME NOT NULL DEFAULT now(),
	updated TIME NOT NULL DEFAULT now(),

	PRIMARY KEY(exchange_id),
	CONSTRAINT fk_account
		FOREIGN KEY(account_id)
			REFERENCES s_account_accounts(account_id) ON DELETE SET NULL
);
