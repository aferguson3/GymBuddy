package main

import (
	"fmt"
)

type PurchasedPlates struct {

	lbs map[string]byte
	kgs map[string]byte
}

func (h *HomeGym) initialize() {
	h.Plates.lbs = map[string]byte{
		"2.5":0,
		"5": 0,
		"10": 0,
		"25": 0,
		"35": 0,
		"45": 0,
	}
	h.Plates.kgs = map[string]byte{
		"1.25": 0,
		"2.5":  0,
		"5":    0,
		"10":   0,
		"15":   0,
		"20":   0,
		"25":   0,
	}
}
type HomeGym struct {
	
	FreedomUnits bool 
	Plates PurchasedPlates
	MaxBarWeight int 
	BarWeight int 
	//weight combos: hasmap of all possible weight combinations, ranked by # of weights

}


func (h *HomeGym) BuyPlate(plate float32) {
	// evaluate if plate is of standard size

	if h.FreedomUnits{
		var ownedPlates = &h.Plates.lbs 
	} else {
		var ownedPlates = &h.Plates.kgs
	}

	if ownedPlates[plate] == 0 {

}

func main() {


	// create a home gym
	homeGym := HomeGym{
		FreedomUnits: true,
		Plates: PurchasedPlates{},
		MaxBarWeight: 45,
		BarWeight: 45,
	}
	homeGym.initialize()

	fmt.Println(homeGym.Plates.lbs)
}