-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "accounts" (
  id bigserial,
  uid uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  messenger_id varchar(255) UNIQUE NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  username varchar(255),
  nickname varchar(255),
  membership smallint,
  role smallint,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  PRIMARY KEY (id)
);

comment on column accounts.role is '帳戶平台身份 1) admin 2) manager 3) normal user';
comment on column accounts.membership is '帳戶學校身份 1) master 2) phd 3) faculty 4) professor 5) alumni';


CREATE TABLE IF NOT EXISTS "wallets" (
  id bigserial,
  uid uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  owner_id bigserial,
  amount decimal(10,2) DEFAULT 0,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  last_paied_at timestamp DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (owner_id) REFERENCES accounts(id) ON DELETE CASCADE
);

comment on column wallets.last_paied_at is '最後付清欠款的時間';

-- +goose Down
DROP TABLE IF EXISTS "wallets";
DROP TABLE IF EXISTS "accounts";