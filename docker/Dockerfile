FROM golang:1.22.0-alpine3.19 as builder
COPY ./LRUcache /go/LRUcache
WORKDIR /go/LRUcache
RUN go mod vendor
RUN cd ./cmd && \
go build -o main .

FROM alpine:latest
WORKDIR /go/
EXPOSE 8100
COPY --from=builder /go/LRUcache/cmd/main ./app
CMD ["./app", "-listening-address=0.0.0.0:8100"]

