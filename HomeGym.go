package main
import (
	"maps"
	"fmt"
	"strconv"
)

func (h *HomeGym) init() {
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
func getPlates(h* HomeGym) (ownedPlates *map[string]byte, standardPlates *map[string]bool){
	if h.FreedomUnits{
		ownedPlates = &h.Plates.lbs 
		standardPlates = &StandardWeights_.lbs
	} else {
		ownedPlates = &h.Plates.kgs
		standardPlates = &StandardWeights_.kgs
	}
	return
}

func (h *HomeGym) BuyPlates(plate string, amount byte) {
	var ownedPlates, standardPlates = getPlates(h)
	
	if _, ok:= (*standardPlates)[plate]; ok{
		(*ownedPlates)[plate] =	(*ownedPlates)[plate] + amount
	}else{
		logger.Printf("%v: is a non-standard weight plate", plate)
	}

	h.CalculateMaxWeight()
}


func (h *HomeGym) SellPlates(plate string, amount byte) (err error){
	ownedPlates, standardPlates := getPlates(h)

	if _, ok:=(*standardPlates)[plate]; ok{
		result := int((*ownedPlates)[plate]) - int(amount); if result < 0 {
			err = fmt.Errorf("HomeGym.FreedomUnits:%v %v:%v has < %v plates", h.FreedomUnits, plate, (*ownedPlates)[plate], amount)
			return 
		}

		if (*ownedPlates)[plate] > 0 {
			(*ownedPlates)[plate] = (*ownedPlates)[plate] - amount

		}else {
			logger.Printf("No %v plates to remove", plate)
		}

	}else{
		logger.Printf("%v: is a non-standard weight plate", plate)
	}

	h.CalculateMaxWeight()
	return
}
func (h * HomeGym) CalculateMaxWeight() (err error){
	ownedPlates, _ := getPlates(h)
	var totalWeight float32

	for plate, count := range maps.All(*ownedPlates) {
		var plate_int, err = strconv.ParseFloat(plate, 32); if err != nil {
			logger.Printf("Error converting %v to int", plate)
			return err
		}
		totalWeight = totalWeight +  float32(plate_int) * float32(count)
	}
	h.MaxWeight = totalWeight
	return
}

func(h* HomeGym) WeightCombos() {
	/*
	Known: purchased plates
	Want: weight(s) to achieve, minimal plates used, even on both sides
	Weight : [bar weight] + [weight plates used]
	225: 45(bar) 45 45 35 35 10 10
	*/
	var ownedPlates map[string]byte
	if h.FreedomUnits {
		ownedPlates = h.Plates.lbs
	}else{
		ownedPlates = h.Plates.kgs
	}

	
}

func(h* HomeGym) GetCombo() (err error){
	//try to retrive combo of resired weight
	//if solution possible -->
	//else get next lowest and highest solution
}