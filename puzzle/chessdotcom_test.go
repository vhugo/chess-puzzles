package puzzle_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/notnil/chess"
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

func (m *httpmock) PuzzleOne() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{
			"title": "Gyorgy BAKCSI, Jubilli de Finlande 1966",
			"url": "https://www.chess.com/forum/view/daily-puzzles/10-22-2018-gyorgy-bakcsi-jubilli-de-finlande-1966",
			"publish_time": 1540191600,
			"fen": "2b4r/6N1/4Rp2/4n1bQ/5k2/3rN2p/5P1K/6BB w k - 0 1",
			"pgn": "[Date \"????.??.??\"]\r\n[FEN \"2b4r/6N1/4Rp2/4n1bQ/5k2/3rN2p/5P1K/6BB w k - 0 1\"]\r\n\r\n1. Kxh3 Bxe6+ 2. Nxe6#\r\n*",
			"image": "https://www.chess.com/dynboard?fen=2b4r/6N1/4Rp2/4n1bQ/5k2/3rN2p/5P1K/6BB%20w%20k%20-%200%201&size=2"
		}`))

		return
	}
	return httptest.NewServer(
		http.HandlerFunc(f),
	)
}
func (m *httpmock) PuzzleTwo() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{
			"title": "Dimitri Gurevich Tactic",
			"comments": "",
			"url": "https://www.chess.com/forum/view/daily-puzzles/9192008---dimitri-gurevich-tactic",
			"publish_time": 1221807600,
			"fen": "2r2rk1/ppBQ1pp1/7p/6q1/4n3/4nB2/PPP3PP/1K2R2R b - - 0 1",
			"pgn": "[Date \"????.??.??\"]\r\n[Result \"*\"]\r\n[FEN \"2r2rk1/ppBQ1pp1/7p/6q1/4n3/4nB2/PPP3PP/1K2R2R b - - 0 1\"]\r\n\r\n1...Rxc7 2. Qxc7 Nd2+ 3. Kc1 Nec4 4. Rd1 Nb3+ 5. Kb1 Qc1+ 6. Rxc1 Ncd2#\r\n*",
			"image": "https://www.chess.com/dynboard?fen=2r2rk1/ppBQ1pp1/7p/6q1/4n3/4nB2/PPP3PP/1K2R2R%20b%20-%20-%200%201&size=2"
		}`))

		return
	}
	return httptest.NewServer(
		http.HandlerFunc(f),
	)
}
func (m *httpmock) PuzzleThree() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{
			"title": "Going to the Edge and Back",
			"comments": "",
			"url": "https://www.chess.com/forum/view/daily-puzzles/862011---going-to-the-edge-and-back",
			"publish_time": 1312614000,
			"fen": "1r3k2/5ppB/2pbb3/p5BQ/7P/6P1/Pq3P2/5RK1 w - - 0 1",
			"pgn": "[Date \"????.??.??\"]\r\n[Result \"*\"]\r\n[FEN \"1r3k2/5ppB/2pbb3/p5BQ/7P/6P1/Pq3P2/5RK1 w - - 0 1\"]\r\n\r\n1. Bf5 g6 2. Qh6+ Qg7 3. Bxe6 fxe6 4. Be7+ Bxe7 5. Qf4+ Qf6 6. Qxb8+\r\n*",
			"image": "https://www.chess.com/dynboard?fen=1r3k2/5ppB/2pbb3/p5BQ/7P/6P1/Pq3P2/5RK1%20w%20-%20-%200%201&size=2"
		}`))

		return
	}
	return httptest.NewServer(
		http.HandlerFunc(f),
	)
}

func TestChessDotCom(t *testing.T) {
	t.Run("new chess dot com", func(t *testing.T) {
		for _, tc := range []struct {
			m, url string
			ok     bool
		}{
			{"invalid url", "--\ninvalid\n--", false},
			{"valid url", "www.chess.com", true},
		} {
			t.Run(tc.m, func(t *testing.T) {
				_, err := puzzle.NewChessDotCom(tc.url, time.Minute, time.Millisecond)
				ok := (err == nil)
				if tc.ok != ok {
					t.Fatalf("got %v, wanted %v for err %q", ok, tc.ok, err)
				}
			})
		}
	})
}
func TestChessDotComPuzzle(t *testing.T) {
	var mock httpmock

	type expected struct {
		answer bool
		score  puzzle.Score
		done   bool
	}

	for _, tc := range []struct {
		m        string
		url      string
		moves    []string
		expected expected
	}{
		{
			m:   "puzzle one",
			url: mock.PuzzleOne().URL,
			moves: []string{
				"Kxh3", "Bxe6+", "Nxe6#",
			},
			expected: expected{
				answer: true,
				score:  puzzle.SUCCESS,
				done:   true,
			},
		},
		{
			m:   "puzzle two",
			url: mock.PuzzleTwo().URL,
			moves: []string{
				"Rxc7", "Qxc7", "Nd2+", "Kc1", "Nec4", "Rd1", "Nb3+",
				"Kb1", "Qc1+", "Rxc1", "Ncd2#",
			},
			expected: expected{
				answer: true,
				score:  puzzle.SUCCESS,
				done:   true,
			},
		},
		{
			m:   "puzzle three",
			url: mock.PuzzleThree().URL,
			moves: []string{
				"Bf5", "g6", "Qh6+", "Qg7", "Bxe6", "fxe6", "Be7+",
				"Bxe7", "Qf4+", "Qf6", "Qxb8+",
			},
			expected: expected{
				answer: true,
				score:  puzzle.SUCCESS,
				done:   true,
			},
		},
	} {
		t.Run(tc.m, func(t *testing.T) {
			var err error
			puzzler, err := puzzle.NewChessDotCom(tc.url, time.Minute, time.Millisecond)
			if err != nil {
				t.Fatal("unexpected error: ", err)
			}

			var gc *chess.Game
			var newPuzzle func(*chess.Game)

			newPuzzle, err = puzzler.NewGame()
			if err != nil {
				t.Fatal("unexpected error: ", err)
			}

			gc = chess.NewGame(newPuzzle)
			for _, m := range tc.moves {
				move, err := chess.AlgebraicNotation{}.Decode(gc.Position(), m)
				if err != nil {
					t.Fatal("unexpected error: ", err)
				}

				IsAnswer := puzzler.Answer(move)
				if IsAnswer != tc.expected.answer {
					t.Fatalf("got %v, want %v", IsAnswer, tc.expected.answer)
				}

				if err := gc.Move(move); err != nil {
					t.Fatal("unexpected error: ", err)
				}
			}

			if puzzler.Score() != tc.expected.score {
				t.Fatalf("got %v, want %v", puzzler.Score(), tc.expected.score)
			}

			if puzzler.Done() != tc.expected.done {
				t.Fatalf("expected puzzle to be done, but it was not")
			}

			newPuzzle, err = puzzler.NewGame()
			if err != nil {
				t.Fatal("unexpected error: ", err)
			}

			gc = chess.NewGame(newPuzzle)
			for _, m := range tc.moves {
				expectedNextMove, err := chess.AlgebraicNotation{}.Decode(gc.Position(), m)
				if err != nil {
					t.Fatal("unexpected error: ", err)
				}

				nextMove := puzzler.NextMove()
				if expectedNextMove.String() != nextMove.String() {
					t.Fatalf("got %v, want %v", nextMove, expectedNextMove)
				}

				if err := gc.Move(nextMove); err != nil {
					t.Fatal("unexpected error: ", err)
				}
			}

			if puzzler.Done() != tc.expected.done {
				t.Fatalf("expected puzzle to be done, but it was not")
			}
		})
	}
}
