FROM golang:1.16-alpine AS build
RUN apk add gcc musl-dev
WORKDIR /src
RUN go get github.com/cespare/reflex
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/im-users -ldflags "-s -w" ./cmd/serve

FROM alpine:3.14
RUN apk --no-cache -U upgrade
WORKDIR /app
COPY --from=build /app/im-users .
COPY --from=build /src/swagger/swagger.yaml ./swagger/
USER guest
CMD ["/app/im-users"]
