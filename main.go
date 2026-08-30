package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var x uint32 = 0b_10000000_00000000_00100100_00000000

	buf1 := make([]uint8, 0, 4)
	buf2 := make([]uint8, 0, 4)
	
	b := binary.LittleEndian
	buf1 = b.AppendUint32(buf1, x)
	fmt.Println(buf1)

	b2 := binary.BigEndian
	buf2 = b2.AppendUint32(buf2, x)
	fmt.Println(buf2)
}