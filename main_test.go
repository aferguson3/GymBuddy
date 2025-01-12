package main

import (
	"maps"
	"strconv"
	"testing"
)
func setup() (HomeGym){
	var homeGym = HomeGym{
		FreedomUnits: true,
		Plates: PurchasedPlates{make(map[string]byte), make(map[string]byte)},
	}
	homeGym.init()
	StandardWeights_.init()
	return homeGym
}
func TestBuyPlateValidWeights(t *testing.T) {
	
	var homeGym = setup()
	var lbsPlates = maps.Keys(StandardWeights_.lbs)
	var kgsPlates = maps.Keys(StandardWeights_.kgs)
	
	for plate:= range lbsPlates {
		homeGym.BuyPlates(plate, 7)
		got := homeGym.Plates.lbs[plate]
		want:= byte(7)

		if got != want {
			t.Errorf("Plate: %v got %d, want %d",plate, got, want)
		}

	}

	homeGym.FreedomUnits = false
	for plate:= range kgsPlates {
		homeGym.BuyPlates(plate, 12)
		got := homeGym.Plates.kgs[plate]
		want:= byte(12)

		if got != want {
			t.Errorf("Plate: %v got %d, want %d",plate, got, want)
		}

	}
}


func TestBuyPlateInvalidWeights(t *testing.T) {
	var homeGym = setup()
	var plates_1 =[]string {"3", "1000", "33", "7", "232","0", "-1"}

	for i:=0; i < len(plates_1); i++ {
		plate := plates_1[i]
		homeGym.FreedomUnits = true
		homeGym.BuyPlates(plate, 1)
		got := homeGym.Plates.lbs[plate]
		want := byte(0)
	
		if got != want {
			t.Errorf("Plate: %v got %d, want %d",plate, got, want)
		}
	}

	for i:=0; i < len(plates_1); i++ {
		plate := plates_1[i]
		homeGym.FreedomUnits = false
		homeGym.BuyPlates(plate, 1)
		got := homeGym.Plates.kgs[plate]
		want := byte(0)
		
		if got != want {
			t.Errorf("Plate: %v got %d, want %d",plate, got, want)
		}
	}
}

func TestSellPlateValidWeights(t *testing.T) {
	
	var homeGym = setup()
	var lbsPlates = maps.Keys(StandardWeights_.lbs)
	var kgsPlates = maps.Keys(StandardWeights_.kgs)
	var numPlates = byte(6)

	for key := range maps.Keys(homeGym.Plates.lbs){
		homeGym.Plates.lbs[key] = numPlates

	}
	for key := range maps.Keys(homeGym.Plates.kgs){
		homeGym.Plates.kgs[key] = numPlates
	}

	for plate:= range lbsPlates {
		err := homeGym.SellPlates(plate,1)
		got := homeGym.Plates.lbs[plate]
		want:= numPlates-1
		
		if err != nil {
			logger.Println(err)
		}
		if got != want {
			t.Errorf("Plate: %v got %d, want %d",plate, got, want)
		}
		ReportErrors(true)

	}

	homeGym.FreedomUnits = false
	for plate:= range kgsPlates {
		homeGym.SellPlates(plate,1)
		got := homeGym.Plates.kgs[plate]
		want:= numPlates-1

		if got != want {
			t.Errorf("Plate: %v got %d, want %d",plate, got, want)
		}

	}
}

func TestSellPlateInvalidWeights(t *testing.T) {
	var homeGym = setup()
	var plates_1 =[]string {"3", "1000", "33", "7", "232","0", "-1"}

	for i:=0; i < len(plates_1); i++ {
		plate := plates_1[i]
		homeGym.FreedomUnits = true
		homeGym.SellPlates(plate,1)
		got := homeGym.Plates.lbs[plate]
		want := byte(0)
		
		if got != want {
			t.Errorf("Plate: %v got %d, want %d",plate, got, want)
		}
	}

	for i:=0; i < len(plates_1); i++ {
		plate := plates_1[i]
		homeGym.FreedomUnits = false
		homeGym.SellPlates(plate,1)
		got := homeGym.Plates.kgs[plate]
		want := byte(0)
		
		if got != want {
			t.Errorf("Plate: %v got %d, want %d",plate, got, want)
		}
	}
}

func TestSellPlateInvalidAmount(t *testing.T) {
	
	var homeGym = setup()
	var lbsPlates = maps.Keys(StandardWeights_.lbs)
	var kgsPlates = maps.Keys(StandardWeights_.kgs)
	var numPlates = byte(6)

	for key := range maps.Keys(homeGym.Plates.lbs){
		homeGym.Plates.lbs[key] = numPlates

	}
	for key := range maps.Keys(homeGym.Plates.kgs){
		homeGym.Plates.kgs[key] = numPlates
	}

	homeGym.FreedomUnits = true
	for plate:= range lbsPlates {
		err := homeGym.SellPlates(plate,numPlates + 5)
		got := err != nil
		want:= true
		
		if err != nil {
			logger.Println(err)
		}
		if got != want {
			t.Errorf("Plate: %v got %v, want %v",plate, got, want)
		}

	}

	homeGym.FreedomUnits = false
	for plate:= range kgsPlates {
		err :=homeGym.SellPlates(plate,numPlates+ 8)
		got := err != nil
		want:= true

		if err != nil {
			logger.Println(err)
		}
		if got != want {
			t.Errorf("Plate: %v got %v, want %v",plate, got, want)
		}

	}

}

func getMaxWeight(plates map[string]byte)  (totalWeight float32)  {
	for plate, amount := range maps.All(plates) {
		var currPlate, _ = strconv.ParseFloat(plate, 32)
		totalWeight = totalWeight + float32(currPlate) * float32(amount)
	}
	return 
}
func TestCalculateMaxWeight(t *testing.T) {
	var homeGym = setup()
	homeGym.FreedomUnits = true
	newPlates := map[string]byte {"5":5,"45":3,"10":5,"2.5":2,"25":10}

	for plate, amount := range maps.All(newPlates) {
		homeGym.BuyPlates(plate, amount)
	}
	
	expectedMaxWeight := getMaxWeight(newPlates)
	got:= homeGym.MaxWeight
	want := expectedMaxWeight

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}