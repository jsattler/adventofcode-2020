package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func Day16Part1() {
	input, err := ioutil.ReadFile("../../assets/input-16-sample")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(input))
	fmt.Printf("Solution day 16 part 1: %d\n", 1)
}
