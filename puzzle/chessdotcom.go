package puzzle

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode"

	"github.com/notnil/chess"
)

// ChessDotCom represent the adapter for puzzles from chess.com
type ChessDotCom struct {
	httpcli   *http.Client
	baseURL   *url.URL
	timeout   time.Duration
	expires   time.Duration
	expiresAt time.Time
	puzzle    ChessDotComPuzzle
	score     Score
	done      bool
	answers   []*chess.Move
	game      *chess.Game
}

// ChessDotComPuzzle represents the data structure for puzzle from chess.com API\
// ref.: www.chess.com/news/view/published-data-api#pubapi-random-daily-puzzle
type ChessDotComPuzzle struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	PublishTime int64  `json:"publish_time"`
	FEN         string `json:"fen"`
	PGN         string `json:"pgn"`
	Image       string `json:"image"`
}

const (
	ChessDotComURL     = "https://api.chess.com/pub/"
	ChessDotComTimeout = time.Minute
	ChessDotComExpires = 15 * time.Second
)

// NewChessDotCom returns a instance to access puzzles from chess.com
func NewChessDotCom(baseURL string, timeout time.Duration, expires time.Duration) (*ChessDotCom, error) {
	bURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &ChessDotCom{
		baseURL: bURL,
		timeout: timeout,
		expires: expires,
	}, nil
}

// NewGame refresh the puzzle information from source when applicable.
// This endpoint for random puzzle from chess.com has new information
// every 15 seconds, so there is no point in sending request before that.
func (c *ChessDotCom) NewGame() (func(*chess.Game), error) {
	if c.hasExpired() {
		if err := c.get("puzzle/random", nil, &c.puzzle); err != nil {
			return nil, err
		}

		var err error
		c.answers, err = c.gatherAnswers()
		if err != nil {
			return nil, err
		}
	}
	return chess.FEN(c.puzzle.FEN)
}

// Answer returns whether or not a move is correct
func (c *ChessDotCom) Answer(m *chess.Move) bool {
	if m != nil && len(c.answers) > 0 && m.String() == c.answers[0].String() {
		c.answers = c.answers[1:]
		return true
	}
	c.score = FAILURE
	return false
}

// NextMove returns the next move in the answer
func (c *ChessDotCom) NextMove() *chess.Move {
	if len(c.answers) == 0 {
		return nil
	}
	next := c.answers[0]
	c.answers = c.answers[1:]
	return next
}

// Hint returns the next move in the answer
func (c *ChessDotCom) Hint() *chess.Move {
	if len(c.answers) == 0 {
		return nil
	}
	return c.answers[0]
}

func (c *ChessDotCom) Score() Score {
	if c.score == NOSCORE {
		if len(c.answers) == 0 {
			c.score = SUCCESS
			return SUCCESS
		}
	}
	return c.score
}

func (c *ChessDotCom) Done() bool {
	return len(c.answers) == 0
}

func (c *ChessDotCom) get(URL string, querystring url.Values, v interface{}) error {
	nURL, err := url.Parse(URL)
	if err != nil {
		return err
	}

	endpoint := c.baseURL.ResolveReference(nURL)
	endpoint.RawQuery = querystring.Encode()
	resp, err := c.doRequest("GET", endpoint.String(), nil)
	if err != nil {
		return err
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf(
			"chess.com responded unexpectly with %q", resp.Status)
	}

	return c.unmarshal(resp, v)
}

func (c *ChessDotCom) unmarshal(r *http.Response, v interface{}) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}
	return nil
}

func (c *ChessDotCom) doRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	r, err := http.NewRequest(method, endpoint, body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: c.timeout}
	return client.Do(r)
}

func (c *ChessDotCom) hasExpired() bool {
	return c.expiresAt.Before(time.Now())
}

func (c *ChessDotCom) gatherAnswers() ([]*chess.Move, error) {
	var answers []*chess.Move

	pgn := strings.ReplaceAll(c.puzzle.PGN, ".", " ")
	n := strings.Split(pgn, "]")
	if len(n) == 0 {
		return nil, fmt.Errorf("puzzle is missing a correct answer format")
	}

	fen, err := chess.FEN(c.puzzle.FEN)
	if err != nil {
		return nil, err
	}

	game := chess.NewGame(fen)
	for _, m := range strings.Fields(n[len(n)-1]) {
		if len(m) == 0 || !unicode.IsLetter(rune(m[0])) {
			continue
		}
		move, err := chess.AlgebraicNotation{}.Decode(game.Position(), m)
		if err != nil {
			return nil, err
		}

		answers = append(answers, move)
		game.Move(move)
	}

	return answers, nil
}
