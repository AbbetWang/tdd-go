package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20000 * time.Millisecond)

	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeDelayedServer(delayedTime time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delayedTime)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
