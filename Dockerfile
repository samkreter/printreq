FROM golang:1.9.2 as builder
WORKDIR /go/src/github.com/samkreter/printreq/
COPY . /go/src/github.com/samkreter/printreq/
# RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o run .

FROM alpine:latest
RUN apk --update add ca-certificates
WORKDIR /
COPY --from=builder /go/src/github.com/samkreter/printreq/run .

EXPOSE 8081

CMD ["./run"]