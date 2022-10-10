package main

import (
	"BlockChain/part1-basic-prototype/BLC"
	"fmt"
)

func main() {
	bc := BLC.CreateBlockChain("xlao block")
	bc.AddBlockToBlockChain("send 100RMB to xlao")
	bc.AddBlockToBlockChain("send 50RMB to xlao")

	for i, v := range bc.Blocks {
		fmt.Println(i, ":", v)
	}
}
