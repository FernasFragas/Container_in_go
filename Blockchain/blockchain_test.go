package Blockchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockchain(t *testing.T) {
	// Create a new blockchain
	bc := NewBlockchain()

	// Add blocks to the blockchain
	bc.AddBlock("Block 1 Data")
	bc.AddBlock("Block 2 Data")

	// Check the validity of the blockchain
	valid := bc.IsValid()

	assert.True(t, valid)
}
