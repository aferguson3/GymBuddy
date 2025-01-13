package main

import (
	"bytes"
	"fmt"
	"log"

)

var StandardWeights_ = StandardWeights{}

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
	StandardWeights_.init()
	homeGym := HomeGym{
		FreedomUnits: true,
		Plates: PurchasedPlates{make(map[string]byte), make(map[string]byte)},
	}

	logger.Println(homeGym)
	ReportErrors(true)
}

