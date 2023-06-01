package store

import (
	"errors"
	"fmt"
)

var myStore store

const (
	GET    = "get"
	PUT    = "put"
	DELETE = "del"
	BYE    = "bye"
)

var (
	ErrKeyNotFound = errors.New("Key not found")
)

type store struct {
	kvStore        map[string]string
	requestChannel chan KvRequest
}

type KvRequest struct {
	Command         string
	Key             string
	Value           string
	ResponseChannel chan KvResponse
	Error           error
}

type KvResponse struct {
	Value string
	Error error
}

func NewStore() {
	myStore = store{make(map[string]string), make(chan KvRequest)}
	go myStore.monitor()
}

func (store *store) monitor() {
	for {
		response := KvResponse{Error: nil}
		request := <-myStore.requestChannel

		switch request.Command {

		case BYE:

		case GET:
			v, err := getValue(request.Key)
			response.Error = err
			response.Value = v

		case PUT:
			putEntry(request.Key, request.Value)

		case DELETE:
			deleteEntry(request.Key)
		}

		request.ResponseChannel <- response
	}
}

func AddRequestToRequestChannel(request KvRequest) {
	myStore.requestChannel <- request
}

func getValue(key string) (string, error) {
	v, ok := myStore.kvStore[key]

	if !ok {
		fmt.Println("key not found in getValue")
		return "", ErrKeyNotFound
	}

	return v, nil
}

func putEntry(key string, value string) {
	myStore.kvStore[key] = value
}

func deleteEntry(key string) error {
	_, ok := myStore.kvStore[key]
	if !ok {
		return ErrKeyNotFound
	}
	delete(myStore.kvStore, key)
	return nil
}
