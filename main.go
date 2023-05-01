package main

import (
	"dragon-city-utils/calculators/elements"
	"dragon-city-utils/calculators/orb_recall"
	"fmt"
)

func main() {
	strongs, _ := elements.CalculateStrongs([]string{"terra", "legend", "time"})
	fmt.Println("Strongs: ", strongs)

	weaknesses := elements.CalculateWeaknesses("primal")
	fmt.Println("Weaknesses: ", weaknesses)

	orbRecallGain, _ := orb_recall.CalculateOrbRecallGain(15, 2)
	fmt.Print("Orb recall gain: ", orbRecallGain)
}