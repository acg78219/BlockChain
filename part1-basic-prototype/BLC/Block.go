package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

// 区块链中的单位 block
type Block struct {
	// 1. 高度
	Height int64
	// 2. 上一个 block 的 hash
	PrevBlockHash []byte
	// 3. 数据
	Data []byte
	// 4. 时间戳
	Timestamp int64
	// 5. hash
	Hash []byte
	// 6. 工作量证明
	Nonce int64
}

// SetHash 设置 Block 的 hash
// 这个 hash 是将 Block 中的所有属性组合起来再 Sum256 生成最终 Hash
func (block *Block) SetHash() {
	// 1. 将 Height 转换为 []byte
	heightBytes := IntToHex(block.Height)

	// 2. 将时间戳转换为 []byte (和 Height 转换的方式不一样)
	// 先将 int64 转换为 string；再转换为 []byte
	timeString := strconv.FormatInt(block.Timestamp, 2) // 2 表示转换为二进制
	timeBytes := []byte(timeString)

	// 3. 组合所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})

	// 4. 生成 Hash
	hash := sha256.Sum256(blockBytes)
	// 由于 Sum256 返回一个 [32]byte，但是我们的 Hash 类型是 []byte 切片，所以需要转换一下
	block.Hash = hash[:]
}

// NewBlock 新建一个 Block
func NewBlock(height int64, prevBlock []byte, data string) *Block {
	block := &Block{
		Height:        height,
		PrevBlockHash: prevBlock,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
	}
	// 工作量证明
	pow := NewProofOfWork(block)
	// 不断挖矿进行工作
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block
}

// CreateFirstBlock 创建 BlockChain 中的第一个 Block
func CreateFirstBlock(data string) *Block {
	return NewBlock(1,
		// Hash 是 64 位的，而 byte 中一个数字表示 2 位，所以需要 32 个 0
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		data)
}

// Serialize 将 block 序列化为字节数组
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	
	encoder := gob.NewEncoder(&result)

	// 将 block 序列化
	err := encoder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(blockBytes []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	// 将 block 反序列化
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
