-- name: CreateUser :execresult
INSERT INTO users (email, password, first_name, last_name, is_active)
    VALUES (?, ?, ?, ?, ?);
