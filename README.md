# Golang

## go mod 

```
 go mod init demo
 ```

 ## to run

 ```
 go run main.go
 ```

  ## to build

 ```
 go build main.go
 ```

## to build with output name

 ```
go build -o hello main.go
 ```

## bring down the size of the binary/executable using build command
```
go build -ldflags="-s -w" -o small-hello main.go
```
## to check the size
```
ls -lh *
or
du -sh *
```

## to do escape analysis

```
go build -gcflags="-m" main.go
```

## to display all os and arch
```
go tool dist list
```

# cross compilation

```
GOOS=linux GOARCH=amd64 go build -o demo-linux  main.go
```

## go compile

```
go tool compile -o output.o main.go
```

## link 

```
go tool link -o demo output.o
```
