package main

import (
	"bytes"
	"fmt"
	"log"
	"maps"
	"strconv"
)

var StandardWeights_ = StandardWeights{}

var (
	buf bytes.Buffer
	logger = log.New(&buf, "",log.Lshortfile | log.Ltime)
)

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
	CurrentWeight float32 
	//weight combos: hasmap of all possible weight combinations, ranked by # of weights

}

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
func (h * HomeGym) CalculateMaxWeight() {
	ownedPlates, _ := getPlates(h)
	var totalWeight float32

	for plate, count := range maps.All(*ownedPlates) {
		plate_int, err := strconv.ParseFloat(plate, 32); if err != nil {
			logger.Printf("Error converting %v to int", plate)
		}
		totalWeight = totalWeight +  float32(plate_int) * float32(count)
	}
	h.MaxWeight = totalWeight
}

func main(){
	StandardWeights_.init()
	homeGym := HomeGym{
		FreedomUnits: true,
		Plates: PurchasedPlates{make(map[string]byte), make(map[string]byte)},
	}

	logger.Println(homeGym)
	ReportErrors(true)
}

func ReportErrors(x bool) {
	if x {
		fmt.Print(&buf)

	}
}