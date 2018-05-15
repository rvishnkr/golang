## ver3 Manifest

* /ver3 - Using Gorilla WebSockers
  * /server/main.go - server
  * /web-client/index.html - web client
  * /web-client/style.css - web client styling
  * /web-client/app.js - helper functions
  * README.md - build and usage instructions

## Building the project

* from within the `ver3/server/` folder run `go build server.go` to build the server excutable

## Usage

### Server: 
* If you haven't built it yet, from within the server directory, run `go run server.go`
* If you have already built the server, from the server directory, run`./server`

### Client:
* In a web browser, after the server is running, navigate to `http://localhost:8000/`
