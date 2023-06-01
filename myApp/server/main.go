package main

import (
	"fmt"
	"log"
	"myapp/server/handler"
	"myapp/server/store"
	"net"
)

func main() {

	store.NewStore()

	listner, err := net.Listen("tcp4", ":1234")
	if err != nil {
		panic(err) //log
	}

	defer func() { _ = listner.Close() }()

	for {
		fmt.Println("server")
		c, err := listner.Accept()
		if err != nil {
			log.Printf("Accept error %s", err)
		}

		//_, err = c.Write([]byte("123456\n"))
		//_, err = c.Write([]byte("aaa"))
		if err != nil {
			log.Printf("Write error %s\n", err)

		}

		//fmt.Printf("Write: %d\n", n)
		//go handler.Handle(c)

		go func() {
			log.Printf("%s: start conn", c.RemoteAddr())
			defer log.Printf("%s: close conn", c.RemoteAddr())
			for {

				shutdown, err := handler.Handle(c)
				if err != nil {
					break
				}
				if shutdown {
					break
				}
			}
		}()
	}

}
