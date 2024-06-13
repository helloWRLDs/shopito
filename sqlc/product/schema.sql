CREATE TABLE categories(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL
);

CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    img_url VARCHAR(255),
    price INT NOT NULL,
    stock INT NOT NULL,
    category_id BIGINT DEFAULT 0 REFERENCES categories(id) ON DELETE CASCADE
);

