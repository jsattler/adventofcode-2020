package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// timestamp - number of minutes since reference point
// first line is earliest timestamp you can depart on a bus
// second line is the bus ids in service
// entries with x are out of service --> ignore them
// figure out earliest bus you can take to the airport
func Day13Part1(){
	input, err := ioutil.ReadFile("assets/input-13-sample")
	if err != nil {
		log.Fatal(err)
	}
	inputStr := strings.Split(string(input), "\n")
	fmt.Println(inputStr)

}
