FROM golang:alpine AS build-stage
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
RUN make build

# run tests in the container 
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# deploy binary into a lean image 
# FROM gcr.io/distroless/base-debian12 AS build-release-stage 
FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
WORKDIR /app

# copy the binary from build stage to the second image
COPY --from=build-stage /usr/local/bin/migrate /usr/local/bin
COPY --from=build-stage /app/migrations ./migrations/
COPY --from=build-stage /app/server .
COPY --from=build-stage /app/cmd/server/entrypoint.sh .
COPY --from=build-stage /app/config/*.yaml ./config/


EXPOSE 8080 

# USER nonroot:nonroot

ENTRYPOINT ["/app/entrypoint.sh"]
