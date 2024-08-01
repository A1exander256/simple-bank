CREATE TABLE users(
    guid            UUID PRIMARY KEY    NOT NULL,
    username        VARCHAR(50) UNIQUE  NOT NULL,
    hashed_password VARCHAR(100)        NOT NULL,
    email           VARCHAR(100)        NOT NULL,
    created_at      TIMESTAMPTZ         NOT NULL DEFAULT now(),
    updated_at      TIMESTAMPTZ         NOT NULL DEFAULT now()
);