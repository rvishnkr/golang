// A TCP client written in Go.
// @authors ravi

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
	"emoji"
)

var host = flag.String("host", "localhost", "The hostname or IP to connect to; defaults to \"localhost\".")
var port = flag.Int("port", 8000, "The port to connect to; defaults to 8000.")

func main() {
	flag.Parse()

	dest := *host + ":" + strconv.Itoa(*port)
	fmt.Printf("Connecting to %s...\n", dest)

	//Connect to a server using `Dial`
	conn, err := net.Dial("tcp", dest)

	if err != nil {
		if _, t := err.(*net.OpError); t {
			fmt.Println("Some problem connecting.")
		} else {
			fmt.Println("Unknown error: " + err.Error())
		}
		os.Exit(1)
	}
	fmt.Print("Enter name: ")
	scanner := bufio.NewScanner(conn)
	uName := scanner.Text()
	_, err = conn.Write([]byte(uName))
	if err != nil {
		fmt.Println("Error writing to stream.")
		os.Exit(1)
	}


	go readConnection(conn)

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "/quit\n"{
			os.Exit(0)
		}
		if text == "/clear\n"{
			print("\033[H\033[2J")
		}

		conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Error writing to stream.")
			break
		}
	}
}

func readConnection(conn net.Conn) {

	for {
		scanner := bufio.NewScanner(conn)

		for {
			ok := scanner.Scan()
			text := scanner.Text()

			fmt.Println("", text)

			if !ok {
				fmt.Println(emoji.Emojify(":heavy_exclamation_mark:Yo! Server quit unexpectedly.:heavy_exclamation_mark:"))
				os.Exit(1)
			}
		}
	}
}
