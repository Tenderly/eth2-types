package types

import (
	fssz "github.com/ferranbt/fastssz"
)

// SSZUint64 is a bytes slice that satisfies the fast-ssz interface.
type SSZBytes []byte

func (s SSZBytes) GetTree() (*fssz.Node, error) {
	panic("not implemented")
}

// HashTreeRoot hashes the uint64 object following the SSZ standard.
func (b *SSZBytes) HashTreeRoot() ([32]byte, error) {
	return fssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith hashes the uint64 object with the given hasher.
func (b *SSZBytes) HashTreeRootWith(hh fssz.HashWalker) error {
	indx := hh.Index()
	hh.PutBytes(*b)
	hh.Merkleize(indx)
	return nil
}
