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

type MinoMaster struct {
	MinoType      int
	Height, Width int
	shape         [][]rune
}

func MakeMinoShapes() (map[string]*MinoMaster, map[int]*MinoMaster) {
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

	shapeMap := map[string]*MinoMaster{}
	shapeReverseMap := map[int]*MinoMaster{}
	for i, shape := range shapes {
		h := len(shape)
		w := len(shape[0])
		joined := Join(shape, "\n")
		mm := &MinoMaster{
			MinoType: i,
			Height:   h,
			Width:    w,
			shape: Map(shape, func(s string, _ int) []rune {
				return []rune(s)
			}),
		}
		shapeMap[joined] = mm
		shapeReverseMap[i] = mm
	}
	return shapeMap, shapeReverseMap
}
