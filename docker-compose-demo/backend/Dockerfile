# syntax=docker/dockerfile:1.4
FROM --platform=$BUILDPLATFORM golang:1.22 AS builder

WORKDIR /code

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    GOPROXY=https://goproxy.cn,direct \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o bin/backend main.go

FROM scratch
COPY --from=builder /code/bin/backend /usr/local/bin/backend
CMD ["/usr/local/bin/backend"]