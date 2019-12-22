package puzzle_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/vhugo/chess-puzzles/puzzle"
)

type httpmock struct{}

func (m *httpmock) NotFound() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	return httptest.NewServer(
		http.HandlerFunc(f),
	)
}

func (m *httpmock) FailedJSON() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid"))
		return
	}
	return httptest.NewServer(
		http.HandlerFunc(f),
	)
}

func (m *httpmock) GoodResponse() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{
   "title": "title",
   "url": "URL",
   "publish_time": 1513584000,
   "fen": "FEN",
   "pgn": "PGN",
   "image":"the link to the image"
}`))
		return
	}
	return httptest.NewServer(
		http.HandlerFunc(f),
	)
}

func TestChessDotCom(t *testing.T) {
	var mock httpmock

	t.Run("new chess dot com", func(t *testing.T) {
		for _, tc := range []struct {
			url string
			ok  bool
		}{
			{"--\ninvalid\n--", false},
			{mock.NotFound().URL, false},
			{mock.FailedJSON().URL, false},
			{mock.GoodResponse().URL, true},
		} {
			_, err := puzzle.NewChessDotCom(tc.url, time.Minute, time.Millisecond)
			ok := (err == nil)
			if tc.ok != ok {
				t.Fatalf("got %v, wanted %v", ok, tc.ok)
			}
		}
	})

	t.Run("fen", func(t *testing.T) {
		p, err := puzzle.NewChessDotCom(
			mock.GoodResponse().URL,
			time.Minute,
			time.Millisecond)
		if err != nil {
			t.Fatalf("no error expected, got %q", err)
		}

		if p.FEN() != "FEN" {
			t.Fatalf("got %q, want %q", p.FEN(), "FEN")
		}
	})
}
