package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Person Structure.
type Person struct {
	First string
}

func main() {
	p1 := Person{
		First: "Salvador",
	}

	p2 := Person{
		First: "Joe",
	}

	// Create slice of person's.
	xp := []Person{p1, p2}

	bs, err := json.Marshal(xp)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(bs))
}
