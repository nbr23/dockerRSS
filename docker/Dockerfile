FROM golang:1.19-alpine as builder
ENV PORT=8080

WORKDIR /app

COPY atom /app/atom
COPY dockerhub /app/dockerhub
COPY *.go go.* /app

RUN go build

FROM alpine

COPY --from=builder /app/dockerRSS /usr/local/bin/dockerRSS

EXPOSE $PORT

CMD ["dockerRSS"]