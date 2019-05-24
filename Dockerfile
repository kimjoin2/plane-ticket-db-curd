FROM golang:1.11-alpine AS build
COPY ./ /go/src/plane-ticket-db-curd/
WORKDIR /go/src/plane-ticket-db-curd/
RUN go build -o plane-ticket-db-curd

FROM alpine:3.7
RUN apk --no-cache --update upgrade \
  && apk add --update --no-cache tzdata ca-certificates \
  && update-ca-certificates --fresh
RUN mkdir -p /app
COPY --from=build /go/src/plane-ticket-db-curd/plane-ticket-db-curd /app/plane-ticket-db-curd

EXPOSE 9001

ENTRYPOINT ["/app/plane-ticket-db-curd"]