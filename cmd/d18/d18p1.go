package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"strconv"
	"strings"
)

var opa = map[string]struct {
	precedence int
	rAssoc     bool
}{
	"*": {0, false}, // MUL and PLUS have same precedence and are not right associative
	"+": {0, false}, // MUL and PLUS have same precedence and are not right associative
}

// Using the shunting yard algorithm with reverse polish notation to evaluate expressions
func Day18Part1() {
	expressions := aocio.ReadFileSplit("../../assets/input-18", "\n")
	sum := float64(0)
	for _, expr := range expressions {
		expr = strings.ReplaceAll(expr, "(", "( ")
		expr = strings.ReplaceAll(expr, ")", " )")
		rpn := shuntingYard(expr)
		sum += evaluateRPN(rpn)
	}

	fmt.Printf("Solution day 18 part 1: %d\n", int(sum))
}


// https://rosettacode.org/wiki/Parsing/Shunting-yard_algorithm
// with adaptions to follow precedence rules
func shuntingYard(e string) (rpn string) {
	var stack []string
	for _, tok := range strings.Fields(e) {
		switch tok {
		case "(":
			stack = append(stack, tok) // push "(" to stack
		case ")":
			var op string
			for {
				// pop item ("(" or operator) from stack
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // discard "("
				}
				rpn += " " + op // add operator to result
			}
		default:
			if o1, isOp := opa[tok]; isOp {
				for len(stack) > 0 {
					op := stack[len(stack)-1]
					if o2, isOp := opa[op]; !isOp || o1.precedence > o2.precedence ||
						o1.precedence == o2.precedence && o1.rAssoc {
						break
					}
					stack = stack[:len(stack)-1]
					rpn += " " + op
				}
				stack = append(stack, tok)
			} else {
				if rpn > "" {
					rpn += " "
				}
				rpn += tok
			}
		}
	}
	for len(stack) > 0 {
		rpn += " " + stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return
}

// https://rosettacode.org/wiki/Parsing/RPN_calculator_algorithm
// with adaptions to follow precedence rules
func evaluateRPN(input string) (res float64){
	var stack []float64
	for _, tok := range strings.Fields(input) {
		switch tok {
		case "+":
			stack[len(stack)-2] += stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case "*":
			stack[len(stack)-2] *= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default:
			f, _ := strconv.ParseFloat(tok, 64)
			stack = append(stack, f)
		}
	}
	return stack[0]
}