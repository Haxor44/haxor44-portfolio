FROM golang:1.26.0-alpine AS builder
WORKDIR /app
  
# caching dependencies sepearately from source changes
COPY go.mod go.sum ./
RUN go mod download

COPY . . 
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /server ./cmd/server

# ---- Runtime stage ----
FROM alpine:3.20
RUN apk add --no-cache ca-certificates && \
    addgroup -S app && adduser -S -G app app
WORKDIR /app

COPY --from=builder /server ./server
COPY --from=builder /app/web ./web

USER app
EXPOSE 3000
ENTRYPOINT ["./server"]