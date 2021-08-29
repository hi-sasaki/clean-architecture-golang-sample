FROM golang:alpine as builder
WORKDIR /go/src/github.com/alexellis/href-counter/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app  ./cmd/api/main.go  ./cmd/api/wire_gen.go
RUN apk update && apk add --no-cache ca-certificates tzdata && update-ca-certificates


FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/alexellis/href-counter/app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo

CMD ["./app"]