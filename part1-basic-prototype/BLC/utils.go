package BLC

import (
	"bytes"
	"encoding/binary"
)

// IntToHex 将 int64 转换为 []byte
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}
