-- +goose Up
CREATE TABLE IF NOT EXISTS "products" (
  id bigserial,
  uid uuid UNIQUE NOT NULL DEFAULT uuid_generate_v4(),
  name varchar(255) NOT NULL UNIQUE,
  price decimal(10,2) CHECK (price >= 0) DEFAULT 0,
  cost decimal(10,2) CHECK (cost >= 0) DEFAULT 0,
  quantity integer CHECK (quantity >= 0) DEFAULT 0,
  image_url varchar(255),
  category smallint DEFAULT 1,
  created_at timestamp DEFAULT now(),
  updated_at timestamp DEFAULT now(),
  PRIMARY KEY (id)
);


comment on column products.category is '商品類別 1) snake 2) drink';

CREATE INDEX idx_products_price ON "products" (price);
CREATE INDEX idx_products_cost ON "products" (cost);
CREATE INDEX idx_products_category ON "products" (category);

-- +goose Down
DROP TABLE IF EXISTS "products";