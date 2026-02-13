package cube

type Side struct {
	Index  int
	Colors []string
}
type CubeModel struct {
	Sides []Side
}

func (cube *CubeModel) scrambleCube() {

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
