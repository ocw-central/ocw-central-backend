FROM migrate/migrate

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

COPY ./docker/migrate/migrate.sh /usr/local/bin
RUN chmod +x /usr/local/bin/migrate.sh


CMD [ "migrate.sh" ]

