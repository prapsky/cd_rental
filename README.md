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

## API
1. [Collection](#collection)
2. [User](#user)

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
For the first, install PostgreSQL to your local. <br>
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
For example, the name of PostgreSQL's user is prapsky. <br>
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
Show tables at database using this command:
```
$ \dt
```
For example we should see:
```
cdrental-# \dt
           List of relations
 Schema |    Name    | Type  |  Owner
--------+------------+-------+----------
 public | collection | table | prapsky
(1 row)

cdrental-#
```
To quit from PostgreSQL, we can using this command:
```
$ \q
```

### Configure at Code
For example, the name of PostgreSQL's user is prapsky, the name of database is cdrental. <br> Edit at db_connection.go:
```
DB, err = sql.Open("postgres", "user=prapsky dbname=cdrental sslmode=disable")
```

### Create Tables
#### Collection Table
Create Collection table:
```
CREATE TABLE collection (id SERIAL PRIMARY KEY NOT NULL, date_time TIMESTAMP NOT NULL, title TEXT NOT NULL, category TEXT NOT NULL, quantity INT DEFAULT 0 NOT NULL, rate INT DEFAULT 0 NOT NULL);
```
#### Users Table
Create Users table:
```
CREATE TABLE users (id SERIAL PRIMARY KEY NOT NULL, date_time TIMESTAMP NOT NULL, name TEXT NOT NULL, phone_number TEXT NOT NULL, address TEXT NOT NULL);
```

### Collection
#### POST - /collection
Request
```
{
    "title": "Star Wars",
    "category": "Sci-Fi",
    "quantity": 20,
    "rate": 15000
}
```
Response Body (Status: 201 Created)
```
{
    "id": 1,
    "dateTime": "2020-04-18T23:52:40.238858+07:00",
    "title": "Star Wars",
    "category": "Sci-Fi",
    "quantity": 20,
    "rate": 15000
}
```
#### GET - /collection/{collection_id}
Example: /collection/1 <br> Response Body (Status: 200 OK)
```
{
    "id": 1,
    "dateTime": "2020-04-18T23:52:40.238858Z",
    "title": "Star Wars",
    "category": "Sci-Fi",
    "quantity": 20,
    "rate": 15000
}
```
#### GET - /collection/all
Example: /collection/all <br> Response Body (Status: 200 OK)
```
{
    "collections": [
        {
            "id": 1,
            "dateTime": "2020-04-18T23:52:40.238858Z",
            "title": "Star Wars",
            "category": "Sci-Fi",
            "quantity": 20,
            "rate": 15000
        },
        {
            "id": 2,
            "dateTime": "2020-04-19T00:21:04.292774Z",
            "title": "Captain America",
            "category": "Sci-Fi",
            "quantity": 20,
            "rate": 10000
        },
        {
            "id": 3,
            "dateTime": "2020-04-19T15:21:58.669116Z",
            "title": "James Bond",
            "category": "Action",
            "quantity": 10,
            "rate": 10000
        },
        {
            "id": 4,
            "dateTime": "2020-04-19T15:23:53.729481Z",
            "title": "La La Land",
            "category": "Drama",
            "quantity": 10,
            "rate": 7500
        },
        {
            "id": 5,
            "dateTime": "2020-04-19T15:24:19.87131Z",
            "title": "The Social Network",
            "category": "Drama",
            "quantity": 5,
            "rate": 5000
        }
    ]
}
```
#### PUT - /collection/{collection_id}
Example: /collection/2 <br>
Request
```
{
    "id": 2,
    "title": "Captain America",
    "category": "Sci-Fi",
    "quantity": 19,
    "rate": 10000
}
```
Response Body (Status: 200 OK)
```
{
    "id": 2,
    "dateTime": "2020-04-19T19:25:41.436054+07:00",
    "title": "Captain America",
    "category": "Sci-Fi",
    "quantity": 19,
    "rate": 10000
}
```

### User
#### POST - /user
Request
```
{
	"name": "Ihsan",
	"phoneNumber": "085624136133",
	"address": "Jalan K no.11 Jakarta Selatan"
}
```
Response Body (Status: 201 Created)
```
{
    "id": 11,
    "dateTime": "2020-04-19T18:33:00.141695+07:00",
    "name": "Ihsan",
    "phoneNumber": "085624136133",
    "address": "Jalan K no.11 Jakarta Selatan"
}
```
#### GET - /user/{user_id}
Example: /user/1 <br> 
Response Body (Status: 200 OK)
```
{
    "id": 1,
    "dateTime": "2020-04-19T17:09:26.710061Z",
    "name": "Jeffrey",
    "phoneNumber": "085624136123",
    "address": "Jalan A no.1 Jakarta Selatan"
}
```
#### GET - /user/all
Example: /user/all <br> 
Response Body (Status: 200 OK)
```
{
    "users": [
        {
            "id": 1,
            "dateTime": "2020-04-19T17:09:26.710061Z",
            "name": "Jeffrey",
            "phoneNumber": "085624136123",
            "address": "Jalan A no.1 Jakarta Selatan"
        },
        {
            "id": 2,
            "dateTime": "2020-04-19T18:25:52.696274Z",
            "name": "Jose",
            "phoneNumber": "085624136124",
            "address": "Jalan B no.2 Jakarta Selatan"
        },
        {
            "id": 3,
            "dateTime": "2020-04-19T18:27:45.002081Z",
            "name": "Alvi",
            "phoneNumber": "085624136125",
            "address": "Jalan C no.3 Jakarta Selatan"
        },
        {
            "id": 4,
            "dateTime": "2020-04-19T18:28:49.214656Z",
            "name": "Sidney",
            "phoneNumber": "085624136126",
            "address": "Jalan D no.4 Jakarta Selatan"
        },
        {
            "id": 5,
            "dateTime": "2020-04-19T18:29:27.414545Z",
            "name": "Kenny",
            "phoneNumber": "085624136127",
            "address": "Jalan E no.5 Jakarta Selatan"
        },
        {
            "id": 6,
            "dateTime": "2020-04-19T18:30:08.336506Z",
            "name": "Joe",
            "phoneNumber": "085624136128",
            "address": "Jalan F no.6 Jakarta Selatan"
        },
        {
            "id": 7,
            "dateTime": "2020-04-19T18:31:02.737774Z",
            "name": "Jussar",
            "phoneNumber": "085624136129",
            "address": "Jalan G no.7 Jakarta Selatan"
        },
        {
            "id": 8,
            "dateTime": "2020-04-19T18:31:34.965212Z",
            "name": "Gea",
            "phoneNumber": "085624136130",
            "address": "Jalan H no.8 Jakarta Selatan"
        },
        {
            "id": 9,
            "dateTime": "2020-04-19T18:32:01.429844Z",
            "name": "Ary",
            "phoneNumber": "085624136131",
            "address": "Jalan I no.9 Jakarta Selatan"
        },
        {
            "id": 10,
            "dateTime": "2020-04-19T18:32:28.686109Z",
            "name": "Remy",
            "phoneNumber": "085624136132",
            "address": "Jalan J no.10 Jakarta Selatan"
        },
        {
            "id": 11,
            "dateTime": "2020-04-19T18:33:00.141695Z",
            "name": "Ihsan",
            "phoneNumber": "085624136133",
            "address": "Jalan K no.11 Jakarta Selatan"
        }
    ]
}
```
