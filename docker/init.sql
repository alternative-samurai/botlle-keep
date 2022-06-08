DROP DATABASE IF EXISTS bottle;
CREATE DATABASE bottle;
\connect bottle

CREATE TABLE public.bottles(
    bottle_id   uuid PRIMARY KEY,
    bottle_name varchar(32) NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL
)