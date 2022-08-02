package inputdata

type InjectMatcher struct {
	Words []string
	Want  string
	Last  string
}

func (i *InjectMatcher) Validation() error {
	return nil
}
