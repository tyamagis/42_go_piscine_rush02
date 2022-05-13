package piscine

type Core struct {
	Minos []int
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
	}
}
