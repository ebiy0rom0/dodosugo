package interactor

import (
	"dodosugo/usecase/inputdata"
	"dodosugo/usecase/inputport"
	"dodosugo/usecase/outputport"
	"math/rand"
	"strings"
	"time"
)

type matcher struct {
	out outputport.MatcherOutputport
}

func NewMatcher(o outputport.MatcherOutputport) inputport.MatcherInputport {
	return &matcher{
		out: o,
	}
}

func (m *matcher) MatchingToContinuesText(inject inputdata.InjectMatcher) {
	m.out.Display()
	rand.Seed(time.Now().UnixMicro())

	var stack string
	for {
		next := inject.Words[rand.Intn(len(inject.Words))]
		m.out.Receiver(next)

		stack += next
		if stack == inject.Want {
			break
		}
		if strings.HasPrefix(inject.Want, stack) {
			continue
		}
		m.out.Receiver("\n")
		stack = ""
	}
	m.out.Receiver(inject.Last)
}
