FROM --platform=$BUILDPLATFORM node:21-alpine as nodebuilder
USER node

COPY --chown=node ./front /app
WORKDIR /app/front
RUN npm i && npm run build

FROM golang:1.21-alpine as builder
ENV PORT=8080

WORKDIR /app

COPY go.* /app
RUN go mod download

COPY atom /app/atom
COPY dockerhub /app/dockerhub
COPY *.go /app

RUN go build

FROM alpine

COPY --from=builder /app/dockerRSS /usr/local/bin/dockerRSS
COPY --from=nodebuilder /app/dist /usr/local/bin/dist

EXPOSE $PORT

ENV HTTP_STATIC_DIR=/usr/local/bin/dist

CMD ["dockerRSS"]