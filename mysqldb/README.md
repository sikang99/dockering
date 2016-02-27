README.md
=========

# Docker Images
https://github.com/docker-library/official-images/
https://hub.docker.com/_/mariadb/

# Awesome MySQL
https://shlomi-noach.github.io/awesome-mysql/

# Golang, Docker, and databases
https://medium.com/@Martynas/golang-docker-and-databases-125a5c1894a8#.lbfchfqp6

# Installing and using MariaDB via Docker
https://mariadb.com/kb/en/mariadb/installing-and-using-mariadb-via-docker/

#  MySQL Documentation
http://dev.mysql.com/doc/

```
$ mysql --user=root --passwd=root
CREATE DATABASE testdb;
SELECT VERSION(), NOW(), CURRENT_DATE, USER();
USE testdb;

$ mysql mysql --user=root --password=root testdb
CREATE TABLE names (id serial, name varchar(32));
SHOW TABLES;
SHOW CREATE TABLE t1\G;
DESCRIBE names;
INSERT INTO names VALUES ('1','Diane');
SELECT id, name FROM names;
SELECT MAX(id) AS id FROM name;
```

