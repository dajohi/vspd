package database

import (
	"encoding/json"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

type Receipt struct {
	Request  interface{} `json:"request"`
	Response interface{} `json:"response"`
}

func (vdb *VspDatabase) AddReceipt(ticketHash string, receipt Receipt) error {
	return vdb.db.Update(func(tx *bolt.Tx) error {
		receiptsBkt := tx.Bucket(vspBktK).Bucket(receiptsBktK)

		receiptBytes, err := json.Marshal(receipt)
		if err != nil {
			return fmt.Errorf("could not marshal ticket: %v", err)
		}

		hashBytes := []byte(ticketHash)
		return receiptsBkt.Put(hashBytes, receiptBytes)
	})
}
