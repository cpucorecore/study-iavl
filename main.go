package main

import (
	"fmt"
	"log"

	"github.com/cosmos/iavl"
	db "github.com/tendermint/tm-db"
)

func main() {
	tree, err := iavl.NewMutableTree(db.NewMemDB(), 0)
	if err != nil {
		log.Fatal(err)
	}

	tree.Set([]byte("e"), []byte{5})
	tree.Set([]byte("d"), []byte{4})
	tree.Set([]byte("c"), []byte{3})
	tree.Set([]byte("b"), []byte{2})
	tree.Set([]byte("a"), []byte{1})

	rootHash, version, err := tree.SaveVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Saved version %v with root hash %x\n", version, rootHash)

	// Output tree structure, including all node hashes (prefixed with 'n')
	fmt.Println(tree.String())
}
