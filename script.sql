CREATE DATABASE urls
    WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Table: public.url

-- DROP TABLE IF EXISTS public.url;

CREATE TABLE IF NOT EXISTS public.url
(
    id text COLLATE pg_catalog."default",
    longurl text COLLATE pg_catalog."default",
    shorturl text COLLATE pg_catalog."default"
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.url
    OWNER to postgres;