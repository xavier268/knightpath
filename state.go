package main

import (
	"fmt"
	"math/bits"
)

const VERB = 1

// States
type State struct {
	occ  uint64 // bit map of all positions occupied, incl current position
	try  uint64 // bit map of already tried next moves
	pos  int    // current position of knight
	prev *State // previous state
}

// New start state for a given position
func NewState(pos int) *State {
	return &State{
		occ: 1 << pos,
		pos: pos,
	}
}

var (
	ErrBlocked error = fmt.Errorf("blocked")
)

const FULL = 0xFFFFFFFF

func Solve(from *State) error {

	maxDepth := 0

	//time.Sleep(time.Second / 4)
	for {
		for from.PossibleMove() != 0 {

			for i := 0; i < 64; i++ {
				if (uint64(1)<<i)&from.PossibleMove() != 0 { // try all acceptable moves from 'from'
					from.try = from.try | (1 << i) // mark as tried
					if bits.OnesCount64(from.occ) > maxDepth {
						maxDepth = bits.OnesCount64(from.occ)
						if VERB >= 1 {
							fmt.Printf("Max depth : %d\n", maxDepth)
						}
					}
					if VERB >= 2 {
						fmt.Printf("trying %d,\tocc:%064b\tdepth:%d/%d\n", i, from.occ, bits.OnesCount64(from.occ), maxDepth)
					}
					//time.Sleep(time.Second / 4)

					ns := NewState(i)
					ns.occ = ns.occ | from.occ
					ns.prev = from

					// iterate with ns
					from = ns
					break
				}
			}
		}
		if from.PossibleMove() == 0 && !from.Solved() { // back tracking
			from = from.prev
			if from == nil {
				return ErrBlocked
			}
		} else {
			return nil
		}

		if from.Solved() {
			return nil
		}
		// else loop from top ...
	}
}

func (from *State) Solved() bool {
	return from.occ == FULL
}

func (from *State) PossibleMove() uint64 {
	return ((from.occ | from.try) ^ FULL) & CanDo[from.pos]
}
