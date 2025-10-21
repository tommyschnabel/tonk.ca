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
 
# s6-rc services
COPY root/ /
