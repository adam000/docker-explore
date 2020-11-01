FROM golang:1.15.3 as build
ENV GO111MODULE on

WORKDIR /go/src/app

# Warm up the module cache.
# Only copy in go.mod and go.sum to increase Docker cache hit rate.
COPY go.mod go.sum /src/
WORKDIR /src
RUN go mod download
# Now build the whole tree.
COPY . /src
RUN go build -o app

FROM gcr.io/distroless/base-debian10
COPY --from=build /src/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
