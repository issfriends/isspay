FROM golang:latest AS builder
WORKDIR /isspay
ENV GO111MODULE=on 

COPY . .

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build  -mod=vendor -o main

FROM alpine:latest
ARG BUILD_TIME
ARG SHA1_VER


RUN apk update && \
    apk upgrade && \
    apk add --no-cache curl tzdata && \
    apk add ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /isspay
COPY --from=builder /isspay/main /isspay/main

RUN ls
ENV SHA1_VER=${SHA1_VER}
ENV BUILD_TIME=${BUILD_TIME}
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser && \
    chown -R appuser:appuser /isspay
USER isspay

CMD ["./main", "isspay"]
