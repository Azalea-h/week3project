package handler_test

import (
	"myapp/server/handler"
	"myapp/server/store"
	"testing"
)

func TestDecode(t *testing.T) {
	// <setup code>
	t.Run("bye", func(t *testing.T) {
		v, err := handler.Decode("bye")

		request := store.KvRequest{
			Command:         "bye",
			ResponseChannel: make(chan store.KvResponse),
			Error:           nil,
		}

		if v != request || err != nil {
			t.Error("Expected", request, "got", v)
		}
	})

	t.Run("put13key212stored value", func(t *testing.T) {
		v, err := handler.Decode("put13key212stored value")

		// request := store.KvRequest{
		// 	Command:         "put",
		// 	Key:             "key",
		// 	Value:           "stored value",
		// 	ResponseChannel: make(chan store.KvResponse),
		// 	Error:           nil,
		// }

		if v.Command != "put" || err != nil {
			t.Error("Expected", "put", "got", v.Command)
		}
	})

	t.Run("get11a", func(t *testing.T) {
		v, err := handler.Decode("get11a")

		request := store.KvRequest{
			Command:         "get",
			Key:             "a",
			ResponseChannel: make(chan store.KvResponse),
			Error:           nil,
		}

		if v != request || err != nil {
			t.Error("Expected", request, "got", v)
		}
	})

	t.Run("del11a", func(t *testing.T) {
		v, err := handler.Decode("del11a")

		request := store.KvRequest{
			Command:         "del",
			Key:             "a",
			ResponseChannel: make(chan store.KvResponse),
			Error:           nil,
		}

		if v != request || err != nil {
			t.Error("Expected", request, "got", v)
		}
	})

}
