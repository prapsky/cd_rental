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
3. [Rent](#rent)
4. [Rent All](#rent-all)
5. [Return](#return)
6. [Payment](#payment)

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
#### Rent Table
Create Rent table:
```
CREATE TABLE rent (id SERIAL PRIMARY KEY NOT NULL, date_time TIMESTAMP NOT NULL, queue_number INT NOT NULL, user_id INT REFERENCES users(id), cd_id INT REFERENCES collection(id), rent_quantity INT DEFAULT 0 NOT NULL);
```
#### Rent All Table
Create Rent All table:
```
CREATE TABLE rentall (id SERIAL PRIMARY KEY NOT NULL, date_time TIMESTAMP NOT NULL, queue_number INT NOT NULL);
```
#### Return Table
Create Return table:
```
CREATE TABLE return (id SERIAL PRIMARY KEY NOT NULL, date_time TIMESTAMP NOT NULL, rent_all_id INT REFERENCES rentall(id));
```
#### Payment Table
Create Payment table:
```
CREATE TABLE payment (id SERIAL PRIMARY KEY NOT NULL, date_time TIMESTAMP NOT NULL, return_id INT REFERENCES return(id), total_payment INT DEFAULT 0 NOT NULL);
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

### Rent
#### POST - /rent
Request
```
{
    "queueNumber": 1,
    "userId": 1,
    "cdId": 1,
    "rentQuantity": 1
}
```
Response Body (Status: 201 Created)
```
{
    "id": 1,
    "dateTime": "2020-04-19T22:38:40.12395+07:00",
    "queueNumber": 1,
    "userId": 1,
    "cdId": 1,
    "rentQuantity": 1
}
```
Ensure that quantity of that cd_id (cd_id = 1) is substracted by rent_quantity (rent_quantity = 1). <br>
For example, initial quantity of that cd_id (cd_id = 1) in collection is 20. After substracted by rent_quantity (rent_quantity = 1), the final quantity is 19. <br>
Check by GET request to /collection/1 <br> Response Body (Status: 200 OK)
```
{
    "id": 1,
    "dateTime": "2020-04-19T22:38:40.149754Z",
    "title": "Star Wars",
    "category": "Sci-Fi",
    "quantity": 19,
    "rate": 15000
}
```
#### GET - /rent/{rent_id}
Example: /rent/1 <br> 
Response Body (Status: 200 OK)
```
{
    "id": 1,
    "dateTime": "2020-04-19T22:38:40.12395Z",
    "queueNumber": 1,
    "userId": 1,
    "cdId": 1,
    "rentQuantity": 1
}
```
#### GET - /rent/queue/{queue_number}
Example: /rent/queue/1 <br> 
Response Body (Status: 200 OK)
```
{
    "rents": [
        {
            "id": 1,
            "dateTime": "2020-04-19T22:38:40.12395Z",
            "queueNumber": 1,
            "userId": 1,
            "cdId": 1,
            "rentQuantity": 1
        },
        {
            "id": 2,
            "dateTime": "2020-04-19T22:59:31.291183Z",
            "queueNumber": 1,
            "userId": 1,
            "cdId": 2,
            "rentQuantity": 2
        }
    ]
}
```

### Rent All
#### POST - /rent/all
Request
```
{
    "rents": [
        {
            "id": 1,
            "dateTime": "2020-04-19T22:38:40.12395Z",
            "queueNumber": 1,
            "userId": 1,
            "cdId": 1,
            "rentQuantity": 1
        },
        {
            "id": 2,
            "dateTime": "2020-04-19T22:59:31.291183Z",
            "queueNumber": 1,
            "userId": 1,
            "cdId": 2,
            "rentQuantity": 2
        }
    ]
}
```
Response Body (Status: 201 Created)
```
{
    "id": 1,
    "dateTime": "2020-04-20T01:20:00.274273+07:00",
    "queueNumber": 1,
    "userId": 1,
    "rents": [
        {
            "id": 1,
            "dateTime": "2020-04-19T22:38:40.12395Z",
            "queueNumber": 1,
            "userId": 1,
            "cdId": 1,
            "rentQuantity": 1
        },
        {
            "id": 2,
            "dateTime": "2020-04-19T22:59:31.291183Z",
            "queueNumber": 1,
            "userId": 1,
            "cdId": 2,
            "rentQuantity": 2
        }
    ]
}
```
#### GET - /rent/all/{rent_all_id}
Example: /rent/all/1 <br> 
Response Body (Status: 200 OK)
```
{
    "id": 1,
    "dateTime": "2020-04-20T01:20:00.274273Z",
    "queueNumber": 1,
    "userId": 1,
    "rents": [
        {
            "id": 1,
            "dateTime": "2020-04-19T22:38:40.12395Z",
            "queueNumber": 1,
            "userId": 1,
            "cdId": 1,
            "rentQuantity": 1
        },
        {
            "id": 2,
            "dateTime": "2020-04-19T22:59:31.291183Z",
            "queueNumber": 1,
            "userId": 1,
            "cdId": 2,
            "rentQuantity": 2
        }
    ]
}
```

### Return
#### POST - /return/all
Request
```
{
    "rentAllId": 1,
    "queueNumber": 1,
    "userId": 1,
    "rents": [
        {
            "id": 1,
            "dateTime": "2020-04-19T22:38:40.12395Z",
            "queueNumber": 1,
            "userId": 1,
            "cdId": 1,
            "rentQuantity": 1
        },
        {
            "id": 2,
            "dateTime": "2020-04-19T22:59:31.291183Z",
            "queueNumber": 1,
            "userId": 1,
            "cdId": 2,
            "rentQuantity": 2
        }
    ]
}
```
Response Body (Status: 201 Created)
```
{
    "id": 1,
    "dateTime": "2020-04-21T18:05:27.258262+07:00",
    "rentAllId": 1,
    "userId": 1,
    "returns": [
        {
            "cdId": 1,
            "returnQuantity": 1,
            "rentDays": 2,
            "ratePerDay": 15000,
            "totalRate": 30000
        },
        {
            "cdId": 2,
            "returnQuantity": 2,
            "rentDays": 2,
            "ratePerDay": 10000,
            "totalRate": 40000
        }
    ]
}
```
#### GET - /return/all/{return_id}
Example: /return/all/1 <br> 
Response Body (Status: 200 OK)
```
{
    "id": 1,
    "dateTime": "2020-04-21T18:05:27.258262Z",
    "rentAllId": 1,
    "userId": 1,
    "returns": [
        {
            "cdId": 1,
            "returnQuantity": 1,
            "rentDays": 2,
            "ratePerDay": 15000,
            "totalRate": 30000
        },
        {
            "cdId": 2,
            "returnQuantity": 2,
            "rentDays": 2,
            "ratePerDay": 10000,
            "totalRate": 40000
        }
    ]
}
```

### Payment
#### POST - /payment
Request
```
{
    "returnId": 1,
    "userId": 1,
    "returns": [
        {
            "cdId": 1,
            "returnQuantity": 1,
            "rentDays": 2,
            "ratePerDay": 15000,
            "totalRate": 30000
        },
        {
            "cdId": 2,
            "returnQuantity": 2,
            "rentDays": 2,
            "ratePerDay": 10000,
            "totalRate": 40000
        }
    ]
}
```
Response Body (Status: 201 Created)
```
{
    "id": 1,
    "dateTime": "2020-04-21T21:46:54.988508+07:00",
    "returnId": 1,
    "userId": 1,
    "totalPayment": 70000,
    "returns": [
        {
            "cdId": 1,
            "returnQuantity": 1,
            "rentDays": 2,
            "ratePerDay": 15000,
            "totalRate": 30000
        },
        {
            "cdId": 2,
            "returnQuantity": 2,
            "rentDays": 2,
            "ratePerDay": 10000,
            "totalRate": 40000
        }
    ]
}
```
#### GET - /payment/{payment_id}
Example: /payment/1 <br> 
Response Body (Status: 200 OK)
```
{
    "id": 1,
    "dateTime": "2020-04-21T21:46:54.988508Z",
    "returnId": 1,
    "userId": 1,
    "totalPayment": 70000,
    "returns": [
        {
            "cdId": 1,
            "returnQuantity": 1,
            "rentDays": 2,
            "ratePerDay": 15000,
            "totalRate": 30000
        },
        {
            "cdId": 2,
            "returnQuantity": 2,
            "rentDays": 2,
            "ratePerDay": 10000,
            "totalRate": 40000
        }
    ]
}
```
Ensure that quantity of that cd_id (cd_id = 1) is added by return_quantity (return_quantity = 1). <br>
For example, initial quantity of that cd_id (cd_id = 1) in collection is 19. After added by return_quantity (return_quantity = 1), the final quantity is 20. <br>
Check by GET request to /collection/1 <br> Response Body (Status: 200 OK)
```
{
    "id": 1,
    "dateTime": "2020-04-21T22:24:31.37984Z",
    "title": "Star Wars",
    "category": "Sci-Fi",
    "quantity": 20,
    "rate": 15000
}
```
Ensure that quantity of that cd_id (cd_id = 2) is added by return_quantity (return_quantity = 2). <br>
For example, initial quantity of that cd_id (cd_id = 2) in collection is 18. After added by return_quantity (return_quantity = 2), the final quantity is 20. <br>
Check by GET request to /collection/2 <br> Response Body (Status: 200 OK)
```
{
    "id": 2,
    "dateTime": "2020-04-21T22:24:31.436342Z",
    "title": "Captain America",
    "category": "Sci-Fi",
    "quantity": 20,
    "rate": 10000
}
```
