package outputport

type MatcherOutputport interface {
	Receiver(string)
	Display()
}
