FROM golang:1.27.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o ./server ./CMD/main.go

FROM alpine:3.21

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/server ./

RUN mkdir -p public/img public/docs && chmod +x /app/server

EXPOSE 3000

CMD ["/app/server"]



# FROM golang:1.27.1-trixie AS builder

# WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/server ./cmd/main.go

# FROM alpine:3.24

# WORKDIR /app

# RUN apk add --no-cache ca-certificates

# COPY --from=builder /app/server ./server
# COPY --from=builder /app/public ./public

# LABEL org.opencontainers.image.source="https://github.com/habibmrizki/Koda-b9-EventHub-BE"

# EXPOSE 3000

# CMD ["/app/server"]