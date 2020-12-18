package asdefaults

import (
	"math/big"

	"github.com/google/go-cmp/cmp"
)

func numberComparer() cmp.Option {
	return cmp.Comparer(func(i, j *big.Float) bool {
		return (i == nil && j == nil) || (i != nil && j != nil && i.Cmp(j) == 0)
	})
}

var cmpOpts = []cmp.Option{
	numberComparer(),
}
