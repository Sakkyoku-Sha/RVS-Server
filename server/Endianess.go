package server

import (
	"encoding/binary"
	"unsafe"
)

//GetByteOrder determines what "Endianess" is being used
//THIS SHOULD ONLY EVER BE USED WHEN CONVERTING FROM THE TCP INPUT to InputRequests
func GetByteOrder() binary.ByteOrder {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		return binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		return binary.BigEndian
	default:
		panic("Could not determine native endianness.")
	}
}
