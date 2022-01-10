package coder

import (
	"fmt"
	"github.com/klauspost/reedsolomon"
)

func encode(data []byte) (splits [][][]byte) {
	enc, _ := reedsolomon.New(5, 3) //change dataShards, parityShards values !
	shards, _ := enc.Split(data)
	err := enc.Encode(shards)
	if err != nil {
		panic(err)
	}

	// Check that it verifies
	ok, err := enc.Verify(shards)
	if ok && err == nil {
		fmt.Println("encode ok")
	}

	split1 := make([][]byte, 8)
	split2 := make([][]byte, 8)
	split3 := make([][]byte, 8)

	// Split the shards
	// find  x!
	/*
	for i := range shards {
		split1[i] = shards[i][:x]
		split2[i] = shards[i][:x:]
		split3[i] = shards[i][:total-x]
	}

	 */
	splits = append(splits, split1, split2, split3)
	return splits
}


