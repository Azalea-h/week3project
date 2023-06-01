package handler

import (
	"errors"
	"fmt"
	"io"

	"myapp/server/store"
	"net"
	"strconv"
)

var (
	ErrInvalidProtocol = errors.New("Invalid protocol")
)

func Handle(c net.Conn) (bool, error) {

	request, err := Decode(c)
	if err != nil {
		fmt.Println("print err")
		c.Write([]byte("err"))
		return false, err

	} else {

		store.AddRequestToRequestChannel(request)
		response := <-request.ResponseChannel
		if response.Error == store.ErrKeyNotFound {
			c.Write([]byte("nil"))

		} else if response.Error != nil {
			c.Write([]byte("err")) // check more errors

		}
		if request.Command == store.PUT || request.Command == store.DELETE {
			c.Write([]byte("ack"))

		} else if request.Command == store.GET {
			secondArgument := fmt.Sprintf("%d", len(response.Value))
			firstArgument := len(secondArgument)
			responseString := fmt.Sprintf("val%d%v%s", firstArgument, secondArgument, response.Value)
			fmt.Println(response.Value)

			fmt.Println(firstArgument)

			fmt.Println(secondArgument)
			fmt.Println(responseString)
			c.Write([]byte(responseString)) //respond ack, nil etc
		} else if request.Command == store.BYE {
			return true, nil
		}
	}
	return false, nil
}

func ReadBytes(c io.Reader, i int) string {
	result := make([]byte, i)
	c.Read(result)
	return string(result)
}

func Decode(c io.Reader) (store.KvRequest, error) {
	//put 1 3 key 2 12 stored value
	command := ReadBytes(c, 3)

	r := store.KvRequest{
		ResponseChannel: make(chan store.KvResponse),
		Error:           nil,
	}

	key, err := getArgument(c)
	if err != nil {
		fmt.Println("decode error invalid protocol")
		return store.KvRequest{}, ErrInvalidProtocol
	}

	r.Key = key

	switch command {
	case "bye":
		r.Command = store.BYE
	case "get":
		r.Command = store.GET
	case "del":
		r.Command = store.DELETE
	case "put":
		r.Command = store.PUT
		r.Value, err = getArgument(c)
		if err != nil {
			return store.KvRequest{}, ErrInvalidProtocol
		}
	default:
		return store.KvRequest{}, ErrInvalidProtocol
	}
	return r, nil

}

func getArgument(c io.Reader) (string, error) { // add errors
	first := ReadBytes(c, 1)
	firstConverted, err := strconv.Atoi(first)
	if err != nil {
		return "", ErrInvalidProtocol
	}
	second := ReadBytes(c, firstConverted)
	fmt.Println("second is ", second)
	secondConverted, err := strconv.Atoi(second)
	if err != nil {
		fmt.Println("inside if error: ")
		return "", ErrInvalidProtocol
	}
	return ReadBytes(c, secondConverted), nil
}
