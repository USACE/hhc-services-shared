ARG ALPINE_TAG=3.22
#
# -------- Stage 1 --------
# -------- Builder --------
FROM alpine:${ALPINE_TAG} AS builder

# version of the archive for the geometry migration
ARG GEOM_VERSION=v0.1.1

RUN apk add --no-cache postgis

ENV TIGER=/tiger \
    TIGER_OUT=/tmp

WORKDIR ${TIGER}

# add the release asset for the tagged version defined in GEOM_VERSION
# tiger data for counties
ADD --checksum=sha256:89535e9a798e4feca0a0e786041274b4591b61186a0f9a62b8f478e1d70771e2 https://github.com/USACE/hhc-services-shared/releases/download/${GEOM_VERSION}/tl_2024_us_county.tar.gz ${TIGER}
# tiger data for states
ADD --checksum=sha256:c5e25ab502688993cc509721011da08b104602bc0e2eec8bb5e2da9c72955ec2 https://github.com/USACE/hhc-services-shared/releases/download/${GEOM_VERSION}/tl_2024_us_state.tar.gz ${TIGER}
# HUC6
ADD --checksum=sha256:dc9c7e0044eb9686988e0cff0608fcce5e5658a556b21c8b982ba91a01178910 https://github.com/USACE/hhc-services-shared/releases/download/${GEOM_VERSION}/wbdhu6.tar.gz ${TIGER}
# HUC8
# ADD --checksum=sha256:84d9dc1caf680506cb79938ec4ebc77558050a4027917f9a4e2aa2f64ac7b537 https://github.com/USACE/hhc-services-shared/releases/download/${GEOM_VERSION}/wbdhu8.tar.gz ${TIGER}

RUN mkdir -p ${TIGER_OUT} \
    && for f in *.tar.gz; do tar -xzf "$f"; done \
    && shp2pgsql -s 4269:4326 -d tl_2024_us_county tiger_data.county > ${TIGER_OUT}/tiger_county.sql \
    && shp2pgsql -s 4269:4326 -d tl_2024_us_state  tiger_data.state  > ${TIGER_OUT}/tiger_state.sql \
    && shp2pgsql -s 4326 -d wbdhu6 hhc.wbdhu6 > ${TIGER_OUT}/wbdhu6.sql

# -------- Stage 2 --------
# -------- Runner --------
FROM ghcr.io/cwbi-apps/flyway:latest AS runner

# env vars for FLYWAY
ENV FLYWAY_EDITION=community \
    FLYWAY_LOCATIONS=filesystem:/flyway/sql

RUN apk upgrade --no-cache \
    && apk add postgresql-client

COPY ./main /flyway/sql/main

WORKDIR /tmp
COPY ./dist_geom ./dist_geom
COPY --from=builder /tmp/ ./

# entrypoint for flyway with arguments (CMD)
WORKDIR /flyway

ENTRYPOINT ["flyway"]

CMD ["-connectRetries=60", "migrate"]
