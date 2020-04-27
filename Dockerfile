FROM alpine

RUN mkdir /app

COPY track-progress /app

WORKDIR /app

CMD ["/app/track-progress"]