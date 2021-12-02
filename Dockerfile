FROM golang:1.16-alpine AS build
ARG APP_NAME
RUN apk add gcc musl-dev
WORKDIR /src
RUN go get github.com/cespare/reflex
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/$APP_NAME -ldflags "-s -w" ./cmd/serve

FROM alpine:3.14
ARG APP_NAME
RUN apk --no-cache -U upgrade
WORKDIR /app
COPY --from=build /app/$APP_NAME .
USER guest
CMD ["/app/$APP_NAME"]
