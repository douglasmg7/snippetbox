# Source image.
FROM golang:1.23.1-alpine3.20 AS snippetbox

RUN mkdir -p /code

COPY . /code

WORKDIR /code

RUN go build -o snippetbox ./cmd/web

# Tiny image.
FROM alpine:latest

RUN mkdir -p /app

WORKDIR /app

# Exectuable only.
COPY --from=snippetbox /code/snippetbox /app
# Templates and statics.
COPY --from=snippetbox /code/ui /app/ui

CMD [ "/app/snippetbox" ]