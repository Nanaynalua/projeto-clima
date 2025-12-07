FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copiar arquivos de dependência
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fonte
COPY . .

# Compilar binário
RUN go build -o worker main.go

# Imagem final
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/worker .

CMD ["./worker"]
