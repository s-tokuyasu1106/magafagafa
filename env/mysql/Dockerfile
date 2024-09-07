FROM mysql:latest

ENV TZ Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && chown -R mysql:root /var/lib/mysql/

COPY my.cnf /etc/mysql/conf.d/my.cnf

CMD ["mysqld"]

EXPOSE 3306