package rules

type status struct {
	win  string
	foul string
	ok   string
}

//Status abstracts methods that return status strings
type Status interface {
	Win() string
	Foul() string
	Ok() string
}

//EnglishStatus is a Satus implementation with messages in english
var EnglishStatus Status = status{
	win:  "Blackjack!",
	foul: "Foul",
	ok:   "Ok",
}

//Win returns the status message corresponding to a player getting exactly 21 points
func (s status) Win() string {
	return s.win
}

//Foul returns the status message corresponding to a player getting more than 21 points
func (s status) Foul() string {
	return s.foul
}

//Ok returns the status message corresponding to a player getting less than 21 points
func (s status) Ok() string {
	return s.ok
}
