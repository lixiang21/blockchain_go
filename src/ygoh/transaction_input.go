package main

import "bytes"

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte // 交易ID
	Vout      int    // 存储该输出在这笔交易中所有输出的索引（因为一笔交易可能有多个输出，需要有信息指明是具体的哪一个）
	Signature []byte
	PubKey    []byte
}

// UsesKey checks whether the address initiated the transaction
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}
