package engine

import (
	"testing"
)

func TestMakeMove(t *testing.T) {
	game := FromFen("rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1")
	move := NewMove(F3, G4, WhiteBishop, NoPiece, NoType, 0)
	game.position.MakeMove(move)
	fen := game.Fen()
	expected := "rnbqkbnr/pPp1pppp/4P3/3pP3/4p1B1/6N1/PP3PPP/RNBQK2R b KQkq - 1 1"
	if fen != expected {
		t.Errorf("Move was not generated properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", expected)
	}
}

func TestMakeMoveDoublePushPawn(t *testing.T) {
	game := FromFen("rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1")
	move := NewMove(H2, H4, WhitePawn, NoPiece, NoType, 0)
	game.position.MakeMove(move)
	fen := game.Fen()
	expected := "rnbqkbnr/pPp1pppp/4P3/3pP3/4p2P/5BN1/PP3PP1/RNBQK2R b KQkq h3 0 1"
	if fen != expected {
		t.Errorf("Move was not generated properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", expected)
	}
}

func TestMakeMoveCapture(t *testing.T) {
	game := FromFen("rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1")
	move := NewMove(F3, E4, WhiteBishop, BlackPawn, NoType, Capture)
	game.position.MakeMove(move)
	fen := game.Fen()
	expected := "rnbqkbnr/pPp1pppp/4P3/3pP3/4B3/6N1/PP3PPP/RNBQK2R b KQkq - 0 1"
	if fen != expected {
		t.Errorf("Move was not generated properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", expected)
	}
}

func TestMakeMoveCastling(t *testing.T) {
	game := FromFen("rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1")
	move := NewMove(E1, G1, WhiteKing, NoPiece, NoType, KingSideCastle)
	game.position.MakeMove(move)
	fen := game.Fen()
	expected := "rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQ1RK1 b kq - 1 1"
	if fen != expected {
		t.Errorf("Move was not generated properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", expected)
	}
}

func TestMakeMoveEnPassant(t *testing.T) {
	game := FromFen("rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1")
	move := NewMove(E5, D6, WhitePawn, BlackPawn, NoType, EnPassant|Capture)
	game.position.MakeMove(move)
	fen := game.Fen()
	expected := "rnbqkbnr/pPp1pppp/3PP3/8/4p3/5BN1/PP3PPP/RNBQK2R b KQkq - 0 1"
	if fen != expected {
		t.Errorf("Move was not generated properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", expected)
	}
}

func TestMakeMovePromotion(t *testing.T) {
	game := FromFen("rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1")
	move := NewMove(B7, A8, WhitePawn, BlackRook, Queen, Capture)
	game.position.MakeMove(move)
	fen := game.Fen()
	expected := "Qnbqkbnr/p1p1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R b KQk - 0 1"
	if fen != expected {
		t.Errorf("Move was not generated properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", expected)
	}
}

func TestUnMakeMove(t *testing.T) {
	startFen := "rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1"
	game := FromFen(startFen)
	startHash := game.position.Hash()
	move := NewMove(F3, G4, WhiteBishop, NoPiece, NoType, 0)
	ep, tag, hc, _ := game.position.MakeMove(move)
	game.position.UnMakeMove(move, tag, ep, hc)
	fen := game.Fen()
	if fen != startFen {
		t.Errorf("Move was not undone properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", startFen)
	}
	newHash := game.position.Hash()
	if startHash != newHash {
		t.Errorf("Move was not undone properly\nGot hash: %d\n", newHash)
		t.Errorf("But expected: %d\n", startHash)
	}
}

func TestUnMakeMoveDoublePushPawn(t *testing.T) {
	startFen := "rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1"
	game := FromFen(startFen)
	startHash := game.position.Hash()
	move := NewMove(H2, H4, WhitePawn, NoPiece, NoType, 0)
	ep, tag, hc, _ := game.position.MakeMove(move)
	game.position.UnMakeMove(move, tag, ep, hc)
	fen := game.Fen()
	if fen != startFen {
		t.Errorf("Move was not undone properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", startFen)
	}
	newHash := game.position.Hash()
	if startHash != newHash {
		t.Errorf("Move was not undone properly\nGot hash: %d\n", newHash)
		t.Errorf("But expected: %d\n", startHash)
	}
}

func TestUnMakeMoveCapture(t *testing.T) {
	startFen := "rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1"
	game := FromFen(startFen)
	startHash := game.position.Hash()
	move := NewMove(F3, E4, WhiteBishop, BlackPawn, NoType, Capture)
	ep, tag, hc, _ := game.position.MakeMove(move)
	game.position.UnMakeMove(move, tag, ep, hc)
	fen := game.Fen()
	if fen != startFen {
		t.Errorf("Move was not undone properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", startFen)
	}
	newHash := game.position.Hash()
	if startHash != newHash {
		t.Errorf("Move was not undone properly\nGot hash: %d\n", newHash)
		t.Errorf("But expected: %d\n", startHash)
	}
}

func TestUnMakeMoveCastling(t *testing.T) {
	startFen := "rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1"
	game := FromFen(startFen)
	startHash := game.position.Hash()
	move := NewMove(E1, G1, WhiteKing, NoPiece, NoType, KingSideCastle)
	ep, tag, hc, _ := game.position.MakeMove(move)
	game.position.UnMakeMove(move, tag, ep, hc)
	fen := game.Fen()
	if fen != startFen {
		t.Errorf("Move was not undone properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", startFen)
	}
	newHash := game.position.Hash()
	if startHash != newHash {
		t.Errorf("Move was not undone properly\nGot hash: %d\n", newHash)
		t.Errorf("But expected: %d\n", startHash)
	}
}

func TestUnMakeMoveEnPassant(t *testing.T) {
	startFen := "rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1"
	game := FromFen(startFen)
	startHash := game.position.Hash()
	move := NewMove(E5, D6, WhitePawn, BlackPawn, NoType, EnPassant|Capture)
	ep, tag, hc, _ := game.position.MakeMove(move)
	game.position.UnMakeMove(move, tag, ep, hc)
	fen := game.Fen()
	if fen != startFen {
		t.Errorf("Move was not undone properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", startFen)
	}
	newHash := game.position.Hash()
	if startHash != newHash {
		t.Errorf("Move was not undone properly\nGot hash: %d\n", newHash)
		t.Errorf("But expected: %d\n", startHash)
	}
}

func TestUnMakeMovePromotion(t *testing.T) {
	startFen := "rnbqkbnr/pPp1pppp/4P3/3pP3/4p3/5BN1/PP3PPP/RNBQK2R w KQkq d6 0 1"
	game := FromFen(startFen)
	startHash := game.position.Hash()
	move := NewMove(B7, A8, WhitePawn, BlackRook, Queen, Capture)
	ep, tag, hc, _ := game.position.MakeMove(move)
	game.position.UnMakeMove(move, tag, ep, hc)
	fen := game.Fen()
	if fen != startFen {
		t.Errorf("Move was not undone properly\nGot: %s\n", fen)
		t.Errorf("But expected: %s\n", startFen)
	}
	newHash := game.position.Hash()
	if startHash != newHash {
		t.Errorf("Move was not undone properly\nGot hash: %d\n", newHash)
		t.Errorf("But expected: %d\n", startHash)
	}
}

func TestIncrementalEval(t *testing.T) {
	for _, pos := range positions {
		// compute the materials first first
		originalEval := getEval(pos)
		for _, mov := range pos.PseudoLegalMoves() {
			if ep, tg, hc, ok := pos.MakeMove(mov); ok {
				incrementalEval := getEval(pos)
				pos.MaterialAndPSQT()
				fromScratchEval := getEval(pos)
				pos.UnMakeMove(mov, tg, ep, hc)
				decrementalEval := getEval(pos)
				if incrementalEval != fromScratchEval {
					t.Errorf("Updated eval != Fresh eval ->\nMov: %v, Fresh Eval: %d, Incremental Eval: %d, \nPos:%v\n", mov.ToString(), fromScratchEval, incrementalEval, pos.Fen())
				}
				if decrementalEval != originalEval {
					t.Errorf("Undone eval (reverse from unmake) != original eval ->\nMov: %v, Original Eval: %d, Decremental Eval: %d, \nPos:%v\n", mov.ToString(), originalEval, decrementalEval, pos.Fen())
				}
			}
		}
	}
}

func getEval(p *Position) int16 {
	var wcp, bcp int16
	w := []int16{100, 300, 350, 400, 900, 2000}
	for i := 0; i < 6; i++ {
		wcp += p.MaterialsOnBoard[i] * w[i]
		bcp += p.MaterialsOnBoard[i+6] * w[i]
	}
	wcp += p.WhiteEndgamePSQT
	wcp += p.WhiteMiddlegamePSQT
	bcp += p.BlackEndgamePSQT
	bcp += p.BlackMiddlegamePSQT

	return wcp - bcp
}
