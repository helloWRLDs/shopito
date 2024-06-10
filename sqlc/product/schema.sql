CREATE TABLE categories(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(60) NOT NULL
);

CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255),
    img_url VARCHAR(255),
    price INT,
    stock INT,
    category_id INT DEFAULT 0 REFERENCES categories(id) ON DELETE CASCADE
);

