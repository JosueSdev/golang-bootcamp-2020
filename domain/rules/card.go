package rules

type cardFieldNames struct {
	Value string
	Suit  string
	Code  string
}

//CardFieldNames constains the name of the card fields. Can be replaced with reflection.
var CardFieldNames = cardFieldNames{
	"Value",
	"Suit",
	"Code",
}
