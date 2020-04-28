FROM alpine

RUN mkdir /app

COPY /backend/track-progress /app

WORKDIR /app

CMD ["/app/track-progress"]