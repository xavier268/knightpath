package main

import (
	"fmt"
	"math/bits"
)

var CanDo [64]uint64 // for each square, from 0 to 63, bit mask for the acessible squares.

func init() {
	fmt.Println("Starting precomputation of legal moves")
	precomputeCanDo()
	fmt.Println("Finished precomputation of legal moves")
	for _, m := range CanDo {
		fmt.Printf("%064b\t%d\n", m, bits.OnesCount64(m))
	}
}

func precomputeCanDo() {
	for s := range CanDo {
		for i := 0; i < 64; i++ {
			if linked(s, i) {
				CanDo[s] = CanDo[s] | (1 << i) // set bit # i
			}
		}
	}
}

// get the coordinates of a given square.
func sToCoord(square int) (x int, y int) {
	x = square % 8
	y = square / 8
	return x, y
}

/*
// get square for a given coordinate
func coordToS(x, y int) int {
	return x + 8*y
}
*/

// True iff knight can jump from a to b
func linked(sa, sb int) bool {
	xa, ya := sToCoord(sa)
	xb, yb := sToCoord(sb)
	return (xa-xb)*(xa-xb)+(ya-yb)*(ya-yb) == 5
}
