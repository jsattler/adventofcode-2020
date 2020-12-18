package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"log"
	"math/big"
	"strconv"
	"strings"
)

// solved with some help
// initial brute force approach worked for smaller sets
func Day13Part2() {
	inputStr := aocio.ReadFileSplit("../../assets/input-13", "\n")
	_, bids := parseDeparturesWithX(inputStr)
	findTimestamp(bids)
}

func findTimestamp(bids []int) {
	a := make([]*big.Int, 0)
	n := make([]*big.Int, 0)

	for i, bid := range bids {
		if bid == 0 {
			continue
		}
		a = append(a, big.NewInt(int64(bid)))
		n = append(n, big.NewInt(int64(bid-i)))
	}

	res, err := crt(n, a)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Solution day 13 part 2: %s\n", res.String())
}

var one = big.NewInt(1)

// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go (and thanks @ Felix)
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func parseDeparturesWithX(inputStr []string) (int, []int) {
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
		} else {
			bids = append(bids, 0)
		}
	}
	return timestamp, bids
}
