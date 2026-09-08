ARG GO_VERSION=1.27.0

# ----------------------------------------------------------------------
# BUILDER BUILD STAGE
FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates

WORKDIR /opt/go/src

COPY . .

# Install Dependencies
RUN go mod download && go mod verify

# adding our user for the scatch image to run as
RUN adduser -D -u 10001 appuser && echo $(grep appuser /etc/passwd) > /etc/passwd

# ----------------------------------------------------------------------
# DEVELOPMENT BUILD STAGE
FROM builder AS development
RUN go install github.com/go-delve/delve/cmd/dlv@latest \
    && go get github.com/go-delve/delve

# ----------------------------------------------------------------------
# PRODUCTION BUILD STAGE
FROM builder AS production
RUN go build -ldflags="-w -s" -o /opt/go/bin/app

# ----------------------------------------------------------------------
# SCRATCH BUILD STAGE
FROM scratch
COPY --from=production /opt/go/bin/app /app
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt \
    /etc/ssl/certs/ca-certificates.crt

USER 10001

ENTRYPOINT [ "/app" ]
