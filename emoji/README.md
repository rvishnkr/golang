## `emoji` Manifest

* /emoji - Parsing strings into emojis
  * emoji.go - the main package containg a map with the list of all emoji strings agains unicode
  * main.go - to test the emoji package
  * README.md 


## Building the project

### Installing the emoji pakage to your Go installation
* In your computer's Go directory, look for a `/src` folder. Create a folder called `emoji` within this folder and copy the `emoji.go` file here
* Run `go build emoji.go`
* If no errors were reported, run `go install`. This should get the `emoji.a` file setup in you `Go/pkg` folder and the emoji package is now ready for use.

### Testing the emoji package installation
`go build main.go` - to build the excutable

## Usage

* If you haven't built it yet `go run main.go`
* If you have already built main `./main`

