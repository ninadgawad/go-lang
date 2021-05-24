# go-lang
This document will demonstrate development using Go package 

## First Go Program 
- Code is organized into packages
- Package is collection of source files 
- Go 1.13 or later and GO111MODULE environment variable is not set

### Hello.mod 
- Hello World in Go 

[Go Code](https://golang.org/doc/code) 


## Installing GoLang 

#### Step1: 
[Download Golang](https://golang.org/doc/install?download=go1.16.3.windows-amd64.msi)
Validate version `go version`
<br />
`go version go1.16.3 windows/amd64`

#### Step2: Go Code
`go mod init example.com/hello`
____________________________________
`go run .`
____________________________________


#### Step3: Installation
1. MOD - module maintenance
2. ENV - print Go environment information
3. RUN - compile and run Go program
4. FMT - gofmt (reformat) package sources
5. GET - add dependencies to current module and install them
6. GoROOT - Stores Binaries 
7. GoWorkspace - Stores your code 

#### General Project Structure 
```go
project
    bin
    pkg
    src
      github.com
        ninadgawad
          helloWorld
```          

