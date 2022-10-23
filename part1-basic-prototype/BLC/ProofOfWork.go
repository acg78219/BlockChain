package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 256位 hash 中前面至少要有 16 位的 0，表示难度
const targetBit = 16

type ProofOfWork struct {
	// 所在的区块
	Block *Block
	// 大数据存储
	Target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	// 将 pow 中的 Block 转换为 byte[]
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{})

	return data
}

func (pow *ProofOfWork) Run() ([]byte, int64) {

	// 2. 生成 hash
	// 3. 判断 hash 是否有效，如果满足 => 跳出

	nonce := 0
	var hashInt big.Int
	var hash [32]byte
	// 不断死循环挖矿，直到 hash 符合要求
	for {
		// 1. 将 Block 的属性转换为 []byte
		dataBytes := pow.prepareData(nonce)

		// 2. 生成 hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("%x\n", hash)

		// 3. 将 hash 转换为 big.Int
		hashInt.SetBytes(hash[:])

		// 4. 和 target 计算比较：如果 target > hashInt，返回1，说明符合要求
		if pow.Target.Cmp(&hashInt) == 1 {
			break
		}

		nonce = nonce + 1
	}
	return hash[:], int64(nonce)
}

// NewProofOfWork 返回新的工作量证明
func NewProofOfWork(block *Block) *ProofOfWork {
	// 工作量证明的难度
	// 创建一个 big.Int 类型的数值 1
	target := big.NewInt(1)
	// 将 1 左移 256 - targetBit
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}
