package main

import (
	"fmt"
	"ctsWo"
)

func main () {
	fmt.Printf("Main\n")

	power, hr := ctsWo.EnduranceMiles(200, 150)

	fmt.Printf("Endurance Power: %g - %g\n", power[0], power[1]) 
	fmt.Printf("Endurance Heart Rate: %g - %g\n", hr[0], hr[1]) 
}
