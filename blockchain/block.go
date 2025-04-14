package blockchain

import "crypto/sha256"

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}
// logic to derive blockchain hash
func (b *Block) DeriveHash() {
	info := append(b.PrevHash, b.Data...)
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
