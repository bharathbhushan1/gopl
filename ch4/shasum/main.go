package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n", c1, c2)
	fmt.Printf("%d\n", popCount(c1, c2))
	fmt.Printf("%d\n", popCount2(c1, c2))
}

func popCount(hash1 [32]byte, hash2 [32]byte) int {
	var unequalBits int
	for i := 0; i < 32; i++ {
		a, b := hash1[i], hash2[i]
		var j uint8
		for ; j < 8; j++ {
			a1 := a & (1 << j)
			b1 := b & (1 << j)
			if a1 != b1 {
				unequalBits++
			}
		}
	}
	return unequalBits
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount2(hash1 [32]byte, hash2 [32]byte) int {
	var unequalBits int
	for i := 0; i < 32; i++ {
		a, b := hash1[i], hash2[i]
		r := a ^ b
		unequalBits += int(pc[r])
	}
	return unequalBits
}
