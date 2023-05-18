package context

import (
	"fmt"
	"net/http"
	"time"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		fmt.Println(ctx.Value("test"))

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()
		select {
		case d := <-data:
			fmt.Println("response===")
			fmt.Fprint(w, d)
		case <-ctx.Done():
			fmt.Println("cancel===")
			store.Cancel()
		}
		// fmt.Fprint(w, store.Fetch())
	}
}

type Store interface {
	Fetch() string
	Cancel()
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) Fetch() string {
	time.Sleep(4200 * time.Millisecond)
	fmt.Println("response")
	return s.response
}
