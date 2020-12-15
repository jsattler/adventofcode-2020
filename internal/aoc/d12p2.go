package aoc

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// relative to the ship
type Waypoint struct {
	southNorth int
	westEast   int
}

/*
Action N means to move the waypoint north by the given value.
Action S means to move the waypoint south by the given value.
Action E means to move the waypoint east by the given value.
Action W means to move the waypoint west by the given value.
Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
Action F means to move forward to the waypoint a number of times equal to the given value.
*/
func Day12Part2() {

	input, err := ioutil.ReadFile("assets/input-12")
	if err != nil {
		log.Fatal(err)
	}
	inputStr := strings.Split(string(input), "\n")
	actions := parseActions(inputStr)

	start := Waypoint{1, 10}
	ship := Ship{waypoint: start}

	for _, a := range actions {
		ship.navigateWaypoint(a)
	}

	fmt.Printf("Solution day 12 part 2: %d\n", intAbs(ship.westEast)+intAbs(ship.southNorth))
}

func (s *Ship) navigateWaypoint(a Action) {
	switch a.action {
	case moveNorth: // move waypoint north
		s.waypoint.southNorth += a.value
	case moveEast: // move waypoint east
		s.waypoint.westEast += a.value
	case moveSouth: // move waypoint south
		s.waypoint.southNorth -= a.value
	case moveWest: // move waypoint west
		s.waypoint.westEast -= a.value
	case turnLeft: // move waypoint counter-clockwise by value around ship
		s.rotate(-a.value)
	case turnRight: // move waypoint clockwise by value around ship
		s.rotate(a.value)
	case moveForward: // move ship into the direction of the waypoint
		s.forward(a.value)
	}
}

// rotate
func (s *Ship) rotate(degree int) {
	degree = mod(degree, maxDegrees) // make sure if e.g. -270 becomes 90

	x := s.waypoint.westEast
	y := s.waypoint.southNorth

	// we only have to distinguish between three cases
	switch degree {

	case 90: // (x,y) --> (y,-x)
		s.waypoint.westEast = y
		s.waypoint.southNorth = -x
	case 180: // (x,y) --> (-x,-y)
		s.waypoint.westEast = -x
		s.waypoint.southNorth = -y
	case 270: // (x,y) --> (-y, x)
		s.waypoint.westEast = -y
		s.waypoint.southNorth = x
	}

}

// move ship forward towards waypoint
// waypoint is relative and moves accordingly
func (s *Ship) forward(val int) {
	// move the ship to the waypoint by val times
	s.southNorth += s.waypoint.southNorth * val
	s.westEast += s.waypoint.westEast * val
}
