FROM golang:1.16-alpine

RUN apk add --no-cache \
      git \
      tzdata

RUN cp /usr/share/zoneinfo/PRC /etc/localtime && \
      echo "PRC" > /etc/timezone && \
      apk del tzdata
