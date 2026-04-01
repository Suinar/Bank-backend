CREATE SCHEMA IF NOT EXISTS users;
CREATE SCHEMA IF NOT EXISTS accounts;
CREATE SCHEMA IF NOT EXISTS cards;
CREATE SCHEMA IF NOT EXISTS credits;
CREATE SCHEMA IF NOT EXISTS deposits;
CREATE SCHEMA IF NOT EXISTS currencies;

CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    middle_name VARCHAR(50),
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone_number VARCHAR(15) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

CREATE TABLE accounts (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    currency_id UUID NOT NULL,
    name VARCHAR(20) NOT NULL,
    balance BIGINT NOT NULL,
    status SMALLINT NOT NULL CHECK (status IN (0,1,2)),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (currency_id) REFERENCES currenies(id)
);

CREATE TABLE cards (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    account_id UUID NOT NULL,
    number VARCHAR(19) UNIQUE NOT NULL,
    expiry_month SMALLINT NOT NULL,
    expiry_year SMALLINT NOT NULL,
    status SMALLINT NOT NULL CHECK (status IN (0,1,2,3)),

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (currency_id) REFERENCES currenies(id)
);

CREATE TABLE credits (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    currency_id UUID NOT NULL,
    amount UUID NOT NULL,
    interest_rate SMALLINT NOT NULL,
    term_month SMALLINT NOT NULL,
    monthly_payment UUID NOT NULL,
    status SMALLINT NOT NULL CHECK (status IN (0,1,2)),

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (currency_id) REFERENCES currenies(id)
);

CREATE TABLE deposits (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    currency_id UUID NOT NULL,
    amount UUID NOT NULL,
    interest_rate SMALLINT NOT NULL,
    term_month SMALLINT NOT NULL,
    status NOT NULL CHECK (status IN (0,1,2)),

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (currency_id) REFERENCES currenies(id)
);

CREATE TABLE currencies (
    id UUID PRIMARY KEY,
    name VARCHAR UNIQUE NOT NULL,
    symbol CHAR UNIQUE NOT NULL,
    iso_code VARCHAR(4) UNIQUE NOT NULL,
    number_code VARCHAR(3) UNIQUE NOT NULL,
    minor_units VARCHAR(3) NOT NULL
);