package main

import (
	"maps"
	"slices"
	"testing"
)

func testSetup() TestHomeGym {
	var homeGym = HomeGym{
		FreedomUnits:   true,
		PlateInventory: PlateInventory{},
	}
	homeGym.init()
	standardWeights.init()
	return homeGym
}
func (h *TestHomeGym) testSetup() {
	*h = HomeGym{
		FreedomUnits: true,
		BarWeight:    45,
	}
	h.init()
	standardWeights.init()
}
func TestBuyPlateValidWeights(t *testing.T) {

	var homeGym = testSetup()
	var lbsPlates = maps.Keys(standardWeights.lbs)
	var kgsPlates = maps.Keys(standardWeights.kgs)

	for plate := range lbsPlates {
		homeGym.BuyPlates(plate, 7)
		got := homeGym.PlateInventory.lbs[plate]
		want := byte(7)

		if got != want {
			t.Errorf("Plate: %v got %d, want %d", plate, got, want)
		}

	}

	homeGym.FreedomUnits = false
	for plate := range kgsPlates {
		homeGym.BuyPlates(plate, 12)
		got := homeGym.PlateInventory.kgs[plate]
		want := byte(12)

		if got != want {
			t.Errorf("Plate: %v got %d, want %d", plate, got, want)
		}

	}
}

func TestBuyPlateInvalidWeights(t *testing.T) {
	var homeGym = testSetup()
	var plates_1 = []float32{3, 1000, 33, 7, 232, 0, -1}

	for i := 0; i < len(plates_1); i++ {
		plate := plates_1[i]
		homeGym.FreedomUnits = true
		homeGym.BuyPlates(plate, 1)
		got := homeGym.PlateInventory.lbs[plate]
		want := byte(0)

		if got != want {
			t.Errorf("Plate: %v got %d, want %d", plate, got, want)
		}
	}

	for i := 0; i < len(plates_1); i++ {
		plate := plates_1[i]
		homeGym.FreedomUnits = false
		homeGym.BuyPlates(plate, 1)
		got := homeGym.PlateInventory.kgs[plate]
		want := byte(0)

		if got != want {
			t.Errorf("Plate: %v got %d, want %d", plate, got, want)
		}
	}
}

func TestSellPlateValidWeights(t *testing.T) {

	var homeGym = testSetup()
	var lbsPlates = maps.Keys(standardWeights.lbs)
	var kgsPlates = maps.Keys(standardWeights.kgs)
	var numPlates = byte(6)

	for key := range maps.Keys(homeGym.PlateInventory.lbs) {
		homeGym.PlateInventory.lbs[key] = numPlates

	}
	for key := range maps.Keys(homeGym.PlateInventory.kgs) {
		homeGym.PlateInventory.kgs[key] = numPlates
	}

	for plate := range lbsPlates {
		err := homeGym.SellPlates(plate, 1)
		got := homeGym.PlateInventory.lbs[plate]
		want := numPlates - 1

		if err != nil {
			logger.Println(err)
		}
		if got != want {
			t.Errorf("Plate: %v got %d, want %d", plate, got, want)
		}
		ReportErrors(true)

	}

	homeGym.FreedomUnits = false
	for plate := range kgsPlates {
		homeGym.SellPlates(plate, 1)
		got := homeGym.PlateInventory.kgs[plate]
		want := numPlates - 1

		if got != want {
			t.Errorf("Plate: %v got %d, want %d", plate, got, want)
		}

	}
}

func TestSellPlateInvalidWeights(t *testing.T) {
	var homeGym = testSetup()
	var plates_1 = []float32{3, 1000, 33, 7, 232, 0, -1}

	for i := 0; i < len(plates_1); i++ {
		plate := plates_1[i]
		homeGym.FreedomUnits = true
		homeGym.SellPlates(plate, 1)
		got := homeGym.PlateInventory.lbs[plate]
		want := byte(0)

		if got != want {
			t.Errorf("Plate: %v got %d, want %d", plate, got, want)
		}
	}

	for i := 0; i < len(plates_1); i++ {
		plate := plates_1[i]
		homeGym.FreedomUnits = false
		homeGym.SellPlates(plate, 1)
		got := homeGym.PlateInventory.kgs[plate]
		want := byte(0)

		if got != want {
			t.Errorf("Plate: %v got %d, want %d", plate, got, want)
		}
	}
}

func TestSellPlateInvalidAmount(t *testing.T) {

	var homeGym = testSetup()
	var lbsPlates = maps.Keys(standardWeights.lbs)
	var kgsPlates = maps.Keys(standardWeights.kgs)
	var numPlates = byte(6)

	for key := range maps.Keys(homeGym.PlateInventory.lbs) {
		homeGym.PlateInventory.lbs[key] = numPlates

	}
	for key := range maps.Keys(homeGym.PlateInventory.kgs) {
		homeGym.PlateInventory.kgs[key] = numPlates
	}

	homeGym.FreedomUnits = true
	for plate := range lbsPlates {
		err := homeGym.SellPlates(plate, numPlates+5)
		got := err != nil
		want := true

		if err != nil {
			logger.Println(err)
		}
		if got != want {
			t.Errorf("Plate: %v got %v, want %v", plate, got, want)
		}

	}

	homeGym.FreedomUnits = false
	for plate := range kgsPlates {
		err := homeGym.SellPlates(plate, numPlates+8)
		got := err != nil
		want := true

		if err == nil {
			logger.Println("No error, but was expecting one")
		}
		if got != want {
			t.Errorf("Plate: %v got %v, want %v", plate, got, want)
		}

	}

}

func getMaxWeight(plates map[float32]byte) (totalWeight float32) {
	for plate, amount := range maps.All(plates) {
		totalWeight = totalWeight + plate*float32(amount)
	}
	return
}
func TestCalculateMaxWeight(t *testing.T) {
	var homeGym = testSetup()
	homeGym.BarWeight = 45
	homeGym.FreedomUnits = true
	newPlates := map[float32]byte{5: 5, 45: 3, 10: 5, 2.5: 2, 25: 10}

	for plate, amount := range maps.All(newPlates) {
		homeGym.BuyPlates(plate, amount)
	}

	expectedMaxWeight := getMaxWeight(newPlates) + homeGym.BarWeight
	got := homeGym.MaxWeight
	want := expectedMaxWeight

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestCombos(t *testing.T) {
	var homeGym TestHomeGym
	numPlates := byte(10)
	homeGym.testSetup()
	plates_lbs := map[float32]byte{45: numPlates, 35: numPlates, 25: numPlates, 10: numPlates, 5: numPlates, 2.5: numPlates}
	plates_kgs := map[float32]byte{45: numPlates, 35: numPlates, 25: numPlates, 10: numPlates, 5: numPlates, 2.5: numPlates}

	homeGym = TestHomeGym{
		BarWeight:      45,
		FreedomUnits:   true,
		PlateInventory: PlateInventory{(plates_lbs), (plates_kgs)},
	}

	want := [][]float32{
		{45, 45, 45, 45},
		{45, 45, 25, 25},
		{45, 45},
		{},
		{},
		{},
	}
	expected_errors := []byte{3, 4}
	var (
		combo []float32
		err   error
	)

	for i := 0; i < len(want); i++ {
		switch i {
		case 0:
			combo, err = homeGym.GetCombo(homeGym.PlateInventory.lbs, 225)
		case 1:
			combo, err = homeGym.GetCombo(homeGym.PlateInventory.lbs, 185)
		case 2:
			combo, err = homeGym.GetCombo(homeGym.PlateInventory.lbs, 135)
		case 3:
			combo, err = homeGym.GetCombo(homeGym.PlateInventory.lbs, 0)
		case 4:
			combo, err = homeGym.GetCombo(homeGym.PlateInventory.lbs, -100)
		case 5:
			combo, err = homeGym.GetCombo(homeGym.PlateInventory.lbs, homeGym.BarWeight)
		}
		if slices.Contains(expected_errors, byte(i)) {
			t.Logf("expected error: %v", err)
		}

		if err != nil && !slices.Contains(expected_errors, byte(i)) {
			t.Errorf("error: %v", err)
		}

		if !slices.Equal(combo, want[i]) {
			t.Errorf("got: %v, want: %v", combo, want[i])
		}

	}

}
