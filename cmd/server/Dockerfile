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

# copy mod files into app dir and download
COPY go.* ./  
RUN go mod download
RUN go mod verify

# install migrate for db migrations
ARG MIGRATE_VERSION=4.17.1
ADD https://github.com/golang-migrate/migrate/releases/download/v${MIGRATE_VERSION}/migrate.linux-amd64.tar.gz /tmp
RUN tar -xzf /tmp/migrate.linux-amd64.tar.gz -C /usr/local/bin 
# copy source files

# build 
COPY . .
ENV CGO_ENABLED=0
RUN go build -o /tmp/server ./cmd/server/*.go

# run tests in the container 
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# deploy binary into lean image
FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
WORKDIR /app

# copy the binary from build stage to the second image
COPY --from=built /usr/local/bin/migrate /usr/local/bin
COPY --from=built /app/migrations ./migrations/
COPY --from=built /tmp/server /usr/bin/server
COPY --from=built /app/cmd/server/entrypoint.sh .

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]

