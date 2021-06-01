#
# Build and bundle the Vue.js frontend SPA
#
FROM node:14-alpine AS vue-build
WORKDIR /build

COPY http/frontend/package*.json ./
RUN npm install

COPY http/frontend .


RUN npm run build

#
# Build the Go server backend
#
FROM golang:1.16-alpine as go-build

WORKDIR /build/src/

RUN apk update && apk add git gcc musl-dev

COPY http/ ./http
COPY db/ ./db
COPY grpc/ ./grpc

ENV GO111MODULE=on
WORKDIR /build/src/http/server
# Disabling cgo results in a fully static binary that can run without C libs
RUN CGO_ENABLED=0 GOOS=linux go build -o main.go

#
# Assemble the server binary and Vue bundle into a single app
#
FROM alpine
WORKDIR /app

COPY --from=vue-build /build/dist ./dist
COPY --from=go-build /build/src/http/server .
COPY --from=go-build /build/src/db/schema ./db/schema



ENV PORT 8080
EXPOSE 8080
CMD ["/app/main.go"]
