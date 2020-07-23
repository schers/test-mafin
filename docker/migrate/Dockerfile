FROM migrate/migrate

RUN set -xe \
 && apk update \
 && apk --no-cache add \
 bash

COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh

RUN mkdir /migrations

ENTRYPOINT ["/entrypoint.sh"]