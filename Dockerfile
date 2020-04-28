FROM alpine

RUN mkdir /app

COPY /api/track-progress /app

WORKDIR /app

CMD ["/app/track-progress"]