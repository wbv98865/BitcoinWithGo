/**
区块链迭代器
 */
package main

import (
	"log"

	"github.com/boltdb/bolt"
)

// 用于迭代区块链块的结构
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// 从顶开始返回下一个块
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	i.currentHash = block.PrevBlockHash

	return block
}
