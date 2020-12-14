package rules

type cardFieldNames struct {
	Value string
	Suit  string
	Code  string
}

type cardSpecialValue struct {
	Ace   string
	Jack  string
	Queen string
	King  string
}

//CardFieldNames constains the name of the card fields. Can be replaced with reflection.
var CardFieldNames = cardFieldNames{
	"Value",
	"Suit",
	"Code",
}

//CardSpecialValue contains the name of the supported non-numerical cards
var CardSpecialValue = cardSpecialValue{
	Ace:   "ACE",
	Jack:  "JACK",
	Queen: "QUEEN",
	King:  "KING",
}
