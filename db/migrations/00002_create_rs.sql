-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "rumah_sakits" (
    "id" bigserial,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    "nama" character varying(100),
    "lower" character varying(100),
    "upper" character varying(100),
    "start" character varying(100),
    "stop" character varying(100),
    "next_schedule" timestamptz,
    PRIMARY KEY ("id")
) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS "rumah_sakits";
