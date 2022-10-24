package main

import (
	"BlockChain/part1-basic-prototype/BLC"
	"fmt"
)

func main() {
	bc := BLC.CreateBlockChain("xlao block")
	bc.AddBlockToBlockChain("send 100RMB to xlao")
	bc.AddBlockToBlockChain("send 50RMB to xlao")

	// 验证序列化和反序列化是否成功
	for _, block := range bc.Blocks {
		fmt.Println("origin Block.Hash: ", block.Hash)
		//fmt.Println(i, ":", v)
		bytes := block.Serialize()
		fmt.Println("Serialize: ", bytes)

		deBlock := BLC.DeserializeBlock(bytes)
		fmt.Println("deBlock.Hash: ", deBlock.Hash)
	}
}
