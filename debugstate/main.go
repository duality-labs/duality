package main

import (
	cometbftdb "github.com/cometbft/cometbft-db"
)

func main() {
	// specify your db path
	dbPath := "/root/.duality/data"

	// Open the database
	db, err := cometbftdb.NewDB("application", "goleveldb", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepare for iteration over the whole database
	iterator, err := db.Iterator(nil, nil)
	if err != nil {
		panic(err)
	}
	defer iterator.Close()

	// Initialize key count
	root := NewNode(0)
	keyCount := 0

	// Iterate over the keys in the database
	for ; iterator.Valid(); iterator.Next() {
		root.Insert(iterator.Key())
		keyCount += 1
		if keyCount%10000 == 0 {
			root.PrintStats()
		}
	}
}
