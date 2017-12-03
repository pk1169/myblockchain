package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex converts an int64 to a byte array
// 将64位整型转化成二进制八位整型数组
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
