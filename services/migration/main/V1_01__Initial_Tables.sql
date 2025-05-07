-- pgFormatter-ignore

-- uuid extension needed for tables using uuid
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

-- create the divisions, offices, and their relationship
CREATE TABLE IF NOT EXISTS office(
    id varchar(2) NOT NULL
    , code varchar(3) NOT NULL
    , symbol varchar(5) NOT NULL
    , fullname varchar NOT NULL
    , office_type varchar NOT NULL
    , parent_id varchar(2)
    , CONSTRAINT office_pk PRIMARY KEY (id)
    , CONSTRAINT office_unique UNIQUE (code , fullname)
    , CONSTRAINT office_parent_fk FOREIGN KEY (parent_id) REFERENCES office(id)
);

-- load District geometry
CREATE TABLE IF NOT EXISTS office_aor_cw(
    id uuid NOT NULL DEFAULT uuid_generate_v4()
    , office_id varchar(2) NOT NULL
    , geom geometry
    , CONSTRAINT office_aor_cw_pk PRIMARY KEY (id)
    , CONSTRAINT office_id_fk FOREIGN KEY (office_id) REFERENCES office(id)
);

CREATE TABLE IF NOT EXISTS office_aor_reg(
    id uuid NOT NULL DEFAULT uuid_generate_v4()
    , office_id varchar(2) NOT NULL
    , geom geometry
    , CONSTRAINT office_aor_reg_pk PRIMARY KEY (id)
    , CONSTRAINT office_id_fk FOREIGN KEY (office_id) REFERENCES office(id)
);

CREATE TABLE IF NOT EXISTS office_aor_fuds(
    id uuid NOT NULL DEFAULT uuid_generate_v4()
    , office_id varchar(2) NOT NULL
    , geom geometry
    , CONSTRAINT office_aor_fuds_pk PRIMARY KEY (id)
    , CONSTRAINT office_id_fk FOREIGN KEY (office_id) REFERENCES office(id)
);

CREATE TABLE IF NOT EXISTS office_aor_mil(
    id uuid NOT NULL DEFAULT uuid_generate_v4()
    , office_id varchar(2) NOT NULL
    , geom geometry
    , CONSTRAINT office_aor_mil_pk PRIMARY KEY (id)
    , CONSTRAINT office_id_fk FOREIGN KEY (office_id) REFERENCES office(id)
);

