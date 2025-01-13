package main

type PurchasedPlates struct {

	lbs map[string]byte
	kgs map[string]byte
}
type StandardWeights struct {
	lbs map[string]bool
	kgs map[string]bool

}

func (s *StandardWeights) init() {
	s.lbs = map[string]bool{
		"2.5":true,
		"5": true,
		"10": true,
		"25": true,
		"35": true,
		"45": true,
	}
	s.kgs = map[string]bool{
		"1.25": true,
		"2.5":  true,
		"5":    true,
		"10":   true,
		"15":   true,
		"20":   true,
		"25":   true,
	}
}


type HomeGym struct {
	
	FreedomUnits bool 
	Plates PurchasedPlates
	MaxWeight float32 
	BarWeight byte 
	//weight combos: hasmap of all possible weight combinations, ranked by # of weights

}