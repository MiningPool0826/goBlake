package goBlake

import (
	"fmt"
	"testing"
)

func TestCalcDecredHashHex(t *testing.T) {
	input := make([]byte, 192)
	for i := 0; i < 192; i++ {
		input[i] = 0x0
	}
	output := CalcDecredHashHex(input)
	fmt.Println(output)
}

func TestCalcSiaHashHex(t *testing.T) {
	input := make([]byte, 192)
	for i := 0; i < 192; i++ {
		input[i] = 0x0
	}
	output := CalcSiaHashHex(input)
	fmt.Println(output)
}
