-- load District geometry
CREATE TABLE IF NOT EXISTS office_aor(
    id uuid NOT NULL DEFAULT uuid_generate_v4()
    , office_id varchar(2) NOT NULL
    , mission varchar NOT NULL
    , geom geometry
    , CONSTRAINT office_aor_pk PRIMARY KEY (id)
    , CONSTRAINT office_id_fk FOREIGN KEY (office_id) REFERENCES office(id)
);

