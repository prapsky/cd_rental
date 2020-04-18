# CD Rental
Count payment of CD that we rent at CD Rental.

## Setup
1. [go mod](#go-mod)
2. [Unit Test](#unit-test)
3. [Run Program](#run-program)

## Setup PostgreSQL
1. [Setup at Local](#setup-at-local)
2. [Configure at Code](#configure-at-code)
3. [Create Tables](#create-tables)

### go mod
Execute go mod at root this folder using this command:
```
$ go mod init cd_rental
```
Open go.mod then we should see:
```
module cd_rental

go 1.13
```

### Unit Test
First, we need to remove cache by using this command:
```
$ go clean -testcache
```
Execute unit test at root this folder using this command:
```
$ go test ./...
```

### Run Program
Run the program at root this folder using this command:
```
$ go run main.go
```

### Setup at Local
For the first, install PostgreSQL to your local.
Access PostgreSQL using this command:
```
$ psql
```
Then we should see:
```
psql (12.2)
Type "help" for help.

prapsky=#
```
For example, the name of PostgreSQL's user is prapsky.
Create database, for example the name is cdrental, using this command:
```
$ CREATE DATABASE cdrental;
```
Check the database using this command:
```
$ \l
```
Then we should see:
```
                             List of databases
   Name    |  Owner   | Encoding | Collate | Ctype |   Access privileges
-----------+----------+----------+---------+-------+-----------------------
 cdrental  | prapsky  | UTF8     | C       | C     |
 prapsky   | prapsky  | UTF8     | C       | C     |
 template0 | prapsky  | UTF8     | C       | C     | =c/prapsky           +
           |          |          |         |       | prapsky=CTc/prapsky
 template1 | prapsky  | UTF8     | C       | C     | =c/prapsky           +
           |          |          |         |       | prapsky=CTc/prapsky
(4 rows)
```
Connect to database, for example the name is cdrental, using this command:
```
$ \c cdrental
```
Then we should see:
```
You are now connected to database "cdrental" as user "prapsky".
cdrental=#
```

### Configure at Code
For example, the name of PostgreSQL's user is prapsky, the name of database is cdrental. Edit at db_connection.go:
```
DB, err = sql.Open("postgres", "user=prapsky dbname=cdrental sslmode=disable")
```

### Create Tables
Create Collection table:
```
CREATE TABLE collection (id SERIAL PRIMARY KEY NOT NULL, date_time TIMESTAMP NOT NULL, title TEXT NOT NULL, category TEXT NOT NULL, quantity INT DEFAULT 0 NOT NULL, rate INT DEFAULT 0 NOT NULL);
```
