FROM golang:1.25.1 AS builder

# Copy files
COPY search /search
COPY static /static
COPY ocr/results.zip /

# Install deps
RUN apt update && apt install -y unzip

# Unzip ocr results
RUN unzip -:o /results.zip

# Build index and search server
RUN cd /search && make index build

FROM ghcr.io/linuxserver/baseimage-ubuntu:noble

# Search
COPY --from=builder /search/search_server /search_server
COPY --from=builder /search/index.bleve /index.bleve

# Install nginx to serve static files, and proxy to search server
RUN apt update && apt install -y nginx
COPY nginx.conf /etc/nginx/nginx.conf

# Static files
COPY static /static

# s6-rc services
COPY root/ /
