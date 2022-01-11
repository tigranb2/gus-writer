package coder

import (
	"bytes"
	"fmt"
	"github.com/klauspost/reedsolomon"
	"os"
)

func Encode(data []byte) [][]byte {
	enc, _ := reedsolomon.New(2, 1)
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

	return shards
}

func Decode(shards [][]byte) (data []byte) {
	enc, _ := reedsolomon.New(2, 1)

	// Verify the shards
	ok, err := enc.Verify(shards)
	if ok {
		fmt.Println("No reconstruction needed")
	} else {
		fmt.Println("Verification failed. Reconstructing data")
		err = enc.Reconstruct(shards)
		if err != nil {
			fmt.Println("Reconstruct failed -", err)
			os.Exit(1)
		}
		ok, err = enc.Verify(shards)
		if !ok {
			fmt.Println("Verification failed after reconstruction, data likely corrupted.")
			os.Exit(1)
		}
		if err != nil {
			fmt.Println("Verification failed, reconstruction failed -", err)
			os.Exit(1)
		}

	}

	buf := new(bytes.Buffer)
	err = enc.Join(buf, shards, 2)
	if err != nil {
		fmt.Println("Join failed -", err)
		os.Exit(1)
	}

	return buf.Bytes()
}
