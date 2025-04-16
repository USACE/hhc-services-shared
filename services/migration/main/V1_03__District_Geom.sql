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

