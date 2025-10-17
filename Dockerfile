FROM ghcr.io/linuxserver/baseimage-ubuntu:noble

# Static files
COPY static /static

# 
COPY root/ /root-layer/