package rules

import "testing"

var testStatus Status = status{
	win:  "win",
	foul: "foul",
	ok:   "ok",
}

func TestWin(t *testing.T) {
	if testStatus.Win() != "win" {
		t.Fail()
	}
}

func TestFoul(t *testing.T) {
	if testStatus.Foul() != "foul" {
		t.Fail()
	}
}

func TestOk(t *testing.T) {
	if testStatus.Ok() != "ok" {
		t.Fail()
	}
}
