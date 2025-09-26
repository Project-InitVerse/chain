package bsc

import (
	"github.com/Project-InitVerse/chain/rlp"
)

// enrEntry is the ENR entry which advertises `bsc` protocol on the discovery.
type enrEntry struct {
	// Ignore additional fields (for forward compatibility).
	Rest []rlp.RawValue `rlp:"tail"`
}

// ENRKey implements enr.Entry.
func (e enrEntry) ENRKey() string {
	return "bsc"
}
