-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "admins" (
    "id" bigserial,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    "name" character varying(100),
    "username" character varying(100),
    "password" character varying(100),
    "instansi" character varying(100),
    PRIMARY KEY ("id")

) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS "admins";
