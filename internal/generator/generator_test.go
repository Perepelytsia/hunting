package generator

import "testing"

func TestGetSymbolsLen(t *testing.T) {
	amount := CreateField()
	if len(amount) != 144 {
		t.Error("amount should be 144, got", len(amount))
	}
}