package qmd

import (
	"crypto/md5"
	"fmt"
	"github.com/qqax/qrandom"
)

type mdHash [16]byte

func md(p []byte) mdHash {
	return md5.Sum(p)
}

func (mh *mdHash) salted(min, max int) ([]byte, []byte, int, error) {
	if min < 0 || min >= max || max > len(mh)+1 {
		return nil, nil, 0, fmt.Errorf("wrong index")
	}

	mdHashIndex, err := qrandom.IntBetween(int64(min), int64(max))
	if err != nil {
		return nil, nil, 0, err
	}

	saltArray, err := qrandom.CreateByteArray(int(mdHashIndex))
	if err != nil {
		return nil, nil, 0, err
	}

	saltIndex, err := qrandom.IntBetween(0, int64(int(mdHashIndex)))
	if err != nil {
		return nil, nil, 0, err
	}
	saltedSlice := ConcatMultipleSlices(saltArray[:saltIndex], mh[int(mdHashIndex):], saltArray[saltIndex:])

	fmt.Println(mdHashIndex, saltIndex)
	rest := make([]byte, mdHashIndex)
	copy(rest, mh[:int(mdHashIndex)])

	return rest, saltedSlice, int(saltIndex), nil
}
