CREATE TYPE genderType AS ENUM ('Male', 'Female', 'Other');

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL CONSTRAINT usernamechk CHECK (char_length(username) >= 5),
    "email" VARCHAR(255) NOT NULL CONSTRAINT emailchk CHECK (char_length(email) >= 5),
    "first_name" VARCHAR(255)  CONSTRAINT firstnamechk CHECK (char_length(first_name) >= 5),
    "last_name" VARCHAR(255)  CONSTRAINT lastnamechk CHECK (char_length(last_name) >= 5),
    "phone" VARCHAR(255) NOT NULL CONSTRAINT phonechk CHECK (char_length(phone) >= 7),
    "gender" genderType  CONSTRAINT genderchk CHECK(gender in ('Male', 'Female', 'Other')),
    "hash_password" VARCHAR(255) NOT NULL CONSTRAINT passwordchk CHECK (hash_password <> ''),
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now()),
    "updated_at" TIMESTAMPTZ,
    "deleted_at" TIMESTAMPTZ
);