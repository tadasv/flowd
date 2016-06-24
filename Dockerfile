FROM busybox
MAINTAINER Tadas Vilkeliskis <vilkeliskis.t@gmail.com>
EXPOSE 7777
ENTRYPOINT ["/flowd"]
COPY ./src/flowd /flowd
