package ctsWo

import (
	"fmt"
)

func EnduranceMiles (power, hr int) ([]float64, []float64) {
	fmt.Printf("EnduranceMiles Func\n")
	var ctsPower []float64
	var ctsHr []float64

	ctsPower = append (ctsPower, float64(power) * 0.45)
	ctsPower = append (ctsPower, float64(power) * 0.73)

	ctsHr = append (ctsHr, float64(hr) * 0.5)
	ctsHr = append (ctsHr, float64(hr) * 0.91)

	return ctsPower, ctsHr
}
