package piscine

type Core struct {
	// ミノ種別のマスターデータ
	Minos []int
	// ミノ形状のマスターデータ
	MinoMap map[string]int
	// 入力をミノ種別に変換したもの
	GivenMinos []int
}

func NewCore() *Core {
	return &Core{
		Minos: []int{
			MinoO,
			MinoLU,
			MinoLR,
			MinoLD,
			MinoLL,
			MinoJU,
			MinoJR,
			MinoJD,
			MinoJL,
			MinoTU,
			MinoTR,
			MinoTD,
			MinoTL,
			MinoIU,
			MinoIR,
			MinoSU,
			MinoSR,
			MinoZU,
			MinoZR,
		},
		MinoMap: MakeMinoShapes(),
	}
}
