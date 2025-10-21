FROM ghcr.io/linuxserver/baseimage-ubuntu:noble

RUN apt update && apt install golang

# Static files
COPY static /static

# 
COPY root/ /root-layer/
