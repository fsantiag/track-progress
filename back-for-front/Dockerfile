FROM debian:stable-slim

RUN mkdir /app

COPY src/bff /app
COPY wait.sh /app

WORKDIR /app

CMD ["/app/bff"]