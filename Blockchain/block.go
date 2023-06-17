// Package blockchain implements a blockchain data structure
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block represents the block part in the blockchain
/*
Timestamp is the current timestamp (when the block is created)
Data is the actual valuable information containing in the block
PrevBlockHash stores the hash of the previous block
Hash is the hash of the block.
*/
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// NewBlock creates a new block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// SetHash calculates the hash of the block by concatenate the current Block with the previous Block,
// and calculate a SHA-256 hash on the concatenated combination
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

// IsValid checks if the block is valid
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.blocks); i++ {
		currentBlock := bc.blocks[i]
		previousBlock := bc.blocks[i-1]

		// Check if the current block's hash is valid
		if !bytes.Equal(currentBlock.Hash, currentBlock.CalculateHash()) {
			return false
		}

		// Check if the current block is correctly linked to the previous block
		if !bytes.Equal(currentBlock.PrevBlockHash, previousBlock.Hash) {
			return false
		}
	}
	return true
}

// CalculateHash calculates the hash of the block
func (b *Block) CalculateHash() []byte {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	return hash[:]
}
