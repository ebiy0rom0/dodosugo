package presenter

import (
	"dodosugo/usecase/outputport"
	"fmt"
	"io"
)

type matcherPresenter struct {
	writer io.Writer
	ch     chan string
}

func NewMatcherPresenter(w io.Writer, c chan string) outputport.MatcherOutputport {
	return &matcherPresenter{
		writer: w,
		ch:     c,
	}
}

func (m *matcherPresenter) Display() {
	go func() {
		for {
			s := <-m.ch
			fmt.Fprint(m.writer, s)
		}
	}()
}

func (m *matcherPresenter) Receiver(s string) {
	m.ch <- s
}
