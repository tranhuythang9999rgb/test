-- Active: 1721114895358@@127.0.0.1@5432@sell_products
CREATE table users (
    id BIGINT PRIMARY key,
    user_name VARCHAR(128),
    display_name VARCHAR(255),
    password TEXT,
    avatar VARCHAR(255),
    google_id VARCHAR(128),
    create_time INTEGER,
    update_time INTEGER
);

CREATE Table files(
    id BIGINT PRIMARY KEY,
    any_id BIGINT,
    url VARCHAR(1024)
);

CREATE Table products(
    id BIGINT PRIMARY KEY,
    creator_id BIGINT,
    name VARCHAR(255),
    price DECIMAL(10, 2),
    quantity INTEGER,
    description VARCHAR(1024),
    discount_percent DECIMAL(5, 2),
    status_sell BOOLEAN DEFAULT true,
    is_active BOOLEAN,
    create_time INTEGER,
    update_time INTEGER
)

CREATE Table role(
    id PRIMARY PRIMARY KEY,
    title VARCHAR(128),
    slug VARCHAR(128),
    creator_id BIGINT,
    create_time INTEGER,
    update_time INTEGER
);

CREATE Table permission(
    id BIGINT PRIMARY KEY,
    title VARCHAR(128),
    slug VARCHAR(128),
    active int DEFAULT 1,
);

CREATE Table orders(
    id BIGINT PRIMARY KEY,
    product_id BIGINT,
    creator_id BIGINT,
    quantity INT,
    price DECIMAL(10, 2),
    status BOOLEAN,
    create_time INTEGER,
    update_time INTEGER
);