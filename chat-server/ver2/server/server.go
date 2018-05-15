// A TCP server in Go
// To run: $go run server.go -h
// @authors ravi
 
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"log"
	"io"
	"strings"
	"strconv"
	"time"
	"emoji"
)

var addr = flag.String("addr", "", "The address to listen to; default is \"\" (all interfaces).")
var port = flag.Int("port", 8000, "The port to listen on; default is 8000.")

func main() {
	//Parse command line arguments if any
	flag.Parse()

	users := make(map[net.Conn]string) // Map of active connections
	newConnection := make(chan net.Conn) // Handle new connection
	addedUser := make(chan net.Conn)   // Add new connection
	discUser := make(chan net.Conn)    // Users that have left chat
	messages := make(chan string)      // channel that recieves messages from all users


	fmt.Println("Starting server...")

	src := *addr + ":" + strconv.Itoa(*port)
	listener, err := net.Listen("tcp", src)
    if err != nil {
    	fmt.Println(err)
        os.Exit(1)
    }
	fmt.Printf("Listening on %s.\n", src)

	// Close the listener when the application closes.
	defer listener.Close()

	go func() { // Launch routine that will accept connections
        for {
            conn, err := listener.Accept()
            if err != nil {
                fmt.Println(err)
                os.Exit(1)
            }
            newConnection <- conn // Send to handle new user
        }
    }()

	// Run forever
    for { 
                                                                                                                                                     
        select {
        case conn := <-newConnection:

            go func(conn net.Conn) { // Ask user for name and information
                reader := bufio.NewReader(conn)
                // io.WriteString(conn, "Enter name: ")
                userName, _ := reader.ReadString('\n')
                userName = strings.Trim(userName, "\r\n")
                log.Printf("Accepted new user : %s", userName)
                messages <- fmt.Sprintf("Accepted user : [%s]\n\n", userName)

                users[conn] = userName // Add connection

                addedUser <- conn // Add user to pool
            }(conn)

        case conn := <-addedUser: // Launch a new go routine for the newly added user

            go func(conn net.Conn, userName string) {
                reader := bufio.NewReader(conn)
                for { // Run forever and handle this user's messages
                    newMessage, err := reader.ReadString('\n')
                    newMessage = strings.Trim(newMessage, "\r\n")
                    if err != nil {
                        break
                    }

                    if strings.Compare("/help", newMessage) == 0{
                    	newMessage = "\n/help\t\tdisplay a help message\n/time\t\tdisplay the current time\n/clear\t\tclear the client screen\n:emoji-name:\tcorresponding emoji" 
                    }
                    if strings.Compare("/qs", newMessage) == 0{
                        os.Exit(1)
                    }
        		    if strings.Compare("/time", newMessage) == 0{
        		    	newMessage = "It is " + time.Now().String() + "\n"
        		    }
        		    if strings.Compare("/stats", newMessage) == 0{
        		    	newMessage = "\nList of connected users:"
        		    	for k, v := range users {
        		    		log.Println("Connection k: ", k)
       						newMessage += "\nusername: " + v
    					}
                    }

                    // Send to messages channel to all users after processing any emojis
                    messages <- fmt.Sprintf(">>%s: %s \a\n\n", userName, emoji.Emojify(newMessage))
                }

                discUser <- conn // If error occurs, connection has been terminated
                messages <- fmt.Sprintf("%s disconnected\n\n", userName)
            }(conn, users[conn])

        // If a message is recieved from any user
        case message := <-messages:
            for conn, _ := range users { // Send to all users
                go func(conn net.Conn, message string) { // Write to all user connections
                        _, err := io.WriteString(conn, message)
                        if err != nil {
                            discUser <- conn
                        }
                }(conn, message)
                log.Printf("New message:\n %s", message)
                log.Printf("Sent to %d users", len(users))
            }

        // Handle disconnected users
        case conn := <-discUser:
            log.Printf("Client disconnected")
            delete(users, conn)
        }
    }

}
