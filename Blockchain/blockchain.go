package Blockchain

type Blockchain struct {
	blocks []*Block
}

// NewBlockchain creates a new blockchain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// NewGenesisBlock creates a new blockchain, by creating the first block
/*
In any blockchain, there must be at least one block, the first in the chain, this block called genesis block
*/
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
