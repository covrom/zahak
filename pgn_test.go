package main

import (
	"fmt"
	"testing"
)

func TestPGNParsing(t *testing.T) {
	fen := "r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q2/PPPBBPpP/R3K2R b KQkq - 0 1"
	game := FromFen(fen, true)
	actual := game.position.ParseMoves([]string{"g2h1q", "e2f1", "   ", "\n\t", "h8h2"})
	expected := []*Move{
		&Move{G2, H1, Queen, Capture | Check},
		&Move{E2, F1, NoType, 0},
		&Move{H8, H2, NoType, Capture},
	}
	if !equalMoves(expected, actual) {
		fmt.Println("Got:")
		for _, i := range expected {
			fmt.Println(i.ToString(), i.promoType, i.moveTag)
		}
		fmt.Println("Expected:")
		for _, i := range actual {
			fmt.Println(i.ToString(), i.promoType, i.moveTag)
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", len(expected), len(actual)))
	}
}
