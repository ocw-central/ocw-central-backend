FROM mysql:8.0.30

#MySQL設定ファイルをイメージ内にコピー
COPY ./docker/mysql/my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]