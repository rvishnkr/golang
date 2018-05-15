## ver1 Manifest

* /ver1 - A basic TCP server
  * server.go - server
  * README.md - build and usage instructions

## Building the project

From within the `/ver1` folder, `go build server.go` - to build the executable

## Usage

### Server: 
* If you haven't built it yet, from within the `/ver1` folder, run `go run server.go`
* If you have already built the server, from within the `/ver1` folder, run `./server`

### Client:
* Make sure you have `telnet` installed. 
* From a different terminal/console, from within the `/ver1` folder, run `telnet localhost 5005`
* To exit the client, use the escape sequence `^]` and then enter `q` and press return
