package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index        int64
	Timestamp    int64
	Data         string
	PreviousHash string
	Hash         string
}

func calculateHash(block Block) string {
	record := string(block.Index) + string(block.Timestamp) + block.Data + block.PreviousHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(previousBlock Block, data string) Block {
	var newBlock Block

	newBlock.Index = previousBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.PreviousHash = previousBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

func isBlockValid(newBlock, previousBlock Block) bool {
	if previousBlock.Index+1 != newBlock.Index {
		return false
	}

	if previousBlock.Hash != newBlock.PreviousHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func main() {
	var blockchain []Block

	// Generate the genesis block
	genesisBlock := Block{0, time.Now().Unix(), "Genesis Block", "", ""}
	genesisBlock.Hash = calculateHash(genesisBlock)
	blockchain = append(blockchain, genesisBlock)

	// Add some more blocks to the blockchain
	blockchain = append(blockchain, generateBlock(genesisBlock, "First Block"))
	blockchain = append(blockchain, generateBlock(blockchain[1], "Second Block"))

	// Verify the blockchain
	for i := 1; i < len(blockchain); i++ {
		if isBlockValid(blockchain[i], blockchain[i-1]) {
			fmt.Printf("Block %d is valid\n", i)
		} else {
			fmt.Printf("Block %d is invalid\n", i)
		}
	}
}
