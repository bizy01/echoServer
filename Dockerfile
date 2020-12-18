FROM alpine:latest

MAINTAINER liushaobo <liushaobo101@gmail.com>

RUN mkdir -p  /dev-work
RUN mkdir -p  /dev-work/conf
RUN mkdir -p /log

COPY ./conf  /dev-work/conf
COPY ./server /dev-work
WORKDIR /dev-work
RUN chmod +x server

CMD ./server