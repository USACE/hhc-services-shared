ARG GO_VERSION=1.25
ARG APPNAME=hhcshare

# ----------------------------------------------------------------------
# FROM build
FROM golang:${GO_VERSION}-alpine AS gobuild
# name it above or overwrite in IaC.
ARG APPNAME

RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates

WORKDIR /opt/go/src

COPY ./${APPNAME} ./

# Install Dependencies
RUN go mod download && go mod verify

# adding our user for the scatch image to run as
RUN adduser -D -u 10001 appuser && echo $(grep appuser /etc/passwd) > /etc/passwd

# ----------------------------------------------------------------------
# IMAGE FOR DEVELOPMENT
FROM gobuild AS development
RUN go install github.com/go-delve/delve/cmd/dlv@latest \
    && go get github.com/go-delve/delve

# ----------------------------------------------------------------------
# IMAGE FOR PRODUCTION THAT SCRATCH USES
FROM gobuild AS production
RUN go build -ldflags="-w -s" -o /opt/go/bin/app

# ----------------------------------------------------------------------
# SCRATCH IMAGE
FROM scratch
COPY --from=production /opt/go/bin/app /app
COPY --from=gobuild /etc/passwd /etc/passwd
COPY --from=gobuild /etc/ssl/certs/ca-certificates.crt \
    /etc/ssl/certs/ca-certificates.crt

USER 10001

ENTRYPOINT [ "/app" ]
