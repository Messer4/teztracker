-- +migrate Up

CREATE INDEX IF NOT EXISTS ix_operations_public_key_kind
    ON tezos.operations USING btree
    (public_key COLLATE pg_catalog."default", kind COLLATE pg_catalog."default")
      WHERE public_key IS NOT NULL;

-- +migrate Down