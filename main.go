package main

import "fmt"

type side struct {
	index  int
	colors []string
}
type cube struct {
	sides []side
}

func (cube *cube) scrambleCube() {

}
func (cube *cube) printCube() {

}
func createCube(n int) cube {
	var mainCube cube
	colors := []string{"WHITE", "ORANGE", "GREEN", "RED", "BLUE", "YELLOW"}
	for i, color := range colors {
		var sideColors = []string{}
		for j := 0; j < n*n; j++ {
			sideColors = append(sideColors, color)
		}
		newSide := side{i, sideColors}
		mainCube.sides = append(mainCube.sides, newSide)
	}
	return mainCube
}

func main() {
	var mainCube cube = createCube(9)
	fmt.Printf("%s", mainCube.sides[5].colors[0])
}
