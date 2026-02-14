package cube

type Side struct {
	Index  int
	Colors []string
}
type CubeModel struct {
	Sides []Side
}

var moves = map[string]func(int, *CubeModel){
	"R": func(r int, cube *CubeModel) {},
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
	colors := []string{"WHITE", "ORANGE", "GREEN", "RED", "BLUE", "YELLOW"}
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
