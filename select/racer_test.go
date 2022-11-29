package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRace(t *testing.T) {

	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(20 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	slowUrl := slowServer.URL
	fastUrl := fastServer.URL
	want := fastUrl
	got := websiteRacer(slowUrl, fastUrl)

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
	slowServer.Close()
	fastServer.Close()

}
