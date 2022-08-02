package inputport

import "dodosugo/usecase/inputdata"

type MatcherInputport interface {
	MatchingToContinuesText(inputdata.InjectMatcher)
}
