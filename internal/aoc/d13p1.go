package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func Day13Part1(){
	input, err := ioutil.ReadFile("assets/input-13")
	if err != nil {
		log.Fatal(err)
	}
	inputStr := strings.Split(string(input), "\n")
	timestamp, bids := parseDepartures(inputStr)
	earliest := math.MaxInt32
	var takeBid int
	for _, bid := range bids {
		if time := bid - (timestamp % bid); time < earliest {
			earliest = time
			takeBid = bid
		}
	}
	fmt.Printf("Solution day 13 part 1: %d\n", earliest * takeBid)
}

func parseDepartures(inputStr []string) (int, []int){
	timestamp, err := strconv.Atoi(inputStr[0])
	if err != nil {
		log.Fatal(err)
	}
	ids := strings.Split(inputStr[1], ",")
	bids := make([]int, 0)
	for _, id := range ids {
		if id != "x" {
			bid, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			bids = append(bids, bid)
		}
	}
	return timestamp, bids
}