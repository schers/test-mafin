#depricated
FROM migrate/migrate

RUN set -xe \
 && apk update \
 && apk --no-cache add \
 bash

COPY wait-for-it.sh /usr/bin/wait-for-it
RUN chmod +x /usr/bin/wait-for-it

COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]