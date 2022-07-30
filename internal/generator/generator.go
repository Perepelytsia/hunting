package generator

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/perepelytsia/hunting/internal/symbols"
)

func CreateField() int {
	file, err := os.OpenFile(
		"./static/field.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	limit := 16
	x := r.Intn(limit)
	y := r.Intn(limit)
	var symbols [16]rune = symbols.GetSymbols()
	amount := 0
	for j := 0; j < limit; j++ {
		row := ""
		for i := 0; i < limit; i++ {
			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			pos := r.Intn(len(symbols))
			if j == x && i == y {
				row += "x"
			} else {
				row += string(symbols[pos])
			}
			amount++
		}
		log.Printf(row)
		byteSlice := []byte(row + "\n")
		_, err := file.Write(byteSlice)
		if err != nil {
			log.Fatal(err)
		}
	}
	return amount
}
