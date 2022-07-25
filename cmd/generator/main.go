package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
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
	limit := 30
	x := r.Intn(limit)
	y := r.Intn(limit)
	var symbols [16]rune = [16]rune{'0', '#', '1', '#', '2', '#', '3', '4', '#', '5', '#', '6', '7', '#', '8', '9'}
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
		}
		log.Printf(row)
		byteSlice := []byte(row + "\n")
		_, err := file.Write(byteSlice)
		if err != nil {
			log.Fatal(err)
		}
	}
}
