FROM postgres:16
COPY sc-database.init.sql /docker-entrypoint-initdb.d/