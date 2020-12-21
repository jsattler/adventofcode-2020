package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"math"
	"strings"
)

const (
	active = '#'
)

type Coordinate struct {
	x, y, z int
}

type Cubes map[Coordinate]bool

type MinMax3D struct {
	minX, maxX int
	minY, maxY int
	minZ, maxZ int
}

// 1. if cube is active and exactly 2 or 3 of its neighbors are also active the cube remains active
// 2. if cube is inactive but exactly 3 of its neighbors are are active the cube becomes active
// Idea: initialize with origin and expand everytime a new active cube is initialized outside current "shell"
// iterate over each cube and check if the rules from above apply
func Day17Part1() {
	input := aocio.ReadFileSplit("../../assets/input-17", "\n")
	//input = input[:len(input)-1]
	active := parseGrid(input)
	for i := 0; i < 6; i++ {
		active = active.simulate()
	}
	fmt.Printf("Solution day 17 part 1: %d\n", len(active))
}

func (cs Cubes) simulate() (nextActive Cubes) {
	minMax := minMax(cs)
	nextActive = make(Cubes)
	for x := minMax.minX - 1; x <= minMax.maxX+1; x++ {
		for y := minMax.minY - 1; y <= minMax.maxY+1; y++ {
			for z := minMax.minZ - 1; z <= minMax.maxZ+1; z++ {
				coord := Coordinate{x, y, z}
				neighbors := cs.neighborCount(coord)
				if cs[coord] && (neighbors == 2 || neighbors == 3) { // check if coordinate is active and 2 or three neighbors are active
					nextActive[coord] = true
				}
				if !cs[coord] && neighbors == 3 { // check if coordinate is inactive and 3 neighbors are active
					nextActive[coord] = true
				}
			}
		}
	}
	return nextActive
}

func (cs Cubes) neighborCount(coord Coordinate) (nCount int) {
	neighborRange := []int{-1, 0, 1}
	for _, x := range neighborRange {
		for _, y := range neighborRange {
			for _, z := range neighborRange {
				if x == 0 && y == 0 && z == 0 {
					continue // will not count ourselves as a neighbor
				}
				neighbor := Coordinate{coord.x + x, coord.y + y, coord.z + z}
				if _, ok := cs[neighbor]; ok {
					nCount++
				}
			}
		}
	}
	return nCount
}

// get the min/max coordination bounds of the active cubes
func minMax(cs Cubes) (res MinMax3D) {
	for c := range cs {
		res.maxX = int(math.Max(float64(c.x), float64(res.maxX)))
		res.minX = int(math.Min(float64(c.x), float64(res.minX)))
		res.maxY = int(math.Max(float64(c.y), float64(res.maxY)))
		res.minY = int(math.Min(float64(c.y), float64(res.minY)))
		res.maxZ = int(math.Max(float64(c.z), float64(res.maxZ)))
		res.minZ = int(math.Min(float64(c.z), float64(res.minZ)))
	}
	return res
}

func parseGrid(lines []string) (result Cubes) {
	result = make(Cubes)
	for x, line := range lines {
		for y, cube := range line {
			if cube == active {
				c := Coordinate{x, y, 0}
				result[c] = true
			}
		}
	}
	return result
}

func (cs Cubes) print(){
	fmt.Printf("%s\n", strings.Repeat("-", 32))
	fmt.Println("active cubes:", len(cs))
	for c, _ := range cs {
		fmt.Printf("x:%d, y:%d, z:%d\n", c.x, c.y, c.z)
	}
	fmt.Printf("%s\n", strings.Repeat("-", 32))
}