-- +migrate Up

CREATE INDEX IF NOT EXISTS ix_operations_delegate_kind
    ON tezos.operations USING btree
    (delegate COLLATE pg_catalog."default", kind COLLATE pg_catalog."default")
      WHERE delegate IS NOT NULL;

-- +migrate Down