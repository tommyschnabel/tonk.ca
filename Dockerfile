FROM golang:1.25.1 AS builder

# Search
COPY search /search
RUN cd /search && make index build

FROM ghcr.io/linuxserver/baseimage-ubuntu:noble

# Static files
COPY static /static

# Search
COPY --from=builder /search/search_server /search_server
COPY --from=builder /search/index.bleve /index.bleve

# Install nginx to serve static files, and proxy to search server
RUN apt update && apt install -y nginx
COPY nginx.conf /etc/nginx/nginx.conf

# s6-rc services
COPY root/ /
