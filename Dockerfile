# build phase
FROM golang:1.18

WORKDIR /app

COPY . /app

ENV CGO_ENABLED=0

RUN go build -o gotemplate ./cmd

# execution phase
FROM alpine:latest

WORKDIR /

COPY --from=0 /app/gotemplate ./
COPY --from=0 /app/.config ./.config

CMD ["./gotemplate"]
