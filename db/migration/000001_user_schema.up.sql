CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL CONSTRAINT usernamechk CHECK (char_length(username) >= 5),
    "email" VARCHAR(255) NOT NULL CONSTRAINT emailchk CHECK (char_length(email) >= 5),
    "hash_password" VARCHAR(255) NOT NULL CONSTRAINT passwordchk CHECK (hash_password <> ''),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ
);
CREATE UNIQUE INDEX idx_users_username ON users(username);
CREATE UNIQUE INDEX idx_users_email ON users(email);