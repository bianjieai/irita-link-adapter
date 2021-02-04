FROM golang:1.15-alpine3.12 as builder

WORKDIR /irita-link-adapter
COPY . .

RUN apk add make && make install

FROM alpine:3.12

COPY --from=builder /go/bin/irita-link-adapter /usr/local/bin/

EXPOSE 8080
ENTRYPOINT ["irita-link-adapter"]
