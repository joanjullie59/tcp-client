package main

import (
	"bufio"
	"fmt"
	"net"
)

//1. Dial server
//2. Create a function that reads from the terminal and sends to server
//3. Create another function to recieve and pring msgs from server

func main() {
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		fmt.Println("Error accepting:", err.Error())
	}
	go receive(conn)
	go scanFromStdin(conn)
	defer conn.Close()
	select {}
}

func scanFromStdin(conn net.Conn) {
	for {
		var txt string
		fmt.Print("> ")
		_, err := fmt.Scanln(&txt)
		if err != nil {
			break
		}

		if txt == "\n" || txt == "" {
			continue
		}

		b := []byte(txt)
		_, err = conn.Write(b)
		if err != nil {
			fmt.Println("Failed to send data to the server!")
			break
		}
	}
}
func receive(conn net.Conn) {
	rd := bufio.NewReader(conn)
	var inbuf [64]byte
	for {
		n, err := rd.Read(inbuf[:])
		if err != nil {
			fmt.Println("Error reading from the connection")
			break
		}
		fmt.Println(string(inbuf[:n]))
	}
	conn.Close()
}
