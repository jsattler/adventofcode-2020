package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
)

// relative to the ship
type Waypoint struct {
	southNorth int
	westEast   int
}

func Day12Part2() {
	inputStr := aocio.ReadFileSplit("../../assets/input-12", "\n")
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
