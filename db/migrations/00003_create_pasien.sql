-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE "pasiens" (
    "id" bigserial,
    "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    "nama" character varying(100),
    "no_hp" character varying(100),
    "ttl" character varying(100),
    "jk" character varying(100),
    "kode" character varying(100),
    "status" character varying(100),
    "rumah_sakit_id" bigserial,
    PRIMARY KEY ("id"),
    FOREIGN KEY ("rumah_sakit_id") REFERENCES rumah_sakits(id)

) WITH (oids = false);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS "pasiens";
