package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)

var pool map[string]*net.Conn

func main() {
	
	pool = make(map[string]*net.Conn)
	ln, err := net.Listen("tcp", ":8085")

	if err != nil {
		fmt.Println("Error occured while listening")
		return
	}

	fmt.Println("Listening on port 8085")

	for {
		//Accept Connection
		conn, err := ln.Accept()
		if (err != nil) {
			fmt.Println("Error occured while accepting connection")
		}
		//Receive Public Key of Client
		pkey, err := bufio.NewReader(conn).ReadString('\n')
		//If public key received only then accept the connection
		if (err == nil) {
			pkey = strings.Trim(pkey, "\n")
			_, exists := pool[pkey]

			if (exists) {
				conn.Write([]byte("> pkey exists \n"))
			} else {
				pool[pkey] = &conn
				go handleConnection(&conn, &pool, pkey)
			}
		}
	}
}

func handleConnection(conn *net.Conn, pool *map[string]*net.Conn, connPubKey string) {
	(*conn).Write([]byte("> You are connected \n"))
	for {
		input, err := bufio.NewReader(*conn).ReadString('\n')

		if (err != nil) {
			fmt.Println("Error occured from client")
			delete(*pool, connPubKey)
			return
		}

		input = strings.Trim(input, "\n")
		values := strings.Split(input, " ")

		pkey := values[0]
		values[0] = connPubKey + " > "

		input = strings.Join(values, " ")

		_, exists := (*pool)[pkey]

		if (!exists) {
			(*(*pool)[connPubKey]).Write([]byte("Invalid public key\n"))
		} else {
			(*(*pool)[pkey]).Write([]byte(input + "\n"))
		}
	
	}
}