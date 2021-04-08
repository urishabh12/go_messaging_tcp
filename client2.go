package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8085")
	if(err != nil) {
		// handle error
		return
	}
	pkey := "12346"
	conn.Write([]byte(pkey + "\n"))
	fmt.Printf(pkey + "\n")
	go readFromServer(&conn)
	inputReader := bufio.NewReader(os.Stdin)
	
	for {
		//Receivers public key and msg
		input, err := inputReader.ReadString('\n')

		if(err != nil) {
			fmt.Println("Error occured while reading input from console")
			return
		}

		conn.Write([]byte(input + "\n"))
	}
	
}

func readFromServer(conn *net.Conn) {
	msg, err := bufio.NewReader(*conn).ReadString('\n')

	if (err != nil) {
		fmt.Println("Error occured from server")
		return
	}

	fmt.Println(msg)

	for {
		msg, err := bufio.NewReader(*conn).ReadString('\n')

		if(err != nil) {
			fmt.Println("Error occured from server")
			return
		}

		fmt.Print(msg)
	}
}