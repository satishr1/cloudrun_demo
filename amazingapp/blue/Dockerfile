FROM golang:1.17 as builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server

FROM alpine:3.14
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/server /
CMD ["/server"]