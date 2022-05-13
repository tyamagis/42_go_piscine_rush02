package piscine

// ミノタイプ
const MinoTypes = 19
const (
	MinoO = iota
	MinoLU
	MinoLR
	MinoLD
	MinoLL
	MinoJU
	MinoJR
	MinoJD
	MinoJL
	MinoTU
	MinoTR
	MinoTD
	MinoTL
	MinoIU
	MinoIR
	MinoSU
	MinoSR
	MinoZU
	MinoZR
)

const MinoSize = 4

func MakeMinoShapes() map[string]int {
	shapes := make([][]string, 0)
	rotator := func(shape []string, n int) {
		s := shape
		for k := 0; k < n; k++ {
			shapes = append(shapes, s)
			s = RotateLines(s)
		}
	}

	shapeO := []string{
		"##",
		"##",
	}
	shapeL := []string{
		"#.",
		"#.",
		"##",
	}
	// J = mirror of L
	shapeJ := ReverseLines(shapeL)
	shapeT := []string{
		"###",
		".#.",
	}
	shapeS := []string{
		".##",
		"##.",
	}
	// Z = mirror of S
	shapeZ := ReverseLines(shapeS)
	shapeI := []string{
		"#",
		"#",
		"#",
		"#",
	}

	rotator(shapeO, 1)
	rotator(shapeL, 4)
	rotator(shapeJ, 4)
	rotator(shapeT, 4)
	rotator(shapeI, 2)
	rotator(shapeS, 2)
	rotator(shapeZ, 2)

	shapeMap := map[string]int{}
	for i, shape := range shapes {
		shapeMap[Join(shape, "\n")] = i
	}
	return shapeMap
}
