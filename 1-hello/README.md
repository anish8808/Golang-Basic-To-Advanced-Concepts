- To run go
```go run hello.go```

-To know where the binary/exe created when run 
```go run --work hello.go```

- To build and gen a binary/exe
```go build hello.go```

- To build and gen a binary/exe with name 
```go build -o app hello.go```

- To compile the binary 
``` ./app ```

- To build for other operating system(cross compilation)
``` export GOOS=windows```
 ```export GOARCH=amd64```
 ``` go build -o app.exe hello.go```

- To check the go env
``` go to terminal and type go env ```

- To know list of OS and architechure
```go tool dist list```

-To build for other operating system
```GOOD=linux GOARCH=amd46```

-To check the binary is sutaible for os 
``` file $binary name ```

-To install binaries/exe to GOBIN location 
```go install hello.go```

-To set bin path and generate binary there 
``` GOBIN=/Users/anishgond/Documents/GolangCourse/1-hello/bin go install hello.go```
```instead of hello.go we can use github path or somthing this will generate binary in the bin path ```

-To install binaries/exe from from remote repositories
```go install github.com/JitenPalaparthi/urllinter@latest```

-To complie and link
```go tool compile hello.go```
```go tool compile link hello.exe hello.o```


- go mod --> 1. To work with user defined package
             2. To maintaine= dependencies 
    ```go mod init demo``` --> initialise the new module 
    ```go mod tidy ``` --> it uses and fetch the updated module from internet repo  
