-- +goose Up
CREATE TABLE IF NOT EXISTS "orders" (
  id bigserial,
  uid uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  wallet_id int8 NOT NULL,
  status smallint,
  amount decimal(10,2) CHECK (amount >= 0) NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  canceled_at timestamp DEFAULT NULL,
  paied_at timestamp DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE CASCADE
);

comment on column orders.status is '訂單狀態 1) 成功 2) 取消';

CREATE INDEX idx_orders_wallet_id ON "orders" (wallet_id);
CREATE INDEX idx_orders_paied_at ON "orders" (paied_at);

CREATE TABLE IF NOT EXISTS "ordered_products" (
  id bigserial,
  product_id int8 NOT NULL,
  order_id int8 NOT NULL,
  price decimal(10,2) CHECK (price >= 0) NOT NULL DEFAULT 0,
  quantity int CHECK (quantity >= 0) NOT NULL default 0,
  PRIMARY KEY (id),
  FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
  FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);


-- +goose Down
DROP TABLE IF EXISTS "ordered_products";
DROP TABLE IF EXISTS "orders";