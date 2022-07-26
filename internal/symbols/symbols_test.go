package symbols

import "testing"

func TestGetSymbolsLen(t *testing.T) {
	symbols := GetSymbols()
	if len(symbols) != 16 {
		t.Error("result should be 16 symbols, got", len(symbols))
	}
}

func TestGetSymbolsWall(t *testing.T) {
	fail := true
	for _, symbol := range GetSymbols() {
		if symbol == '#' {
			fail = false
			break
		}
	}
	if fail {
		t.Error("there is no #, got emptyness")
	}
}
