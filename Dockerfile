FROM golang:1.20-bullseye AS builder

EXPOSE 80

WORKDIR /usr/src

COPY . .

RUN go build -o fiber-postgres-template .

FROM debian:bullseye-slim

COPY --from=builder /usr/src/fiber-postgres-template /usr/local/bin/fiber-postgres-template

CMD ["fiber-postgres-template"]