ARG PHP_VERSION=8.0

FROM php:${PHP_VERSION}-alpine AS api_platform_php

RUN apk add --no-cache bash

WORKDIR /srv/api

EXPOSE 8000
ENTRYPOINT ["php"]
CMD ["-S", "0.0.0.0:8000", "-t", "public"]
