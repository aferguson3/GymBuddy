package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
)


func ToFloatKeys[V any](oldMap map[string]V) (newMap map[float32]V, err error) {
	newMap = make(map[float32]V)

	for key, value := range maps.All(oldMap) {
		newKey, err := strconv.ParseFloat(key, 32)
		if err != nil {
			err = fmt.Errorf("error converting %v to float", key)
			return nil, err
		}
		newMap[float32(newKey)] = value
	}
	return newMap, err
}

func ToStringKeys[V any](oldMap map[float32]V) (newMap map[string]V) {
	newMap = make(map[string]V)
	for key, value := range maps.All(oldMap) {
		newMap[strconv.FormatFloat(float64(key), 'g', -1, 32)] = value
	}
	return newMap

}
func getPlates(h *HomeGym) (ownedPlates *PlateCount, standardPlates *map[float32]bool) {
	if h.FreedomUnits {
		ownedPlates = &h.PlateCounter.lbs
		standardPlates = &standardWeights.lbs
	} else {
		ownedPlates = &h.PlateCounter.kgs
		standardPlates = &standardWeights.kgs
	}
	return
}

func (h *HomeGym) BuyPlates(plate float32, amount byte) (err error) {
	var ownedPlates, standardPlates = getPlates(h)

	if _, ok := (*standardPlates)[plate]; ok {
		(*ownedPlates)[plate] = (*ownedPlates)[plate] + amount
	} else {
		err = fmt.Errorf("%v: is a non-standard weight plate", plate)
	}

	h.CalculateMaxWeight()
	return
}

func (h *HomeGym) SellPlates(plate float32, amount byte) (err error) {
	ownedPlates, standardPlates := getPlates(h)

	if _, ok := (*standardPlates)[plate]; ok {
		result := int((*ownedPlates)[plate]) - int(amount)
		if result < 0 {
			err = fmt.Errorf("HomeGym.FreedomUnits:%v %v:%v has < %v plates", h.FreedomUnits, plate, (*ownedPlates)[plate], amount)
			return
		}

		if (*ownedPlates)[plate] > 0 {
			(*ownedPlates)[plate] = (*ownedPlates)[plate] - amount

		} else {
			logger.Printf("No %v plates to remove", plate)
		}

	} else {
		err = fmt.Errorf("%v: is a non-standard weight plate", plate)
	}

	h.CalculateMaxWeight()
	return
}

func (h *HomeGym) CalculateMaxWeight() (err error) {
	ownedPlates, _ := getPlates(h)
	var totalWeight float32

	for plate, count := range maps.All(*ownedPlates) {
		totalWeight = totalWeight + plate*float32(count)
	}
	h.MaxWeight = totalWeight + h.BarWeight
	return
}

func (h *HomeGym) WeightCombos() {
	/*
		Known: purchased plates
		Want: weight(s) to achieve, minimal plates used, even on both sides
		Weight : [bar weight] + [weight plates used]
		225: 45(bar) 45 45 35 35 10 10
// 	*/
	// combos := map[float32][]float32
	// currCombo := []float32
	// var plateCounter = PlateCount

	// if h.FreedomUnits {
	// 	plateCounter = h.Plates.lbs
	// } else {
	// 	plateCounter = h.Plates.kgs
	// }

// 	//recursively find all possible weight combos
// 	//if weight achieved, add to combos

// 	//if not possible, return error
}

func (h *HomeGym) GetCombo(plates PlateCount, desiredWeight float32) (combo []float32, err error) {
	targetWeight := desiredWeight - h.BarWeight
	if desiredWeight < h.BarWeight {
		err = fmt.Errorf("desired weight %v is less than bar weight %v", desiredWeight, h.BarWeight)
		return nil, err
	}

	plateCounterKeys := make([]float32, 0, len(plates))
	for key := range plates {
		plateCounterKeys = append(plateCounterKeys, key)
	}
	slices.Sort(plateCounterKeys)
	slices.Reverse(plateCounterKeys)
	var result float32

	for index := 0; index < len(plateCounterKeys); {
		currWeight := plateCounterKeys[index]
		result = targetWeight - 2*float32(currWeight)

		if result >= 0 && plates[currWeight] >= 2 {
			targetWeight = result
			combo = append(combo, float32(currWeight), float32(currWeight))
			plates[currWeight] = plates[currWeight] - 2
			logger.Println(index, result)
		} else {
			index++
		}

		if targetWeight == 0 {
			break
		}

	}
	return combo, err
}
