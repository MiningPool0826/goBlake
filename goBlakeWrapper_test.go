package goBlake

import (
	"fmt"
	"testing"
)

// 776484204c66ec4894d5a3879aeddb3772cac5fc2795ed26d9ef2c68f73764cc
func TestCalcDecredHashHex(t *testing.T) {
	input := make([]byte, 40)
	for i := 0; i < 40; i++ {
		input[i] = 0x0
	}
	output := CalcDecredHashHex(input)
	fmt.Println(output)
}

// 4a1931803561f431decab002e7425f0a8531d5e456a1a47fd9998a2530c0f800
func TestCalcSiaHashHex(t *testing.T) {
	input := make([]byte, 40)
	for i := 0; i < 40; i++ {
		input[i] = 0x0
	}
	output := CalcSiaHashHex(input)
	fmt.Println(output)
}
