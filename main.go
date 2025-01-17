package main

import (
	"bytes"
	"fmt"
	"log"
)

var standardWeights = StandardWeights{}

var (
	buf bytes.Buffer
	logger = log.New(&buf, "",log.Lshortfile | log.Ltime)
)

func ReportErrors(x bool) {
	if x {
		fmt.Print(&buf)

	}
}

func main(){
	standardWeights.init()
	homeGym := HomeGym{
		FreedomUnits: true,
		PlateCounter: PurchasedPlates{make(map[float32]byte), make(map[float32]byte)},
		BarWeight: 45,
	}
	numPlates := 10
	homeGym.BuyPlates(2.5, byte(numPlates))
	homeGym.BuyPlates(5, byte(numPlates))
	homeGym.BuyPlates(45, 2)
	homeGym.BuyPlates(35, byte(numPlates))
	homeGym.BuyPlates(10, byte(numPlates))
	homeGym.BuyPlates(25, byte(numPlates))

	logger.Println(homeGym)
	
	ReportErrors(true)
}

