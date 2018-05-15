## Manifest

* /ver1 - A basic TCP server
  * server.go - server
  * README.md - build and usage instructions
* /ver2 - Extending the basic TCP server functionality
  * /server/server.go - server
  * /client/client.go - client
  * README.md - build and usage instructions
* /ver3 - Using Gorilla WebSockers
  * /server/main.go - server
  * /web-client/index.html - web client
  * /web-client/style.css - web client styling
  * /web-client/app.js - helper functions
  * README.md - build and usage instructions

## Build and Usage

Build and usage instructions for each version is in the corresponding folders' README.md file. All build instructions assume that Golang has been installed and the machine is Unix based. Errors if any could be because a port that the servers needs to listen on is already in use.

## Sources used

* [The Go Programming Language] (https://golang.org/)
* [Go by Example] (https://gobyexample.com/)
* [GitHub - gorilla/websocket: A WebSocket implementation for Go.] (https://github.com/gorilla/websocket)
* [Stack Overflow] (https://stackoverflow.com/)
