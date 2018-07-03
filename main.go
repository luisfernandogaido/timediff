package main

import (
	"os"
	"log"
	"time"
	"fmt"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		log.Fatal("2 argumentos esperados.")
	}
	hora1, hora2 := args[0], args[1]
	h1, err := time.Parse("15:04", hora1)
	if err != nil {
		log.Fatal(err)
	}
	h2, err := time.Parse("15:04", hora2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(h2.Sub(h1))
}
