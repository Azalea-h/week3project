package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Please provide an address!")
	}

	msg := []byte("put13key212stored value\n")
	conn, err := net.Dial("tcp", os.Args[1])
	if err != nil {
		panic(err)
	}
	defer func() { _ = conn.Close() }()
	fmt.Println("client")

	_, err = conn.Write(msg)
	if err != nil {
		log.Printf("Read error %s", err)
	}

	s := bufio.NewScanner(conn)
	s.Scan()
	fmt.Printf("Read message from server: %s\n", s.Text())

	/*
		_, err = conn.Read(msg)
		if err != nil {
			log.Printf("Write error %s", err)
		}
		log.Printf("Read message from server: %s\n", msg)
		//_ = conn.Close()
	*/

}
