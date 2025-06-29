FROM golang:1.22-bullseye AS builder

WORKDIR /app
COPY . .

RUN go build -ldflags="-s -w" -o server main.go



FROM scratch
COPY --from=builder /app/server /server
COPY public/ /public
EXPOSE 80
ENTRYPOINT ["/server"]
