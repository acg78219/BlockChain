package BLC

type BlockChain struct {
	Blocks []*Block
}

// CreateBlockChain 创建区块链， data 是首节点的数据
func CreateBlockChain(data string) *BlockChain {
	// 区块链中至少要有一个节点 block
	block := CreateFirstBlock(data)
	return &BlockChain{[]*Block{block}}
}

func (bl *BlockChain) AddBlockToBlockChain(data string) {
	// 获取区块链中最后一个节点的 height 和 hash
	height := bl.Blocks[len(bl.Blocks)-1].Height
	prevHash := bl.Blocks[len(bl.Blocks)-1].Hash
	// 新建 block
	newBlock := NewBlock(height+1, prevHash, data)
	// 追加到区块链中
	bl.Blocks = append(bl.Blocks, newBlock)
}
