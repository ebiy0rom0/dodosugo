package controller

import (
	"dodosugo/usecase/inputdata"
	"dodosugo/usecase/inputport"
)

type matcherController struct {
	matcher inputport.MatcherInputport
}

func NewMatcherContoller(inputport inputport.MatcherInputport) *matcherController {
	return &matcherController{
		matcher: inputport,
	}
}

func (m *matcherController) Dodosugo(words []string, want string, last string) error {
	inject := inputdata.InjectMatcher{
		Words: words,
		Want:  want,
		Last:  last,
	}

	if err := inject.Validation(); err != nil {
		return err
	}

	m.matcher.MatchingToContinuesText(inject)
	return nil
}
