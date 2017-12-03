package blockchain

import (
	"time"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64 //时间戳
	Data          []byte //数据
	PrevBlockHash []byte //前面的块hash
	Hash          []byte //
	Nonce         int  //随机数
}

// NewBlock creates and returns Block
func NewBlock(data string, prevBlockHash []byte) *Block {

	//分配一个块
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	//工作量证明
	pow := NewProofOfWork(block)

	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}