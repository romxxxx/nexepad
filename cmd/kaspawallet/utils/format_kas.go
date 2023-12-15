package utils

import (
	"fmt"

	"github.com/nexepanet/nexepad/domain/consensus/utils/constants"
)

// Formatnexe takes the amount of sompis as uint64, and returns amount of nexe with 8  decimal places
func Formatnexe(amount uint64) string {
	res := "                   "
	if amount > 0 {
		res = fmt.Sprintf("%19.8f", float64(amount)/constants.SompiPernexepa)
	}
	return res
}
