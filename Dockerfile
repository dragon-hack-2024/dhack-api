# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY wait-for.sh .
RUN chmod +x ./wait-for.sh

EXPOSE 8080
CMD [ "/app/main" ]