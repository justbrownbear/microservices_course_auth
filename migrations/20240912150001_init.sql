-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.users
(
    id bigserial NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    role smallint NOT NULL,
    is_deleted boolean DEFAULT FALSE,
    create_timestamp timestamp without time zone NOT NULL DEFAULT NOW(),
    update_timestamp timestamp without time zone,
    delete_timestamp timestamp without time zone,
    password_hash text NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;

COMMENT ON TABLE public.users
    IS 'Пользователи';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.users;
-- +goose StatementEnd
