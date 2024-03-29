package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Microsecond)

	fastServer := makeDelayedServer(0 * time.Microsecond)
	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
	slowServer.Close()
	fastServer.Close()
}

func makeDelayedServer(delayedTime time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayedTime)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
