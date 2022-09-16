FROM golang:1.17.7-alpine3.15 AS builder
WORKDIR /go/src
COPY ./ . 
RUN go build . 

FROM alpine:3.15

COPY --from=builder /go/src/gotemplater /bin/
