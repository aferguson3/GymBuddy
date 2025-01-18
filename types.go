package main

import "slices"

type PlateSet = map[float32]byte

func getKeys[V comparable](p map[float32]V) (keys []float32) {
	keys = make([]float32, 0, len(p))
	for key := range p {
		keys = append(keys, key)
	}

	return keys
}

func getValues[V any](p map[float32]V, keys []float32) (values []V) {
	values = make([]V, 0, len(p))
	for i := 0; i < len(keys); i++ {
		value := p[keys[i]]
		values = append(values, value)
	}
	return values
}
func sortByKeys[V any](p map[float32]V) (keys []float32, values []V) {
	keys, values = make([]float32, 0, len(p)), make([]V, 0, len(p))
	slices.Sort(keys) ;slices.Reverse(keys)
	
}

type PlateInventory struct {
	lbs PlateSet
	kgs PlateSet
}

type StandardPlateWeights struct {
	lbs map[float32]bool
	kgs map[float32]bool
}

func (s *StandardPlateWeights) init() {
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
	FreedomUnits   bool
	PlateInventory PlateInventory
	MaxWeight      float32
	BarWeight      float32
	//weight combos: hasmap of all possible weight combinations, ranked by # of weights

}

type TestHomeGym = HomeGym

func (h *TestHomeGym) init() {
	h.PlateInventory.lbs = map[float32]byte{
		2.5: 0,
		5:   0,
		10:  0,
		25:  0,
		35:  0,
		45:  0,
	}
	h.PlateInventory.kgs = map[float32]byte{
		1.25: 0,
		2.5:  0,
		5:    0,
		10:   0,
		15:   0,
		20:   0,
		25:   0,
	}
}
