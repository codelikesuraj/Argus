package argus

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	STR_MODE = iota
	JSON_MODE
)

type Argus struct {
	output io.Writer
	mode   int
}

func Init() *Argus {
	return &Argus{
		output: os.Stdout,
		mode:   STR_MODE,
	}
}

func (a *Argus) SetMode(mode int) {
	if mode < 0 || mode > 1 {
		a.SetMode(STR_MODE)
	} else {
		a.SetMode(mode)
	}
}

func (a *Argus) SetOutput(w io.Writer) {
	a.output = w
}

func (a *Argus) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func(t time.Time, r *http.Request) {
			fmt.Fprintln(a.output, "Argus:", r.Method, r.Response.StatusCode, r.URL.RawQuery, time.Since(t).Abs().String())
		}(time.Now(), r)

		next.ServeHTTP(w, r)
	})
}
