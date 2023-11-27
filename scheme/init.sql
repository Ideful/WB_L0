CREATE DATABASE orders_db;
CREATE USER _user WITH PASSWORD '0';
\c orders_db;
SET datestyle = dmy;
CREATE TABLE IF NOT EXISTS orders (
    order_uid VARCHAR(50) PRIMARY KEY,
    track_number VARCHAR(50),
    entry VARCHAR(10),
    locale VARCHAR(5),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(50),
    delivery_service VARCHAR(50),
    shardkey VARCHAR(10),
    sm_id INTEGER,
    date_created TIMESTAMP,
    oof_shard VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS deliveries (
    id SERIAL PRIMARY KEY,
    order_id VARCHAR(50),
    name VARCHAR(100),
    phone VARCHAR(15),
    zip VARCHAR(20),
    city VARCHAR(100),
    address VARCHAR(255),
    region VARCHAR(100),
    email VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS payments (
    id SERIAL PRIMARY KEY,
    order_id VARCHAR(50),
    transaction VARCHAR(50),
    request_id VARCHAR(50),
    currency VARCHAR(3),
    provider VARCHAR(50),
    amount INTEGER,
    payment_dt INTEGER,
    bank VARCHAR(50),
    delivery_cost INTEGER,
    goods_total INTEGER,
    custom_fee INTEGER
);

CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    order_id VARCHAR(50),
    chrt_id INTEGER,
    track_number VARCHAR(50),
    price INTEGER,
    rid VARCHAR(50),
    name VARCHAR(100),
    sale INTEGER,
    size VARCHAR(10),
    total_price INTEGER,
    nm_id INTEGER,
    brand VARCHAR(100),
    status INTEGER
);



GRANT ALL ON TABLE items TO _user;
GRANT ALL ON TABLE payments TO _user;
GRANT ALL ON TABLE deliveries TO _user;
GRANT ALL ON TABLE orders TO _user;
GRANT ALL ON SEQUENCE deliveries_id_seq, payments_id_seq, items_id_seq TO _user;