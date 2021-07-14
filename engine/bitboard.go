package engine

import (
	"fmt"
)

type Bitboard struct {
	blackPawn   uint64
	blackKnight uint64
	blackBishop uint64
	blackRook   uint64
	blackQueen  uint64
	blackKing   uint64
	whitePawn   uint64
	whiteKnight uint64
	whiteBishop uint64
	whiteRook   uint64
	whiteQueen  uint64
	whiteKing   uint64
	whitePieces uint64
	blackPieces uint64
	PieceMap    [64]Piece
}

func (b *Bitboard) GetWhitePieces() uint64 {
	return b.whitePieces
}

func (b *Bitboard) GetBlackPieces() uint64 {
	return b.blackPieces
}

func (b *Bitboard) GetBitboardOf(piece Piece) uint64 {
	switch piece {
	case BlackPawn:
		return b.blackPawn
	case BlackKnight:
		return b.blackKnight
	case BlackBishop:
		return b.blackBishop
	case BlackRook:
		return b.blackRook
	case BlackQueen:
		return b.blackQueen
	case BlackKing:
		return b.blackKing
	case WhitePawn:
		return b.whitePawn
	case WhiteKnight:
		return b.whiteKnight
	case WhiteBishop:
		return b.whiteBishop
	case WhiteRook:
		return b.whiteRook
	case WhiteQueen:
		return b.whiteQueen
	case WhiteKing:
		return b.whiteKing
	}
	return 0
}

func (b *Bitboard) AllPieces() map[Square]Piece {
	allPieces := make(map[Square]Piece, 32)
	for i, p := range b.PieceMap {
		if p != NoPiece {
			allPieces[Square(i)] = p
		}
	}
	return allPieces
}

func (b *Bitboard) UpdateSquare(sq Square, newPiece Piece, oldPiece Piece) {
	// Remove the piece from source square and add it to destination
	b.Clear(sq, oldPiece)
	mask := SquareMask[int(sq)]
	b.PieceMap[sq] = newPiece
	switch newPiece {
	case BlackPawn:
		b.blackPawn |= mask
		b.blackPieces |= mask
	case BlackKnight:
		b.blackKnight |= mask
		b.blackPieces |= mask
	case BlackBishop:
		b.blackBishop |= mask
		b.blackPieces |= mask
	case BlackRook:
		b.blackRook |= mask
		b.blackPieces |= mask
	case BlackQueen:
		b.blackQueen |= mask
		b.blackPieces |= mask
	case BlackKing:
		b.blackKing |= mask
		b.blackPieces |= mask
	case WhitePawn:
		b.whitePawn |= mask
		b.whitePieces |= mask
	case WhiteKnight:
		b.whiteKnight |= mask
		b.whitePieces |= mask
	case WhiteBishop:
		b.whiteBishop |= mask
		b.whitePieces |= mask
	case WhiteRook:
		b.whiteRook |= mask
		b.whitePieces |= mask
	case WhiteQueen:
		b.whiteQueen |= mask
		b.whitePieces |= mask
	case WhiteKing:
		b.whiteKing |= mask
		b.whitePieces |= mask
	}
}

func (b *Bitboard) PieceAt(sq Square) Piece {
	if sq == NoSquare {
		return NoPiece
	}
	return b.PieceMap[sq]
}

func (b *Bitboard) Clear(square Square, piece Piece) {
	if piece == NoPiece {
		return
	}
	mask := SquareMask[int(square)]
	if b.PieceMap[square] == piece {
		b.PieceMap[square] = NoPiece
	}
	switch piece {
	case BlackPawn:
		b.blackPawn &^= mask
		b.blackPieces &^= mask
	case BlackKnight:
		b.blackKnight &^= mask
		b.blackPieces &^= mask
	case BlackBishop:
		b.blackBishop &^= mask
		b.blackPieces &^= mask
	case BlackRook:
		b.blackRook &^= mask
		b.blackPieces &^= mask
	case BlackQueen:
		b.blackQueen &^= mask
		b.blackPieces &^= mask
	case BlackKing:
		b.blackKing &^= mask
		b.blackPieces &^= mask
	case WhitePawn:
		b.whitePawn &^= mask
		b.whitePieces &^= mask
	case WhiteKnight:
		b.whiteKnight &^= mask
		b.whitePieces &^= mask
	case WhiteBishop:
		b.whiteBishop &^= mask
		b.whitePieces &^= mask
	case WhiteRook:
		b.whiteRook &^= mask
		b.whitePieces &^= mask
	case WhiteQueen:
		b.whiteQueen &^= mask
		b.whitePieces &^= mask
	case WhiteKing:
		b.whiteKing &^= mask
		b.whitePieces &^= mask
	}
}

func (b *Bitboard) Move(src Square, dest Square, sourcePiece Piece, destinationPiece Piece) {

	if src == NoSquare || dest == NoSquare {
		return
	}
	// clear destination square
	b.Clear(dest, destinationPiece)
	b.Clear(src, sourcePiece)
	maskDest := SquareMask[int(dest)]
	b.PieceMap[dest] = sourcePiece

	// Remove the piece from source square and add it to destination
	switch sourcePiece {
	case BlackPawn:
		b.blackPawn |= maskDest
		b.blackPieces |= maskDest
	case BlackKnight:
		b.blackKnight |= maskDest
		b.blackPieces |= maskDest
	case BlackBishop:
		b.blackBishop |= maskDest
		b.blackPieces |= maskDest
	case BlackRook:
		b.blackRook |= maskDest
		b.blackPieces |= maskDest
	case BlackQueen:
		b.blackQueen |= maskDest
		b.blackPieces |= maskDest
	case BlackKing:
		b.blackKing |= maskDest
		b.blackPieces |= maskDest
		// Is it a castle?
		if src == E8 && dest == G8 {
			b.Move(H8, F8, BlackRook, NoPiece)
		} else if src == E8 && dest == C8 {
			b.Move(A8, D8, BlackRook, NoPiece)
		}
	case WhitePawn:
		b.whitePawn |= maskDest
		b.whitePieces |= maskDest
	case WhiteKnight:
		b.whiteKnight |= maskDest
		b.whitePieces |= maskDest
	case WhiteBishop:
		b.whiteBishop |= maskDest
		b.whitePieces |= maskDest
	case WhiteRook:
		b.whiteRook |= maskDest
		b.whitePieces |= maskDest
	case WhiteQueen:
		b.whiteQueen |= maskDest
		b.whitePieces |= maskDest
	case WhiteKing:
		b.whiteKing |= maskDest
		b.whitePieces |= maskDest
		// Is it a castle?
		if src == E1 && dest == G1 {
			b.Move(H1, F1, WhiteRook, NoPiece)
		} else if src == E1 && dest == C1 {
			b.Move(A1, D1, WhiteRook, NoPiece)
		}
	}
}

func EmptyBitboard() *Bitboard {
	var pieces [64]Piece
	for i := 0; i < 64; i++ {
		pieces[i] = NoPiece
	}
	bitboard := &Bitboard{PieceMap: pieces}
	return bitboard
}

func StartingBoard() *Bitboard {
	bitboard := EmptyBitboard()
	bitboard.UpdateSquare(A2, WhitePawn, NoPiece)
	bitboard.UpdateSquare(B2, WhitePawn, NoPiece)
	bitboard.UpdateSquare(C2, WhitePawn, NoPiece)
	bitboard.UpdateSquare(D2, WhitePawn, NoPiece)
	bitboard.UpdateSquare(E2, WhitePawn, NoPiece)
	bitboard.UpdateSquare(F2, WhitePawn, NoPiece)
	bitboard.UpdateSquare(G2, WhitePawn, NoPiece)
	bitboard.UpdateSquare(H2, WhitePawn, NoPiece)

	bitboard.UpdateSquare(A7, BlackPawn, NoPiece)
	bitboard.UpdateSquare(B7, BlackPawn, NoPiece)
	bitboard.UpdateSquare(C7, BlackPawn, NoPiece)
	bitboard.UpdateSquare(D7, BlackPawn, NoPiece)
	bitboard.UpdateSquare(E7, BlackPawn, NoPiece)
	bitboard.UpdateSquare(F7, BlackPawn, NoPiece)
	bitboard.UpdateSquare(G7, BlackPawn, NoPiece)
	bitboard.UpdateSquare(H7, BlackPawn, NoPiece)

	bitboard.UpdateSquare(A1, WhiteRook, NoPiece)
	bitboard.UpdateSquare(B1, WhiteKnight, NoPiece)
	bitboard.UpdateSquare(C1, WhiteBishop, NoPiece)
	bitboard.UpdateSquare(D1, WhiteQueen, NoPiece)
	bitboard.UpdateSquare(E1, WhiteKing, NoPiece)
	bitboard.UpdateSquare(F1, WhiteBishop, NoPiece)
	bitboard.UpdateSquare(G1, WhiteKnight, NoPiece)
	bitboard.UpdateSquare(H1, WhiteRook, NoPiece)

	bitboard.UpdateSquare(A8, BlackRook, NoPiece)
	bitboard.UpdateSquare(B8, BlackKnight, NoPiece)
	bitboard.UpdateSquare(C8, BlackBishop, NoPiece)
	bitboard.UpdateSquare(D8, BlackQueen, NoPiece)
	bitboard.UpdateSquare(E8, BlackKing, NoPiece)
	bitboard.UpdateSquare(F8, BlackBishop, NoPiece)
	bitboard.UpdateSquare(G8, BlackKnight, NoPiece)
	bitboard.UpdateSquare(H8, BlackRook, NoPiece)

	return bitboard
}

func (b *Bitboard) IsEndGame(turn Color) bool {
	if turn == White {
		return b.whiteKnight+b.whiteBishop+b.whiteRook+b.whiteQueen == 0
	} else if turn == Black {
		return b.blackKnight+b.blackBishop+b.blackRook+b.blackQueen == 0
	}
	return false
}

// Draw returns visual representation of the board useful for debugging.
func (b *Bitboard) Draw() string {
	pieceUnicodes := []string{"♙", "♘", "♗", "♖", "♕", "♔", "♟", "♞", "♝", "♜", "♛", "♚"}
	s := "\n A B C D E F G H\n"
	for r := 7; r >= 0; r-- {
		s += fmt.Sprint(Rank(r + 1))
		for f := 0; f < len(Files); f++ {
			p := b.PieceAt(SquareOf(File(f), Rank(r)))
			if p == NoPiece {
				s += "-"
			} else {
				s += pieceUnicodes[int(p-1)]
			}
			s += " "
		}
		s += "\n"
	}
	return s
}

func (b *Bitboard) copy() *Bitboard {
	var other [64]Piece
	for i, p := range b.PieceMap {
		other[i] = p
	}
	return &Bitboard{
		b.blackPawn,
		b.blackKnight,
		b.blackBishop,
		b.blackRook,
		b.blackQueen,
		b.blackKing,
		b.whitePawn,
		b.whiteKnight,
		b.whiteBishop,
		b.whiteRook,
		b.whiteQueen,
		b.whiteKing,
		b.whitePieces,
		b.blackPieces,
		other,
	}
}
