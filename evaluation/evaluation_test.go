package evaluation

import (
	"fmt"
	"testing"

	. "github.com/amanjpro/zahak/engine"
)

func TestMaterialValue(t *testing.T) {
	fen := "rnb2bnr/ppqppkpp/8/2p5/4P3/8/PPPP1PPP/RNB1KBNR w KQ - 0 1"
	game := FromFen(fen)

	actual := Evaluate(game.Position(), NewPawnCache(DEFAULT_PAWNHASH_SIZE), NoColor, 0)

	if actual >= 0 {
		t.Errorf("Expected: a negative number\nGot: %d\n", actual)
	}

	fen = "3N1k2/N7/1p2ppR1/1P6/P2pP3/3Pb3/1K5n/8 w - - 0 4"

	game = FromFen(fen)

	actual = Evaluate(game.Position(), NewPawnCache(DEFAULT_PAWNHASH_SIZE), NoColor, 0)

	if actual <= 0 {
		t.Errorf("Expected: a positive number\nGot: %d\n", actual)
	}

	fen = "2k2b1r/ppp1pppp/4b3/1P6/2P3P1/3BKP1P/7B/1R4N1 b - - 0 23"
	game = FromFen(fen)

	actual = Evaluate(game.Position(), NewPawnCache(DEFAULT_PAWNHASH_SIZE), NoColor, 0)

	if actual >= 0 {
		t.Errorf("Expected: a negative number\nGot: %d\n", actual)
	}
}

func TestIsBackwardsPawn(t *testing.T) {
	game := FromFen("k7/5p2/4p1p1/8/8/4P1P1/5P2/K7 w - - 0 1")
	board := game.Position().Board

	actual := board.IsBackwardPawn(uint64(1<<int(E6)), board.GetBitboardOf(BlackPawn), Black)
	expected := false

	if actual != expected {
		t.Error("Non-Backward Pawn - Black: is falsely identified")
	}

	actual = board.IsBackwardPawn(uint64(1<<int(F7)), board.GetBitboardOf(BlackPawn), Black)
	expected = true

	if actual != expected {
		t.Error("Non-Backward Pawn - Black: is not identified")
	}

	actual = board.IsBackwardPawn(uint64(1<<int(G3)), board.GetBitboardOf(WhitePawn), White)
	expected = false

	if actual != expected {
		t.Error("Non-Backward Pawn - White: is falsely identified")
	}

	actual = board.IsBackwardPawn(uint64(1<<int(F2)), board.GetBitboardOf(WhitePawn), White)
	expected = true

	if actual != expected {
		t.Error("Non-Backward Pawn - White: is not identified")
	}
}

func TestPawnStructureEval(t *testing.T) {
	fen := "k7/4pp2/6p1/8/8/4P1P1/5P2/K7 w - - 0 1"
	game := FromFen(fen)

	actual := Evaluate(game.Position(), NewPawnCache(DEFAULT_PAWNHASH_SIZE), NoColor, 0)
	expected := int16(-9)

	if actual != expected {
		err := fmt.Sprintf("Backward Pawn - White:\nExpected: %d\nGot: %d\n", expected, actual)
		t.Errorf(err)
	}

	fen = "k7/5p2/4p1p1/8/8/6P1/4PP2/K7 b - - 0 1"
	game = FromFen(fen)

	actual = Evaluate(game.Position(), NewPawnCache(DEFAULT_PAWNHASH_SIZE), NoColor, 0)
	expected = int16(-9)

	if actual != expected {
		err := fmt.Sprintf("Backward Pawn - Black:\nExpected: %d\nGot: %d\n", expected, actual)
		t.Errorf(err)
	}
}

func TestRookStructureEval(t *testing.T) {
	fen := "k4r2/5p2/8/8/8/8/4P3/K4R2 w - - 0 1"
	game := FromFen(fen)

	actual := Evaluate(game.Position(), NewPawnCache(DEFAULT_PAWNHASH_SIZE), NoColor, 0)
	expected := int16(50)

	if actual != expected {
		err := fmt.Sprintf("Semi-open file - White:\nExpected: %d\nGot: %d\n", expected, actual)
		t.Errorf(err)
	}

	fen = "k4r2/4p3/8/8/8/8/5P2/K4R2 b - - 0 1"
	game = FromFen(fen)

	actual = Evaluate(game.Position(), NewPawnCache(DEFAULT_PAWNHASH_SIZE), NoColor, 0)
	expected = int16(50)

	if actual != expected {
		err := fmt.Sprintf("Semi-open file - Black:\nExpected: %d\nGot: %d\n", expected, actual)
		t.Errorf(err)
	}
}

func TestRookFilesEval(t *testing.T) {
	fen := "kr1rr2r/7p/8/8/1P6/8/8/KR1R3R w - - 0 1"
	game := FromFen(fen)

	whiteRook := game.Position().Board.GetBitboardOf(WhiteRook)
	blackRook := game.Position().Board.GetBitboardOf(BlackRook)
	whitePawn := game.Position().Board.GetBitboardOf(WhitePawn)
	blackPawn := game.Position().Board.GetBitboardOf(BlackPawn)

	actual := RookFilesEval(blackRook, whiteRook, blackPawn, whitePawn)
	expected := Eval{
		blackMG: MiddlegameRookOpenFileAward*2 + MiddlegameRookSemiOpenFileAward*1,
		whiteMG: MiddlegameRookOpenFileAward*1 + MiddlegameRookSemiOpenFileAward*1,
		blackEG: EndgameRookOpenFileAward*2 + EndgameRookSemiOpenFileAward*1,
		whiteEG: EndgameRookOpenFileAward*1 + EndgameRookSemiOpenFileAward*1,
	}

	if actual != expected {
		err := fmt.Sprintf("Expected: %d\nGot: %d\n", expected, actual)
		t.Errorf(err)
	}
}

func TestKingSafetyWhiteOnG(t *testing.T) {
	fen := "1k6/1pp3p1/8/8/5P2/6P1/P7/6K1 w - - 0 1"
	game := FromFen(fen)

	whiteKing := game.Position().Board.GetBitboardOf(WhiteKing)
	blackKing := game.Position().Board.GetBitboardOf(BlackKing)
	whitePawn := game.Position().Board.GetBitboardOf(WhitePawn)
	blackPawn := game.Position().Board.GetBitboardOf(BlackPawn)

	actual := KingSafety(blackKing, whiteKing, blackPawn, whitePawn, false, false)
	expected := Eval{
		blackMG: -(MiddlegameKingZoneMissingPawnPenalty * 1),
		blackEG: -(EndgameKingZoneMissingPawnPenalty * 1),
		whiteMG: -(MiddlegamePawnShieldPenalty*1 + 1*MiddlegameKingZoneOpenFilePenalty),
		whiteEG: -(EndgamePawnShieldPenalty*1 + 1*EndgameKingZoneOpenFilePenalty),
	}

	if actual != expected {
		err := fmt.Sprintf("Expected: %d\nGot: %d\n", expected, actual)
		t.Errorf(err)
	}
}

func TestKingSafetyBlackOnG(t *testing.T) {
	fen := "6k1/p7/6p1/5p2/8/8/1PP3P1/1K6 w - - 0 1"
	game := FromFen(fen)

	whiteKing := game.Position().Board.GetBitboardOf(WhiteKing)
	blackKing := game.Position().Board.GetBitboardOf(BlackKing)
	whitePawn := game.Position().Board.GetBitboardOf(WhitePawn)
	blackPawn := game.Position().Board.GetBitboardOf(BlackPawn)

	actual := KingSafety(blackKing, whiteKing, blackPawn, whitePawn, false, false)
	expected := Eval{
		whiteMG: -(MiddlegameKingZoneMissingPawnPenalty * 1),
		whiteEG: -(EndgameKingZoneMissingPawnPenalty * 1),
		blackMG: -(MiddlegamePawnShieldPenalty*1 + 1*MiddlegameKingZoneOpenFilePenalty),
		blackEG: -(EndgamePawnShieldPenalty*1 + 1*EndgameKingZoneOpenFilePenalty),
	}

	if actual != expected {
		err := fmt.Sprintf("Expected: %d\nGot: %d\n", expected, actual)
		t.Errorf(err)
	}
}

func TestKingSafetySemiOpenFile(t *testing.T) {
	fen := "6k1/ppp3pp/8/8/8/8/P4PPP/1K6 w - - 0 1"
	game := FromFen(fen)

	whiteKing := game.Position().Board.GetBitboardOf(WhiteKing)
	blackKing := game.Position().Board.GetBitboardOf(BlackKing)
	whitePawn := game.Position().Board.GetBitboardOf(WhitePawn)
	blackPawn := game.Position().Board.GetBitboardOf(BlackPawn)

	actual := KingSafety(blackKing, whiteKing, blackPawn, whitePawn, false, false)
	expected := Eval{
		whiteMG: -(MiddlegameKingZoneMissingPawnPenalty * 2),
		whiteEG: -(EndgameKingZoneMissingPawnPenalty * 2),
		blackMG: -(MiddlegameKingZoneMissingPawnPenalty * 1),
		blackEG: -(EndgameKingZoneMissingPawnPenalty * 1),
	}

	if actual != expected {
		err := fmt.Sprintf("Expected: %d\nGot: %d\n", expected, actual)
		t.Errorf(err)
	}
}

func TestKingSafetyBlackNotCastling(t *testing.T) {
	fen := "rnbq1bnr/pppppppp/3k4/8/8/3K4/PPPPPPPP/RNBQ1BNR w - - 0 1"
	game := FromFen(fen)

	whiteKing := game.Position().Board.GetBitboardOf(WhiteKing)
	blackKing := game.Position().Board.GetBitboardOf(BlackKing)
	whitePawn := game.Position().Board.GetBitboardOf(WhitePawn)
	blackPawn := game.Position().Board.GetBitboardOf(BlackPawn)

	actual := KingSafety(blackKing, whiteKing, blackPawn, whitePawn, false, false)
	expected := Eval{
		whiteMG: -MiddlegameNotCastlingPenalty,
		whiteEG: -EndgameNotCastlingPenalty,
		blackMG: -MiddlegameNotCastlingPenalty,
		blackEG: -EndgameNotCastlingPenalty,
	}

	if actual != expected {
		err := fmt.Sprintf("Expected: %d\nGot: %d\n", expected, actual)
		t.Errorf(err)
	}
}
