package main

import (
	"mikolajczak/rubiks/pkg/cube"
)

func main() {
	var mainCube cube.CubeModel = cube.CreateCube(9)
	mainCube.Sides[0].PrintSide(3)
}
