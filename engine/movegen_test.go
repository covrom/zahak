package engine

import (
	"fmt"
	"testing"
)

func TestBishopMoves(t *testing.T) {
	fen := "rnbqkbnr/pPp1pppp/4P3/3pP3/3p4/4B1N1/PP2BPPP/1NRQK2R w Kkq - 0 1"
	g := FromFen(fen)
	ml := NewMoveList(13)
	g.position.slidingQuietMoves(White, Bishop, ml)
	g.position.slidingCaptureMoves(White, Bishop, ml)
	moves := ml.Moves
	expectedMoves := []Move{
		NewMove(E2, F1, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, F3, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, G4, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, H5, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, D3, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, C4, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, B5, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, A6, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E3, D2, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E3, F4, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E3, G5, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E3, H6, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E3, D4, WhiteBishop, BlackPawn, NoType, Capture),
	}
	expectedLen := len(expectedMoves)
	if len(moves) != expectedLen || !equalMoves(expectedMoves, moves) {
		fmt.Println("Got:")
		for _, i := range moves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(moves)))
	}
}

func TestRookMoves(t *testing.T) {
	fen := "rnkqbbnr/ppp1pppp/4P3/3pP3/3P4/4B1N1/PP2BPPP/1NRQK2R w Kkq - 0 1"
	g := FromFen(fen)
	ml := NewMoveList(8)
	g.position.slidingQuietMoves(White, Rook, ml)
	g.position.slidingCaptureMoves(White, Rook, ml)
	moves := ml.Moves
	expectedMoves := []Move{
		NewMove(H1, G1, WhiteRook, NoPiece, NoType, 0),
		NewMove(H1, F1, WhiteRook, NoPiece, NoType, 0),
		NewMove(C1, C2, WhiteRook, NoPiece, NoType, 0),
		NewMove(C1, C3, WhiteRook, NoPiece, NoType, 0),
		NewMove(C1, C4, WhiteRook, NoPiece, NoType, 0),
		NewMove(C1, C5, WhiteRook, NoPiece, NoType, 0),
		NewMove(C1, C6, WhiteRook, NoPiece, NoType, 0),
		NewMove(C1, C7, WhiteRook, BlackPawn, NoType, Capture),
	}
	expectedLen := len(expectedMoves)
	if len(moves) != expectedLen || !equalMoves(expectedMoves, moves) {
		fmt.Println("Got:")
		for _, i := range moves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(moves)))
	}
}

func TestQueenMoves(t *testing.T) {
	fen := "rnbqkbnr/pPp1pppp/4P3/3pP3/3p4/4B1N1/PP2BPPP/1NRQK2R w Kkq - 0 1"
	g := FromFen(fen)
	ml := NewMoveList(6)
	g.position.slidingCaptureMoves(White, Queen, ml)
	g.position.slidingQuietMoves(White, Queen, ml)
	moves := ml.Moves
	expectedMoves := []Move{
		NewMove(D1, D2, WhiteQueen, NoPiece, NoType, 0),
		NewMove(D1, D3, WhiteQueen, NoPiece, NoType, 0),
		NewMove(D1, D4, WhiteQueen, BlackPawn, NoType, Capture),
		NewMove(D1, C2, WhiteQueen, NoPiece, NoType, 0),
		NewMove(D1, B3, WhiteQueen, NoPiece, NoType, 0),
		NewMove(D1, A4, WhiteQueen, NoPiece, NoType, 0),
	}
	expectedLen := len(expectedMoves)
	if len(moves) != expectedLen || !equalMoves(expectedMoves, moves) {
		fmt.Println("Got:")
		for _, i := range moves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(moves)))
	}
}

func TestKingMoves(t *testing.T) {
	fen := "rnbqkbn1/pPp1pppp/4P3/3pP3/3p4/4B1N1/PP1rBPPP/R3K2R w Kkq - 0 1"
	g := FromFen(fen)
	board := g.position.Board
	color := White
	ml := NewMoveList(3)
	taboo := tabooSquares(board, color)
	g.position.kingCaptureMoves(color, ml)
	g.position.kingQuietMoves(taboo, color, ml)
	moves := ml.Moves
	expectedMoves := []Move{
		NewMove(E1, D2, WhiteKing, BlackRook, NoType, Capture),
		NewMove(E1, F1, WhiteKing, NoPiece, NoType, 0),
		NewMove(E1, G1, WhiteKing, NoPiece, NoType, KingSideCastle),
	}
	expectedLen := len(expectedMoves)
	if len(moves) != expectedLen || !equalMoves(expectedMoves, moves) {
		fmt.Println("Got:")
		for _, i := range moves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(moves)))
	}
}

func TestKingCastlingWithOccupiedSquares(t *testing.T) {
	fen := "rnbqkbnr/1p6/p1p3Pp/1B1pp2Q/1P6/B7/P1PP1PPP/RN2K1NR w KQkq - 0 1"
	g := FromFen(fen)
	board := g.position.Board
	color := White
	taboo := tabooSquares(board, color)
	ml := NewMoveList(3)
	g.position.kingCaptureMoves(color, ml)
	g.position.kingQuietMoves(taboo, color, ml)
	moves := ml.Moves
	expectedMoves := []Move{
		NewMove(E1, E2, WhiteKing, NoPiece, NoType, 0),
		NewMove(E1, F1, WhiteKing, NoPiece, NoType, 0),
		NewMove(E1, D1, WhiteKing, NoPiece, NoType, 0),
	}
	expectedLen := len(expectedMoves)
	if len(moves) != expectedLen || !equalMoves(expectedMoves, moves) {
		fmt.Println("Got:")
		for _, i := range moves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(moves)))
	}
}

func TestKingQueenSideCastling(t *testing.T) {
	fen := "rnbqkbnr/1p6/p1p3Pp/1B1pp2Q/1P6/B7/P1PP1PPP/R3K1NR w KQkq - 0 1"
	g := FromFen(fen)
	board := g.position.Board
	color := White
	taboo := tabooSquares(board, color)
	ml := NewMoveList(4)
	g.position.kingCaptureMoves(color, ml)
	g.position.kingQuietMoves(taboo, color, ml)
	moves := ml.Moves
	expectedMoves := []Move{
		NewMove(E1, E2, WhiteKing, NoPiece, NoType, 0),
		NewMove(E1, F1, WhiteKing, NoPiece, NoType, 0),
		NewMove(E1, D1, WhiteKing, NoPiece, NoType, 0),
		NewMove(E1, C1, WhiteKing, NoPiece, NoType, QueenSideCastle),
	}
	expectedLen := len(expectedMoves)
	if len(moves) != expectedLen || !equalMoves(expectedMoves, moves) {
		fmt.Println("Got:")
		for _, i := range moves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(moves)))
	}
}

func TestPawnMovesForWhite(t *testing.T) {
	fen := "rnbqkbn1/pPp1pppp/4P3/3pP3/3p4/4B1N1/PP1rBPPP/R3K2R w Kkq d6 0 1"
	g := FromFen(fen)
	ml := NewMoveList(18)
	color := White
	g.position.pawnQuietMoves(color, ml)
	g.position.pawnCaptureMoves(color, ml)
	moves := ml.Moves
	expectedMoves := []Move{
		NewMove(H2, H4, WhitePawn, NoPiece, NoType, 0),
		NewMove(H2, H3, WhitePawn, NoPiece, NoType, 0),
		NewMove(F2, F4, WhitePawn, NoPiece, NoType, 0),
		NewMove(F2, F3, WhitePawn, NoPiece, NoType, 0),
		NewMove(A2, A4, WhitePawn, NoPiece, NoType, 0),
		NewMove(A2, A3, WhitePawn, NoPiece, NoType, 0),
		NewMove(B2, B4, WhitePawn, NoPiece, NoType, 0),
		NewMove(B2, B3, WhitePawn, NoPiece, NoType, 0),
		NewMove(E5, D6, WhitePawn, BlackPawn, NoType, EnPassant|Capture),
		NewMove(E6, F7, WhitePawn, BlackPawn, NoType, Capture),
		NewMove(B7, A8, WhitePawn, BlackRook, Queen, Capture),
		NewMove(B7, A8, WhitePawn, BlackRook, Rook, Capture),
		NewMove(B7, A8, WhitePawn, BlackRook, Bishop, Capture),
		NewMove(B7, A8, WhitePawn, BlackRook, Knight, Capture),
		NewMove(B7, C8, WhitePawn, BlackBishop, Queen, Capture),
		NewMove(B7, C8, WhitePawn, BlackBishop, Rook, Capture),
		NewMove(B7, C8, WhitePawn, BlackBishop, Bishop, Capture),
		NewMove(B7, C8, WhitePawn, BlackBishop, Knight, Capture),
	}
	expectedLen := len(expectedMoves)
	if len(moves) != expectedLen || !equalMoves(expectedMoves, moves) {
		fmt.Println("Got:")
		for _, i := range moves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(moves)))
	}
}

func TestPawnMovesForBlack(t *testing.T) {
	fen := "rnbqkbnr/ppp3pp/3p1p2/1P4P1/4pP2/N6N/P1PPP2P/R1BQKB1R b KQkq f3 0 1"
	g := FromFen(fen)
	ml := NewMoveList(13)
	color := Black
	g.position.pawnQuietMoves(color, ml)
	g.position.pawnCaptureMoves(color, ml)
	moves := ml.Moves
	expectedMoves := []Move{
		NewMove(H7, H6, BlackPawn, NoPiece, NoType, 0),
		NewMove(H7, H5, BlackPawn, NoPiece, NoType, 0),
		NewMove(G7, G6, BlackPawn, NoPiece, NoType, 0),
		NewMove(F6, F5, BlackPawn, NoPiece, NoType, 0),
		NewMove(F6, G5, BlackPawn, WhitePawn, NoType, Capture),
		NewMove(E4, E3, BlackPawn, NoPiece, NoType, 0),
		NewMove(E4, F3, BlackPawn, WhitePawn, NoType, EnPassant|Capture),
		NewMove(D6, D5, BlackPawn, NoPiece, NoType, 0),
		NewMove(C7, C6, BlackPawn, NoPiece, NoType, 0),
		NewMove(C7, C5, BlackPawn, NoPiece, NoType, 0),
		NewMove(B7, B6, BlackPawn, NoPiece, NoType, 0),
		NewMove(A7, A6, BlackPawn, NoPiece, NoType, 0),
		NewMove(A7, A5, BlackPawn, NoPiece, NoType, 0),
	}
	expectedLen := len(expectedMoves)
	if len(moves) != expectedLen || !equalMoves(expectedMoves, moves) {
		fmt.Println("Got:")
		for _, i := range moves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(moves)))
	}
}

func TestKnightMoves(t *testing.T) {
	fen := "rnbqkbn1/pPp1pppp/4P3/1N1pP3/3p4/4B1N1/PP1rBPPP/R3K2R w Kkq d6 0 1"
	g := FromFen(fen)
	ml := NewMoveList(10)
	g.position.knightQuietMoves(White, ml)
	g.position.knightCaptureMoves(White, ml)
	moves := ml.Moves
	expectedMoves := []Move{
		NewMove(G3, F1, WhiteKnight, NoPiece, NoType, 0),
		NewMove(G3, E4, WhiteKnight, NoPiece, NoType, 0),
		NewMove(G3, F5, WhiteKnight, NoPiece, NoType, 0),
		NewMove(G3, H5, WhiteKnight, NoPiece, NoType, 0),
		NewMove(B5, A7, WhiteKnight, BlackPawn, NoType, Capture),
		NewMove(B5, A3, WhiteKnight, NoPiece, NoType, 0),
		NewMove(B5, C7, WhiteKnight, BlackPawn, NoType, Capture),
		NewMove(B5, C3, WhiteKnight, NoPiece, NoType, 0),
		NewMove(B5, D4, WhiteKnight, BlackPawn, NoType, Capture),
		NewMove(B5, D6, WhiteKnight, NoPiece, NoType, 0),
	}
	expectedLen := len(expectedMoves)
	if len(moves) != expectedLen || !equalMoves(expectedMoves, moves) {
		fmt.Println("Got:")
		for _, i := range moves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(moves)))
	}
}

func TestCastleAndDiscoveredChecks(t *testing.T) {
	fen := "rnbq1bn1/pPp1pppp/4P3/3pP3/3p4/4B1N1/PP1rBPPP/k3K2R w Kkq - 0 1"
	g := FromFen(fen)
	p := g.position
	legalMoves := p.PseudoLegalMoves()
	move := NewMove(E1, G1, WhiteKing, NoPiece, NoType, KingSideCastle)
	if !containsMove(legalMoves, move) {
		fmt.Println("Got:")
		for _, i := range legalMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected to see %s", fmt.Sprintf("%s %d", move.ToString(), move.Tag()))
	}
	move = NewMove(E1, D2, WhiteKing, BlackRook, NoType, Capture)
	if !containsMove(legalMoves, move) {
		fmt.Println("Got:")
		for _, i := range legalMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected to see %s", fmt.Sprintf("%s %d", move.ToString(), move.Tag()))
	}
}

func TestCastleAndPawnAttack(t *testing.T) {
	fen := "r3k2r/p1ppqpb1/1n2pnp1/1b1PN3/1p2P3/P1N2Q2/1PPBBPpP/1R2K2R w Kkq - 0 1"
	g := FromFen(fen)
	board := g.position.Board
	ml := NewMoveList(1)
	color := White
	taboo := tabooSquares(board, color)
	g.position.kingCaptureMoves(color, ml)
	g.position.kingQuietMoves(taboo, color, ml)
	moves := ml.Moves
	expectedMoves := []Move{
		NewMove(E1, D1, WhiteKing, NoPiece, NoType, 0),
	}
	expectedLen := len(expectedMoves)
	if !equalMoves(expectedMoves, moves) {
		fmt.Println(g.position.Board.Draw())
		fmt.Println("Got:")
		for _, i := range moves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(moves)))
	}
}

func TestLegalMoves(t *testing.T) {
	fen := "rn1q1bn1/pPp1pppp/4P3/1N1pP2Q/3p3b/4B3/PP1rBPPP/k3K2R w Kkq d6 0 1"
	g := FromFen(fen)
	p := g.position
	legalMoves := p.PseudoLegalMoves()
	expectedMoves := []Move{
		NewMove(H1, G1, WhiteRook, NoPiece, NoType, 0),
		NewMove(H1, F1, WhiteRook, NoPiece, NoType, 0),
		NewMove(E1, F1, WhiteKing, NoPiece, NoType, 0),
		NewMove(E1, G1, WhiteKing, NoPiece, NoType, KingSideCastle),
		NewMove(E1, D2, WhiteKing, BlackRook, NoType, Capture),
		NewMove(H2, H3, WhitePawn, NoPiece, NoType, 0),
		NewMove(G2, G3, WhitePawn, NoPiece, NoType, 0),
		NewMove(G2, G4, WhitePawn, NoPiece, NoType, 0),
		NewMove(F2, F3, WhitePawn, NoPiece, NoType, 0),
		NewMove(F2, F4, WhitePawn, NoPiece, NoType, 0),
		NewMove(E2, F1, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, D1, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, F3, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, G4, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, D3, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E2, C4, WhiteBishop, NoPiece, NoType, 0),
		NewMove(B2, B3, WhitePawn, NoPiece, NoType, 0),
		NewMove(B2, B4, WhitePawn, NoPiece, NoType, 0),
		NewMove(A2, A3, WhitePawn, NoPiece, NoType, 0),
		NewMove(A2, A4, WhitePawn, NoPiece, NoType, 0),
		NewMove(E3, D4, WhiteBishop, BlackPawn, NoType, Capture),
		NewMove(E3, F4, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E3, G5, WhiteBishop, NoPiece, NoType, 0),
		NewMove(E3, H6, WhiteBishop, NoPiece, NoType, 0),
		NewMove(H5, H7, WhiteQueen, BlackPawn, NoType, Capture),
		NewMove(H5, H6, WhiteQueen, NoPiece, NoType, 0),
		NewMove(H5, F7, WhiteQueen, BlackPawn, NoType, Capture),
		NewMove(H5, G6, WhiteQueen, NoPiece, NoType, 0),
		NewMove(H5, G5, WhiteQueen, NoPiece, NoType, 0),
		NewMove(H5, F5, WhiteQueen, NoPiece, NoType, 0),
		NewMove(H5, G4, WhiteQueen, NoPiece, NoType, 0),
		NewMove(H5, F3, WhiteQueen, NoPiece, NoType, 0),
		NewMove(H5, H4, WhiteQueen, BlackBishop, NoType, Capture),
		NewMove(E5, D6, WhitePawn, BlackPawn, NoType, Capture|EnPassant),
		NewMove(B5, A3, WhiteKnight, NoPiece, NoType, 0),
		NewMove(B5, C3, WhiteKnight, NoPiece, NoType, 0),
		NewMove(B5, A7, WhiteKnight, BlackPawn, NoType, Capture),
		NewMove(B5, C7, WhiteKnight, BlackPawn, NoType, Capture),
		NewMove(B5, D4, WhiteKnight, BlackPawn, NoType, Capture),
		NewMove(B5, D6, WhiteKnight, NoPiece, NoType, 0),
		NewMove(B5, D6, WhiteKnight, NoPiece, NoType, 0),
		NewMove(E6, F7, WhitePawn, BlackPawn, NoType, Capture),
		NewMove(B7, A8, WhitePawn, BlackRook, Queen, Capture),
		NewMove(B7, A8, WhitePawn, BlackRook, Rook, Capture),
		NewMove(B7, A8, WhitePawn, BlackRook, Bishop, Capture),
		NewMove(B7, A8, WhitePawn, BlackRook, Knight, Capture),
	}
	expectedLen := len(expectedMoves)
	if expectedLen != len(legalMoves) || !equalMoves(expectedMoves, legalMoves) {
		fmt.Println("Got:")
		for _, i := range legalMoves {
			fmt.Println(i.ToString(), i.MovingPiece(), i.CapturedPiece(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.MovingPiece(), i.CapturedPiece(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(legalMoves)))
	}
}

func TestDoubleCheckResponses(t *testing.T) {
	fen := "5Q2/8/1q5P/8/6k1/5R2/6P1/2r3K1 w - - 0 1"
	g := FromFen(fen)
	p := g.position
	legalMoves := p.PseudoLegalMoves()
	expectedMoves := []Move{
		NewMove(G1, H2, WhiteKing, NoPiece, NoType, 0), // The only legal move
		NewMove(G1, H1, WhiteKing, NoPiece, NoType, 0),
		NewMove(G1, F1, WhiteKing, NoPiece, NoType, 0),
		NewMove(G1, F2, WhiteKing, NoPiece, NoType, 0),
		NewMove(H6, H7, WhitePawn, NoPiece, NoType, 0),
		NewMove(G2, G3, WhitePawn, NoPiece, NoType, 0),
		NewMove(F3, F1, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, F2, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, F4, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, F5, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, F6, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, F7, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, A3, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, B3, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, C3, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, D3, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, E3, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, G3, WhiteRook, NoPiece, NoType, 0),
		NewMove(F3, H3, WhiteRook, NoPiece, NoType, 0),
		NewMove(F8, G7, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, F7, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, F6, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, F5, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, F4, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, E7, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, D6, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, C5, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, B4, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, A3, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, A8, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, B8, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, C8, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, D8, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, E8, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, G8, WhiteQueen, NoPiece, NoType, 0),
		NewMove(F8, H8, WhiteQueen, NoPiece, NoType, 0),
	}
	if !p.IsInCheck() {
		t.Errorf("Position is wrongfully considered not check for: %s", fen)
	}

	expectedLen := len(expectedMoves)
	if expectedLen != len(legalMoves) || !equalMoves(expectedMoves, legalMoves) {
		fmt.Println("Got:")
		for _, i := range legalMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(legalMoves)))
	}
}

func TestLegalMovesInOpenning(t *testing.T) {
	fen := "rnbqkbnr/ppp3pp/3ppp2/1P6/6P1/N6N/P1PPPP1P/R1BQKB1R w KQkq - 0 1"
	g := FromFen(fen)
	p := g.position
	legalMoves := p.PseudoLegalMoves()
	expectedMoves := []Move{
		NewMove(H1, G1, WhiteRook, NoPiece, NoType, 0),
		NewMove(G4, G5, WhitePawn, NoPiece, NoType, 0),
		NewMove(F2, F3, WhitePawn, NoPiece, NoType, 0),
		NewMove(F2, F4, WhitePawn, NoPiece, NoType, 0),
		NewMove(E2, E3, WhitePawn, NoPiece, NoType, 0),
		NewMove(E2, E4, WhitePawn, NoPiece, NoType, 0),
		NewMove(D2, D3, WhitePawn, NoPiece, NoType, 0),
		NewMove(D2, D4, WhitePawn, NoPiece, NoType, 0),
		NewMove(C2, C3, WhitePawn, NoPiece, NoType, 0),
		NewMove(C2, C4, WhitePawn, NoPiece, NoType, 0),
		NewMove(B5, B6, WhitePawn, NoPiece, NoType, 0),
		NewMove(A1, B1, WhiteRook, NoPiece, NoType, 0),
		NewMove(A3, C4, WhiteKnight, NoPiece, NoType, 0),
		NewMove(A3, B1, WhiteKnight, NoPiece, NoType, 0),
		NewMove(C1, B2, WhiteBishop, NoPiece, NoType, 0),
		NewMove(F1, G2, WhiteBishop, NoPiece, NoType, 0),
		NewMove(H3, G5, WhiteKnight, NoPiece, NoType, 0),
		NewMove(H3, F4, WhiteKnight, NoPiece, NoType, 0),
		NewMove(H3, G1, WhiteKnight, NoPiece, NoType, 0),
	}
	expectedLen := len(expectedMoves)
	if expectedLen != len(legalMoves) || !equalMoves(expectedMoves, legalMoves) {
		fmt.Println("Got:")
		for _, i := range legalMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		fmt.Println("Expected:")
		for _, i := range expectedMoves {
			fmt.Println(i.ToString(), i.PromoType(), i.Tag())
		}
		t.Errorf("Expected different number of moves to be generated%s",
			fmt.Sprintf("\nExpected: %d\nGot: %d\n", expectedLen, len(legalMoves)))
	}
}

func equalMoves(moves1 []Move, moves2 []Move) bool {
	if len(moves1) != len(moves2) {
		return false
	}
	for _, m1 := range moves1 {
		exists := false
		for _, m2 := range moves2 {
			if m1 == m2 {
				exists = true
				break
			}
		}
		if !exists {
			fmt.Println("Missing", m1.ToString(), m1.Tag())
			return false
		}
	}
	return true
}

func containsMove(moves1 []Move, move Move) bool {
	exists := false
	for _, m := range moves1 {
		if m == move {
			exists = true
			break
		}
	}
	return exists
}
