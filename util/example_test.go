package util_test

import (
	"fmt"
	"math"
	"math/big"

	"github.com/romxxxx/nexepad/util/difficulty"

	"github.com/romxxxx/nexepad/util"
)

func ExampleAmount() {

	a := util.Amount(0)
	fmt.Println("Zero Sompi:", a)

	a = util.Amount(1e8)
	fmt.Println("100,000,000 Sompi:", a)

	a = util.Amount(1e5)
	fmt.Println("100,000 Sompi:", a)
	// Output:
	// Zero Sompi: 0 nexe
	// 100,000,000 Sompi: 1 nexe
	// 100,000 Sompi: 0.001 nexe
}

func ExampleNewAmount() {
	amountOne, err := util.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := util.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := util.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := util.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 nexe
	// 0.01234567 nexe
	// 0 nexe
	// invalid nexepa amount
}

func ExampleAmount_unitConversions() {
	amount := util.Amount(44433322211100)

	fmt.Println("Sompi to knexe:", amount.Format(util.AmountKilonexe))
	fmt.Println("Sompi to nexe:", amount)
	fmt.Println("Sompi to Millinexe:", amount.Format(util.AmountMillinexe))
	fmt.Println("Sompi to Micronexe:", amount.Format(util.AmountMicronexe))
	fmt.Println("Sompi to Sompi:", amount.Format(util.AmountSompi))

	// Output:
	// Sompi to knexe: 444.333222111 knexe
	// Sompi to nexe: 444333.222111 nexe
	// Sompi to Millinexe: 444333222.111 mnexe
	// Sompi to Micronexe: 444333222111 μnexe
	// Sompi to Sompi: 44433322211100 Sompi
}

// This example demonstrates how to convert the compact "bits" in a block header
// which represent the target difficulty to a big integer and display it using
// the typical hex notation.
func ExampleCompactToBig() {
	bits := uint32(419465580)
	targetDifficulty := difficulty.CompactToBig(bits)

	// Display it in hex.
	fmt.Printf("%064x\n", targetDifficulty.Bytes())

	// Output:
	// 0000000000000000896c00000000000000000000000000000000000000000000
}

// This example demonstrates how to convert a target difficulty into the compact
// "bits" in a block header which represent that target difficulty .
func ExampleBigToCompact() {
	// Convert the target difficulty from block 300000 in the bitcoin
	// main chain to compact form.
	t := "0000000000000000896c00000000000000000000000000000000000000000000"
	targetDifficulty, success := new(big.Int).SetString(t, 16)
	if !success {
		fmt.Println("invalid target difficulty")
		return
	}
	bits := difficulty.BigToCompact(targetDifficulty)

	fmt.Println(bits)

	// Output:
	// 419465580
}
