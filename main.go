package main

import (
	"fmt"
	"log"
)

func main() {

	l := NewLoader()
	err := l.Scan("./resources")
	if err != nil {
		log.Fatal(err)
	}

	image, err := l.GetImage("assets/fin")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(image)

	fmt.Println("Everything Looks Good!")
}
