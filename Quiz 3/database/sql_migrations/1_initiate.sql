-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE categories (
    id            BIGINT NOT NULL PRIMARY KEY,
    name          VARCHAR(256),
    created_at    TIMESTAMP,
    created_by    VARCHAR(256),
    modified_at   TIMESTAMP,
    modified_by   VARCHAR(256)
);

CREATE TABLE users (
    id            BIGINT NOT NULL PRIMARY KEY,
    username      VARCHAR(256) NOT NULL,
    password      VARCHAR(256) NOT NULL,
    created_at    TIMESTAMP,
    created_by    VARCHAR(256),
    modified_at   TIMESTAMP,
    modified_by   VARCHAR(256)
);
CREATE TABLE books (
    id            BIGINT NOT NULL PRIMARY KEY,
    title         VARCHAR(256),
    description   VARCHAR(256),
    image_url     VARCHAR(256),
    release_year  INT,
    price         INT,
    total_page    INT,
    thickness     VARCHAR(256),
    category_id   BIGINT REFERENCES categories(id) ON DELETE SET NULL,
    created_at    TIMESTAMP,
    created_by    VARCHAR(256),
    modified_at   TIMESTAMP,
    modified_by   VARCHAR(256)
);

-- +migrate StatementEnd