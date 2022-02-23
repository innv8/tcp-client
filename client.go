package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	serverLink := "localhost:5000"
	c, err := net.Dial("tcp", serverLink)
	if err != nil {
		fmt.Println(err)
		return
	}

	// for {
	// 	// reading the input from the terminal
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print(">> ")
	// 	text, _ := reader.ReadString('\n')

	// 	// sending the data
	// 	fmt.Fprintf(c, text+"\n")

	// 	message, _ := bufio.NewReader(c).ReadString('\n')
	// 	fmt.Print("->: " + message)
	// 	if strings.TrimSpace(string(text)) == "STOP" {
	// 		fmt.Println("TCP client exiting...")
	// 		return
	// 	}
	// }

	counter := 0
	for {
		counter++
		fmt.Println("sending Hello world", counter)
		fmt.Fprintf(c, "Hello world %d\n", counter)
		time.Sleep(time.Second)
	}
}
