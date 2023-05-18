package context

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StuStore struct {
	response string
}

func (s *StuStore) Fetch() string {
	return s.response
}

func TestServer(t *testing.T) {
	data := "hello world"
	store := &SpyStore{response: data}
	svr := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	// ctx := context.WithValue(request.Context(), "test", "111111111")
	cancellingCtx, cancel := context.WithCancel(request.Context())
	time.AfterFunc(5*time.Second, cancel)
	// ctx := context.WithValue(cancellingCtx, "test", "wwwww")
	request = request.WithContext(cancellingCtx)

	fmt.Println("send request")
	go func() {
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("cancel from client")
		// cancel()
	}()

	respone := httptest.NewRecorder()
	svr.ServeHTTP(respone, request)
	if respone.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, respone.Body.String(), data)
	}
	if store.cancelled {
		t.Error("it should not have cancelled the store")
	}
}
