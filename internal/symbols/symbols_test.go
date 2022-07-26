package symbols

import "testing"

func TestGetSymbols(t *testing.T) {
	symbols := GetSymbols()
	if len(symbols) != 16 {
		t.Error("result should be 16 symbols, got", len(symbols))
	}
}
