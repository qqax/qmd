package qmd

import (
	"bytes"
	"testing"
)

func TestMd(t *testing.T) {
	checkSum := md([]byte("These pretzels are making me thirsty."))
	testSum := [16]byte{176, 128, 78, 201, 103, 244, 133, 32, 105, 118, 98, 162, 4, 245, 254, 114}
	if checkSum != testSum {
		t.Fatalf(`md check sum is %x, should be %x`, checkSum, testSum)
	}

	store, salted, saltIndex, err := checkSum.salted(4, 12)
	if err != nil {
		return
	}

	restMd := salted[saltIndex : saltIndex+16-len(store)]
	newMd := append(store, salted[saltIndex:saltIndex+16-len(store)]...)

	if *(*[16]byte)(newMd) != testSum {
		t.Fatalf(`joined array is %x, should be %x`, newMd, restMd)
	}
	if !bytes.Equal(store, testSum[:len(store)]) {
		t.Fatalf(`array in store is %x, should be %x`, store, testSum[:len(store)])
	}
	if !bytes.Equal(restMd, testSum[len(store):]) {
		t.Fatalf(`rest array is %x, should be %x`, restMd, testSum[len(store):])
	}
}
