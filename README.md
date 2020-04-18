# CD Rental
Count payment of CD that we rent at CD Rental.

## Setup
1. [go mod](#go-mod)
2. [Unit Test](#unit-test)
3. [Run Program](#run-program)

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
