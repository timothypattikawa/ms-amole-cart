FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY appliaction-*.yml .

EXPOSE 9093 9093
CMD [ "cmd/main" ]