package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"math"
	"strings"
)


type Coordinate4D struct {
	x, y, z, w int
}

type HyperCubes map[Coordinate4D]bool

type MinMax4D struct {
	minX, maxX int
	minY, maxY int
	minZ, maxZ int
	minW, maxW int
}

// same as part 1 just add one dimension
func Day17Part2() {
	input := aocio.ReadFileSplit("../../assets/input-17", "\n")
	//input = input[:len(input)-1]
	active := parseHyperCubes(input)
	for i := 0; i < 6; i++ {
		active = active.simulateHyperCube()
	}
	fmt.Printf("Solution day 17 part 2: %d\n", len(active))
}

func (cs HyperCubes) simulateHyperCube() (nextActive HyperCubes) {
	minMax := minMaxHyperCube(cs)
	nextActive = make(HyperCubes)
	for x := minMax.minX - 1; x <= minMax.maxX+1; x++ {
		for y := minMax.minY - 1; y <= minMax.maxY+1; y++ {
			for z := minMax.minZ - 1; z <= minMax.maxZ+1; z++ {
				for w := minMax.minW - 1; w <= minMax.maxW+1; w++{
					coord := Coordinate4D{x, y, z, w}
					neighbors := cs.neighborCountHyperCube(coord)
					if cs[coord] && (neighbors == 2 || neighbors == 3) { // check if coordinate is active and 2 or three neighbors are active
						nextActive[coord] = true
					}
					if !cs[coord] && neighbors == 3 { // check if coordinate is inactive and 3 neighbors are active
						nextActive[coord] = true
					}
				}
			}
		}
	}
	return nextActive
}

func (cs HyperCubes) neighborCountHyperCube(coord Coordinate4D) (nCount int) {
	neighborRange := []int{-1, 0, 1}
	for _, x := range neighborRange {
		for _, y := range neighborRange {
			for _, z := range neighborRange {
				for _, w := range neighborRange {
					if x == 0 && y == 0 && z == 0 && w == 0{
						continue // will not count ourselves as a neighbor
					}
					neighbor := Coordinate4D{coord.x + x, coord.y + y, coord.z + z, coord.w + w}
					if _, ok := cs[neighbor]; ok {
						nCount++
					}
				}
			}
		}
	}
	return nCount
}

// get the min/max coordination bounds of the active cubes
func minMaxHyperCube(cs HyperCubes) (res MinMax4D) {
	for c := range cs {
		res.maxX = int(math.Max(float64(c.x), float64(res.maxX)))
		res.minX = int(math.Min(float64(c.x), float64(res.minX)))
		res.maxY = int(math.Max(float64(c.y), float64(res.maxY)))
		res.minY = int(math.Min(float64(c.y), float64(res.minY)))
		res.maxZ = int(math.Max(float64(c.z), float64(res.maxZ)))
		res.minZ = int(math.Min(float64(c.z), float64(res.minZ)))
		res.maxW = int(math.Max(float64(c.w), float64(res.maxW)))
		res.minW = int(math.Min(float64(c.w), float64(res.minW)))
	}
	return res
}

func parseHyperCubes(lines []string) (result HyperCubes) {
	result = make(HyperCubes)
	for x, line := range lines {
		for y, cube := range line {
			if cube == active {
				c := Coordinate4D{x, y, 0, 0}
				result[c] = true
			}
		}
	}
	return result
}

func (cs HyperCubes) print(){
	fmt.Printf("%s\n", strings.Repeat("-", 32))
	fmt.Println("active cubes:", len(cs))
	for c, _ := range cs {
		fmt.Printf("x:%d, y:%d, z:%d, w:%d\n", c.x, c.y, c.z, c.w)
	}
	fmt.Printf("%s\n", strings.Repeat("-", 32))
}