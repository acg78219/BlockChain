package BLC

type ProofOfWork struct {
	Block *Block
}

func (pow *ProofOfWork) Run() ([]byte, int64) {
	return nil, 0
}

// NewProofOfWork 返回新的工作量证明
func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{block}
}
