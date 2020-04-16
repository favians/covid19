-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "reports" (
    "id" bigserial,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    "kode" character varying(100),
    "longitude" character varying(100),
    "latitude" character varying(100),
    "kondisi" character varying(100),
    "suhu" character varying(100),
    "demam" character varying(100),
    "rumah_sakit_id" bigint,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("rumah_sakit_id") REFERENCES rumah_sakits(id)

) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS "reports";
