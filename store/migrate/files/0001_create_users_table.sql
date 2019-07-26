-- name: create-table-users

CREATE TABLE users
(
    user_name     VARCHAR(50) PRIMARY KEY,
    user_id       VARCHAR(250),
    user_password VARCHAR
);

-- name: create-index-user-id

CREATE INDEX ix_users_id ON users (user_id);

-- name: insert-default-user

INSERT INTO users (user_name, user_id, user_password)
VALUES ('admin', 'admin', 'admin');