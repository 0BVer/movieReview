FROM arm64v8/mysql:8.0.30

COPY init.sql /docker-entrypoint-initdb.d/init.sql