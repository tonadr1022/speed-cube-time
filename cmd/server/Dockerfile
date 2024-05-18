FROM golang:1.22.3 AS build-stage

# change work dir
WORKDIR /app 

# copy mod files into app dir and download
COPY go.* ./  
RUN go mod download
RUN go mod verify


# copy source files

# build 
COPY . .
RUN make build

# run tests in the container 
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# deploy binary into a lean image 
FROM gcr.io/distroless/base-debian12 AS build-release-stage 

WORKDIR / 

# copy the binary from build stage to the second image
COPY --from=build-stage /speed-cube-time /speed-cube-time

EXPOSE 8080 

USER nonroot:nonroot

CMD ["/speed-cube-time"]
