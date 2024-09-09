package goBlake

import (
	"fmt"
	"testing"
)

func TestCalcDecredHashHex(t *testing.T) {
	input := make([]byte, 180)
	for i := 0; i < 180; i++ {
		input[i] = 0x0
	}
	output := CalcDecredHashHex(input)
	fmt.Println(output)
}

func TestCalcSiaHashHex(t *testing.T) {
	input := make([]byte, 80)
	for i := 0; i < 80; i++ {
		input[i] = 0x0
	}
	output := CalcSiaHashHex(input)
	fmt.Println(output)
}
