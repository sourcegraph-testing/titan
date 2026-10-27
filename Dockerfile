# Builder image
FROM golang:latest@sha256:f38c7f7bbaca5d664e39cd982a1cb5b6f8e999244e9ddb6ec8ba098438b3f4da as builder

RUN apk add --no-cache \
    make \
    git

COPY . /go/src/github.com/distributedio/titan

WORKDIR /go/src/github.com/distributedio/titan

RUN env GOOS=linux CGO_ENABLED=0 make

# Executable image
FROM alpine

COPY --from=builder /go/src/github.com/distributedio/titan/titan /titan/bin/titan
COPY --from=builder /go/src/github.com/distributedio/titan/conf/titan.toml /titan/conf/titan.toml

WORKDIR /titan

EXPOSE 7369

ENTRYPOINT ["./bin/titan"]
