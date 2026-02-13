package main

import (
	"fmt"
	"mikolajczak/rubiks/pkg/cube"
)

func main() {
	var mainCube cube.CubeModel = cube.CreateCube(9)
	fmt.Printf("%s", mainCube.Sides[5].Colors[0])
}
