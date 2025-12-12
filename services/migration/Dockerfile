# *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
# FROM runner
# *~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~*~
FROM ghcr.io/cwbi-apps/flyway:latest AS runner

# env vars for FLYWAY
ENV FLYWAY_EDITION=community
ENV FLYWAY_LOCATIONS=filesystem:/flyway/sql

# version of the archive for the geometry migration
ARG GEOM_VERSION=v0.1.1

# env vars
ENV OPT_MIGRATION=/opt/migration

# Building with cwbi-apps flyway
RUN apk update \
    && apk add postgis postgresql-client unzip curl

# the next two RUNs are for a base image flyway/flyway if needed
# solving possible issue of APT behind a proxy
# RUN cat <<EOF >/etc/apt/apt.conf.d/99proxy
# Acquire::http::Pipeline-Depth 0;
# Acquire::http::No-Cache true;
# Acquire::BrokenProxy true;
# EOF

# RUN apt-get update && apt-get install -y postgis postgresql-client unzip curl

# the main migration into the image
# local migrations for testing is defined in the docker-compose file
COPY ./main /flyway/sql/main

# set the working directory
WORKDIR ${OPT_MIGRATION}

# District Geometry used in a bash script migration
COPY ./dist_geom ./dist_geom

# add the release asset for the tagged version defined in GEOM_VERSION
# tiger data for counties
ADD --checksum=sha256:89535e9a798e4feca0a0e786041274b4591b61186a0f9a62b8f478e1d70771e2 https://github.com/USACE/hhc-services-shared/releases/download/${GEOM_VERSION}/tl_2024_us_county.tar.gz ${OPT_MIGRATION}
# tiger data for states
ADD --checksum=sha256:c5e25ab502688993cc509721011da08b104602bc0e2eec8bb5e2da9c72955ec2 https://github.com/USACE/hhc-services-shared/releases/download/${GEOM_VERSION}/tl_2024_us_state.tar.gz ${OPT_MIGRATION}
# HUC6
ADD --checksum=sha256:dc9c7e0044eb9686988e0cff0608fcce5e5658a556b21c8b982ba91a01178910 https://github.com/USACE/hhc-services-shared/releases/download/${GEOM_VERSION}/wbdhu6.tar.gz ${OPT_MIGRATION}
# HUC8
# ADD --checksum=sha256:84d9dc1caf680506cb79938ec4ebc77558050a4027917f9a4e2aa2f64ac7b537 https://github.com/USACE/hhc-services-shared/releases/download/${GEOM_VERSION}/wbdhu8.tar.gz ${OPT_MIGRATION}

# entrypoint for flyway with arguments (CMD)
WORKDIR /flyway

ENTRYPOINT ["flyway"]

CMD ["-connectRetries=60", "migrate"]
