package cube

import "fmt"

type Side struct {
	Index  int
	Colors []string
}

func (side *Side) PrintSide(cubeSize int) {
	for i := range cubeSize {
		for j := (i * 3); j < (i+1)*cubeSize; j++ {
			fmt.Printf("%s ", side.Colors[j])
		}
		fmt.Print("\n")
	}
}

type CubeModel struct {
	Sides []Side
}

var moves = map[string]func(int, *CubeModel){
	"R": func(r int, cube *CubeModel) {
		sidesOrder := []int8{0, 1, 5, 3, 0}
		for i := range 4 {
			cube.Sides[sidesOrder[i]].Colors[2], cube.Sides[sidesOrder[i+1]].Colors[2] = cube.Sides[sidesOrder[i+1]].Colors[2], cube.Sides[sidesOrder[i]].Colors[2]
			//add the rest of sides and rotating right side by 90 deg
		}
	},
	"L": func(r int, cube *CubeModel) {},
	"U": func(r int, cube *CubeModel) {},
	"D": func(r int, cube *CubeModel) {},
	"F": func(r int, cube *CubeModel) {},
	"B": func(r int, cube *CubeModel) {},
	"r": func(r int, cube *CubeModel) {},
	"l": func(r int, cube *CubeModel) {},
	"u": func(r int, cube *CubeModel) {},
	"d": func(r int, cube *CubeModel) {},
	"f": func(r int, cube *CubeModel) {},
	"b": func(r int, cube *CubeModel) {},
	"x": func(r int, cube *CubeModel) {},
	"y": func(r int, cube *CubeModel) {},
	"z": func(r int, cube *CubeModel) {},
	"M": func(r int, cube *CubeModel) {},
	"E": func(r int, cube *CubeModel) {},
	"S": func(r int, cube *CubeModel) {},
}

func (cube *CubeModel) ScrambleCube(scramble string) {

}
func (cube *CubeModel) printCube() {

}
func CreateCube(n int) CubeModel {
	var mainCube CubeModel
	colors := []string{"W", "O", "G", "R", "B", "Y"}
	for i, color := range colors {
		var sideColors = []string{}
		for j := 0; j < n*n; j++ {
			sideColors = append(sideColors, color)
		}
		newSide := Side{i, sideColors}
		mainCube.Sides = append(mainCube.Sides, newSide)
	}
	return mainCube
}
