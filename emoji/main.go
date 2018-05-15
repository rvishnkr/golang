//After installing the emoji package, run this to test
// @author ravi

package main

import (
    "flag"
    "fmt"

    "emoji"
)

//'flag' is used for command line parsing
//String defines a string flag with specified name, default value, and usage message. 
var messagePtr = flag.String("m", "CS354 :computer: science project. \nGo :horse:", "the message to emojitize")
//The return value is the address of a string variable that stores the value of the flag.


func main() {
    //Parse the command line argument if any
    flag.Parse()

    //Test
    fmt.Println(emoji.Emojify(*messagePtr))
    fmt.Println(emoji.Emojify(":beer:"))
    fmt.Println(emoji.Emojify(":grin:"))
}
