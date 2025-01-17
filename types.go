package main

type PlateCount = map[float32]byte

type PurchasedPlates struct {
	lbs PlateCount
	kgs PlateCount
}
type StandardWeights struct {
	lbs map[float32]bool
	kgs map[float32]bool
}

func (s *StandardWeights) init() {
	s.lbs = map[float32]bool{
		2.5: true,
		5:   true,
		10:  true,
		25:  true,
		35:  true,
		45:  true,
	}
	s.kgs = map[float32]bool{
		1.25: true,
		2.5:  true,
		5:    true,
		10:   true,
		15:   true,
		20:   true,
		25:   true,
	}
}

type HomeGym struct {
	FreedomUnits bool
	PlateCounter PurchasedPlates
	MaxWeight    float32
	BarWeight    float32
	//weight combos: hasmap of all possible weight combinations, ranked by # of weights

}

type TestHomeGym = HomeGym

func (h *TestHomeGym) init() {
	h.PlateCounter.lbs = map[float32]byte{
		2.5: 0,
		5:   0,
		10:  0,
		25:  0,
		35:  0,
		45:  0,
	}
	h.PlateCounter.kgs = map[float32]byte{
		1.25: 0,
		2.5:  0,
		5:    0,
		10:   0,
		15:   0,
		20:   0,
		25:   0,
	}
}
