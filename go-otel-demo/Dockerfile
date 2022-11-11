##
## Build
##
FROM golang:1.17.6-alpine3.15 AS build

#
WORKDIR /app

COPY . .
RUN go build -o /web-app

##
## Deploy
##
FROM alpine:3.15

WORKDIR /app

COPY --from=build /web-app ./
COPY ./.env ./

CMD ["./web-app"]