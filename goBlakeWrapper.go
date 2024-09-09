package goBlake

// #cgo CFLAGS: -I. -IgoBlake
// #cgo LDFLAGS: -L. -LgoBlake -lblake
// #include "blake_hash.h"
import "C"
import (
	"encoding/hex"
	"unsafe"
)

func CalcDecredHash(bytesInput []byte) []byte {
	// avoid of the situation of empty bytes array
	bytesInput = append(bytesInput, byte(0x0))
	input := (*C.char)(unsafe.Pointer(&bytesInput[0]))
	inputLen := C.uint(len(bytesInput) - 1)

	bytesOutput := make([]byte, 32)
	output := (*C.char)(unsafe.Pointer(&bytesOutput[0]))

	C.decred_hash(output, input, inputLen)

	return bytesOutput
}

func CalcDecredHashHex(bytesInput []byte) string {
	bytesRet := CalcDecredHash(bytesInput)
	return hex.EncodeToString(bytesRet)
}

func CalcSiaHash(bytesInput []byte) []byte {
	// avoid of the situation of empty bytes array
	bytesInput = append(bytesInput, byte(0x0))
	input := (*C.char)(unsafe.Pointer(&bytesInput[0]))
	inputLen := C.uint(len(bytesInput) - 1)

	bytesOutput := make([]byte, 32)
	output := (*C.char)(unsafe.Pointer(&bytesOutput[0]))

	C.sia_hash(output, input, inputLen)

	return bytesOutput
}

func CalcSiaHashHex(bytesInput []byte) string {
	bytesRet := CalcSiaHash(bytesInput)
	return hex.EncodeToString(bytesRet)
}
