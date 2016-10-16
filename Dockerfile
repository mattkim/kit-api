FROM alpine:latest

MAINTAINER Edward Muller <edward@heroku.com>

WORKDIR "/opt"

ADD .docker_build/kit-api /opt/bin/kit-api
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/kit-api"]

