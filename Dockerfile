FROM golang:1.20 as builder

WORKDIR /go/src/github.com/kubemeta/cle-simple

COPY cmd/bootstrap .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/clectl github.com/kubemeta/cle-simple/cmd/clectl

FROM alpine:3.17.2

COPY --from=builder /go/bin/cle-simple /usr/local/bin/cle-simple

COPY config /etc/config

EXPOSE 9083

ENTRYPOINT ["/usr/local/bin/clectl"]

