# syntax=docker/dockerfile:1

## Build step 1 - Frontend
FROM node:lts AS frontend

WORKDIR /app/frontend
COPY frontend ./

RUN rm -f ./dist/gitkeep && \
    ls -l && \
    npm install && \
    npm run prod-js && \
    npm run prod-css

## Build step 2 - Backend
FROM golang:1.22.1-bookworm AS backend

WORKDIR /app/backend

COPY backend ./
COPY --from=frontend /app/frontend/dist ./pkg/server/routes/web/static/dist
COPY --from=frontend /app/frontend/src ./pkg/server/routes/web/static

RUN rm -f ./pkg/server/routes/web/static/gitkeep && \
    go mod tidy && \
    go build -o /app/app ./cmd/main.go

## Build step 3 - Delpoyment
FROM debian:stable-slim

WORKDIR /app
COPY --from=backend /app/app ./

ARG CA_CERTIFICATES_VERSION=20230311        # https://packages.debian.org/bookworm/ca-certificates
ARG CURL_VERSION=7.88.1-10+deb12u5          # https://packages.debian.org/bookworm/curl

RUN apt-get update && \
    apt-get install -y --no-install-recommends ca-certificates=${CA_CERTIFICATES_VERSION} curl=${CURL_VERSION}

HEALTHCHECK CMD curl --fail http://localhost:80

#checkov:skip=CKV_DOCKER_3:Irrelevant

EXPOSE 80/tcp

ENTRYPOINT ["./app"]
