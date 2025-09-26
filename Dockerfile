# Build Geth in a stock Go builder container
FROM golang:1.24.2-alpine as builder

RUN apk add --no-cache make gcc  g++ musl-dev linux-headers git bash  libstdc++

ADD . /go-ethereum
RUN cd /go-ethereum && make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates curl jq tini libstdc++
COPY --from=builder /go-ethereum/build/bin/geth /usr/local/bin/

EXPOSE 8545 8546 8547 30303 30303/udp
ENTRYPOINT ["geth"]
