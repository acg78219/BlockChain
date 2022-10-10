package main

import (
	"BlockChain/part1-basic-prototype/BLC"
	"fmt"
)

func main() {
	block := BLC.CreateFirstBlock("xlao block")
	fmt.Println(block)
}
