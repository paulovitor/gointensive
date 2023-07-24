FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o api cmd/api/main.go

# menor imagem docker, vem zerada
FROM scratch

COPY --from=builder /app/api /

CMD ["./api"]