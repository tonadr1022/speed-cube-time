FROM golang:alpine AS built
RUN apk update && \
    apk add curl \
    git \
    bash \
    make \
    ca-certificates && \
    rm -rf /var/cache/apk/*

# change work dir
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
ENV CGO_ENABLED=0
RUN ls
RUN go build -o /tmp/server ./cmd/server/main.go


# install migrate for db migrations
ARG MIGRATE_VERSION=4.17.1
ADD https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz /tmp
RUN tar -xzf /tmp/migrate.linux-amd64.tar.gz -C /usr/local/bin 

# deploy binary into lean image
FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
WORKDIR /app

# copy the binary from build stage to the second image
COPY --from=built /usr/local/bin/migrate /usr/local/bin
COPY --from=built /app/migrations ./migrations/
COPY --from=built /tmp/server /usr/bin/server
COPY --from=built /app/cmd/server/entrypoint.sh .

RUN chmod +x /app/entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]

