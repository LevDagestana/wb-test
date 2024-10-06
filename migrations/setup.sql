CREATE TABLE orders (
    order_uid VARCHAR(255) PRIMARY KEY,
    track_number VARCHAR(255),
    entry VARCHAR(50),
    locale VARCHAR(10),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(255),
    delivery_service VARCHAR(255),
    shardkey VARCHAR(10),
    sm_id INT,
    date_created TIMESTAMP,
    oof_shard VARCHAR(10)
);

CREATE TABLE delivery (
    id VARCHAR(255) PRIMARY KEY,
    order_uid VARCHAR(255),
    name VARCHAR(255),
    phone VARCHAR(50),
    zip VARCHAR(20),
    city VARCHAR(100),
    address VARCHAR(255),
    region VARCHAR(100),
    email VARCHAR(255),
   
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid) ON DELETE CASCADE
);

CREATE TABLE payment (
    id VARCHAR(255) PRIMARY KEY,
    order_uid VARCHAR(255),
    transaction VARCHAR(255),
    request_id VARCHAR(50),
    currency VARCHAR(10),
    provider VARCHAR(50),
    amount INT,
    payment_dt BIGINT,
    bank VARCHAR(50),
    delivery_cost INT,
    goods_total INT,
    custom_fee INT,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid) ON DELETE CASCADE
);

CREATE TABLE items (
    id VARCHAR(255) PRIMARY KEY,
    order_uid VARCHAR(255),
    chrt_id INT,
    track_number VARCHAR(255),
    price INT,
    rid VARCHAR(255),
    name VARCHAR(255),
    sale INT,
    size VARCHAR(50),
    total_price INT,
    nm_id INT,
    brand VARCHAR(100),
    status INT,
    FOREIGN KEY (order_uid) REFERENCES orders(order_uid) ON DELETE CASCADE
);
