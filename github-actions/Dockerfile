FROM youyo/awscredswrap:latest

LABEL maintainer "youyo <1003ni2@gmail.com>"

RUN apk add --update --no-cache bash

COPY entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
