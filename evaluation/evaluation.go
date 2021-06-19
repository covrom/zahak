package evaluation

import (
	"math/bits"

	. "github.com/amanjpro/zahak/engine"
)

type Eval struct {
	blackMG int16
	whiteMG int16
	blackEG int16
	whiteEG int16
}

const CHECKMATE_EVAL int16 = 30000
const MAX_NON_CHECKMATE int16 = 25000
const PawnPhase int16 = 0
const KnightPhase int16 = 1
const BishopPhase int16 = 1
const RookPhase int16 = 2
const QueenPhase int16 = 4
const TotalPhase int16 = PawnPhase*16 + KnightPhase*4 + BishopPhase*4 + RookPhase*4 + QueenPhase*2
const HalfPhase = TotalPhase / 2
const Tempo int16 = 5

const BlackKingSideMask = uint64(1<<F8 | 1<<G8 | 1<<H8 | 1<<F7 | 1<<G7 | 1<<H7)
const WhiteKingSideMask = uint64(1<<F1 | 1<<G1 | 1<<H1 | 1<<F2 | 1<<G2 | 1<<H2)
const BlackQueenSideMask = uint64(1<<C8 | 1<<B8 | 1<<A8 | 1<<A7 | 1<<B7 | 1<<C7)
const WhiteQueenSideMask = uint64(1<<C1 | 1<<B1 | 1<<A1 | 1<<A2 | 1<<B2 | 1<<C2)

const BlackAShield = uint64(1<<A7 | 1<<A6)
const BlackBShield = uint64(1<<B7 | 1<<B6)
const BlackCShield = uint64(1<<C7 | 1<<C6)
const BlackFShield = uint64(1<<F7 | 1<<F6)
const BlackGShield = uint64(1<<G7 | 1<<G6)
const BlackHShield = uint64(1<<H7 | 1<<H6)
const WhiteAShield = uint64(1<<A2 | 1<<A3)
const WhiteBShield = uint64(1<<B2 | 1<<B3)
const WhiteCShield = uint64(1<<C2 | 1<<C3)
const WhiteFShield = uint64(1<<F2 | 1<<F3)
const WhiteGShield = uint64(1<<G2 | 1<<G3)
const WhiteHShield = uint64(1<<H2 | 1<<H3)

// Piece Square Tables
// Middle-game
var EarlyPawnPst = [64]int16{
	0, 0, 0, 0, 0, 0, 0, 0,
	92, 132, 66, 108, 92, 126, 12, -31,
	-11, -12, 21, 20, 61, 74, 15, -18,
	-24, -7, -4, 15, 15, 12, 1, -27,
	-35, -24, -13, 5, 9, 3, -9, -33,
	-35, -30, -20, -22, -10, -10, 6, -25,
	-43, -25, -33, -34, -31, 9, 11, -33,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var EarlyKnightPst = [64]int16{
	-208, -82, -49, -46, 55, -122, -24, -138,
	-79, -39, 81, 38, 30, 75, 1, -11,
	-36, 72, 47, 65, 94, 138, 69, 55,
	10, 34, 25, 56, 26, 76, 18, 34,
	11, 34, 36, 26, 44, 27, 37, 11,
	1, 15, 30, 37, 49, 36, 48, 11,
	-2, -19, 16, 26, 25, 41, 21, 17,
	-97, 9, -26, -9, 21, 7, 14, 7,
}

var EarlyBishopPst = [64]int16{
	-24, 26, -83, -45, -21, -36, 11, 14,
	6, 50, 14, 0, 54, 76, 43, -24,
	23, 68, 82, 66, 59, 78, 51, 27,
	34, 41, 41, 71, 59, 53, 35, 24,
	38, 54, 49, 60, 69, 49, 48, 40,
	35, 58, 56, 53, 59, 77, 58, 45,
	46, 66, 60, 48, 58, 66, 84, 48,
	7, 39, 38, 30, 40, 36, 3, 19,
}

var EarlyRookPst = [64]int16{
	-3, 13, -21, 22, 21, -19, 0, -9,
	0, -2, 31, 32, 57, 58, -7, 19,
	-38, -10, -7, -8, -31, 22, 41, -20,
	-46, -33, -16, -2, -17, 12, -19, -40,
	-53, -46, -29, -27, -13, -27, -1, -39,
	-53, -35, -32, -32, -19, -8, -15, -33,
	-49, -27, -34, -24, -16, 5, -14, -73,
	-20, -20, -13, -2, -1, -2, -38, -17,
}

var EarlyQueenPst = [64]int16{
	-63, -31, -15, -21, 37, 35, 30, 11,
	-33, -62, -24, -27, -69, 26, -7, 18,
	-18, -23, -15, -44, -9, 25, 0, 18,
	-37, -36, -38, -48, -32, -22, -35, -27,
	-8, -39, -23, -24, -23, -18, -17, -13,
	-25, 9, -9, -1, -4, 1, 7, -1,
	-24, -1, 20, 12, 18, 26, 7, 18,
	16, -3, 10, 25, 0, -7, -10, -27,
}

var EarlyKingPst = [64]int16{
	-50, 97, 97, 51, -47, -13, 42, 47,
	101, 32, 25, 64, 21, 13, -21, -63,
	37, 43, 55, 9, 26, 62, 59, -15,
	-20, -5, 10, -23, -18, -20, -24, -57,
	-46, 11, -35, -72, -73, -51, -61, -79,
	-5, -16, -27, -58, -59, -46, -21, -37,
	15, 18, -10, -58, -38, -15, 10, 18,
	-9, 36, 14, -64, -9, -35, 28, 26,
}

// Endgame
var LatePawnPst = [64]int16{
	0, 0, 0, 0, 0, 0, 0, 0,
	168, 148, 132, 101, 115, 103, 154, 191,
	85, 81, 59, 31, 16, 22, 58, 71,
	22, 4, -9, -28, -19, -13, 0, 9,
	21, 9, -1, -9, -9, -8, -2, 5,
	5, 0, -9, -5, -4, -8, -16, -11,
	16, -2, 5, 2, 10, -9, -15, -9,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var LateKnightPst = [64]int16{
	-46, -50, -16, -39, -44, -34, -73, -98,
	-31, -14, -48, -18, -28, -52, -34, -62,
	-36, -40, -13, -13, -30, -39, -38, -61,
	-24, -10, 6, 7, 13, -8, -2, -28,
	-29, -24, -1, 13, 2, 4, -5, -25,
	-33, -14, -18, 0, -8, -18, -32, -33,
	-44, -28, -19, -15, -12, -30, -30, -53,
	-22, -56, -29, -15, -28, -28, -58, -74,
}

var LateBishopPst = [64]int16{
	-9, -26, -5, -9, -5, -9, -15, -28,
	-11, -17, -4, -16, -16, -24, -16, -15,
	-4, -19, -18, -19, -16, -16, -9, -3,
	-7, 0, 2, -5, -1, -4, -8, 0,
	-14, -9, 1, 4, -11, -2, -15, -11,
	-12, -10, 0, -1, 3, -13, -10, -15,
	-16, -24, -13, -5, -3, -17, -21, -32,
	-18, -8, -18, -3, -9, -13, -3, -11,
}

var LateRookPst = [64]int16{
	7, 1, 13, 2, 4, 12, 6, 7,
	8, 11, 0, 1, -17, -7, 12, 2,
	11, 8, 2, 3, 3, -8, -9, 1,
	13, 8, 14, 2, 4, 6, 2, 14,
	16, 19, 16, 12, 4, 8, -2, 5,
	12, 12, 8, 12, 4, -1, 5, -2,
	9, 4, 13, 12, 1, -4, -4, 16,
	5, 11, 9, 0, -2, -3, 10, -14,
}

var LateQueenPst = [64]int16{
	31, 57, 52, 52, 40, 36, 26, 58,
	5, 47, 49, 63, 91, 42, 55, 39,
	-1, 22, 18, 85, 67, 49, 53, 37,
	44, 43, 44, 69, 77, 59, 92, 68,
	0, 58, 43, 68, 53, 55, 72, 53,
	28, -22, 32, 20, 31, 38, 47, 43,
	-3, -7, -19, 0, 5, 0, -14, -9,
	-19, -16, -5, -29, 21, -12, 2, -26,
}

var LateKingPst = [64]int16{
	-72, -51, -33, -29, -4, 17, -3, -14,
	-32, 0, 1, -3, 4, 28, 15, 21,
	2, 3, 4, 4, 4, 27, 26, 14,
	-7, 11, 14, 22, 18, 26, 21, 10,
	-14, -15, 19, 27, 29, 22, 9, 0,
	-19, -8, 8, 22, 24, 17, 3, -3,
	-28, -19, 5, 11, 12, 4, -11, -20,
	-53, -47, -23, 3, -26, -4, -37, -55,
}

var MiddlegameBackwardPawnPenalty int16 = 10
var EndgameBackwardPawnPenalty int16 = 3
var MiddlegameIsolatedPawnPenalty int16 = 10
var EndgameIsolatedPawnPenalty int16 = 4
var MiddlegameDoublePawnPenalty int16 = 2
var EndgameDoublePawnPenalty int16 = 27
var MiddlegamePassedPawnAward int16 = 2
var EndgamePassedPawnAward int16 = 12
var MiddlegameAdvancedPassedPawnAward int16 = 12
var EndgameAdvancedPassedPawnAward int16 = 64
var MiddlegameCandidatePassedPawnAward int16 = 32
var EndgameCandidatePassedPawnAward int16 = 50
var MiddlegameRookOpenFileAward int16 = 46
var EndgameRookOpenFileAward int16 = 0
var MiddlegameRookSemiOpenFileAward int16 = 15
var EndgameRookSemiOpenFileAward int16 = 19
var MiddlegameVeritcalDoubleRookAward int16 = 10
var EndgameVeritcalDoubleRookAward int16 = 10
var MiddlegameHorizontalDoubleRookAward int16 = 25
var EndgameHorizontalDoubleRookAward int16 = 12
var MiddlegamePawnFactorCoeff int16 = 0
var EndgamePawnFactorCoeff int16 = 0
var MiddlegameMobilityFactorCoeff int16 = 6
var EndgameMobilityFactorCoeff int16 = 3
var MiddlegameAggressivityFactorCoeff int16 = 1
var EndgameAggressivityFactorCoeff int16 = 6
var MiddlegameInnerPawnToKingAttackCoeff int16 = 0
var EndgameInnerPawnToKingAttackCoeff int16 = 0
var MiddlegameOuterPawnToKingAttackCoeff int16 = 4
var EndgameOuterPawnToKingAttackCoeff int16 = 0
var MiddlegameInnerMinorToKingAttackCoeff int16 = 18
var EndgameInnerMinorToKingAttackCoeff int16 = 0
var MiddlegameOuterMinorToKingAttackCoeff int16 = 11
var EndgameOuterMinorToKingAttackCoeff int16 = 1
var MiddlegameInnerMajorToKingAttackCoeff int16 = 17
var EndgameInnerMajorToKingAttackCoeff int16 = 0
var MiddlegameOuterMajorToKingAttackCoeff int16 = 8
var EndgameOuterMajorToKingAttackCoeff int16 = 5
var MiddlegamePawnShieldPenalty int16 = 10
var EndgamePawnShieldPenalty int16 = 9
var MiddlegameNotCastlingPenalty int16 = 26
var EndgameNotCastlingPenalty int16 = 5
var MiddlegameKingZoneOpenFilePenalty int16 = 36
var EndgameKingZoneOpenFilePenalty int16 = 0
var MiddlegameKingZoneMissingPawnPenalty int16 = 15
var EndgameKingZoneMissingPawnPenalty int16 = 0
var MiddlegamePawnIslandPenalty int16 = 3
var EndgamePawnIslandPenalty int16 = 7

// // Middle-game
// var EarlyPawnPst = [64]int16{
// 	0, 0, 0, 0, 0, 0, 0, 0,
// 	96, 131, 66, 105, 97, 126, 15, -31,
// 	-14, -13, 21, 21, 62, 74, 16, -18,
// 	-24, -6, -3, 16, 15, 12, 1, -29,
// 	-35, -23, -13, 5, 10, 3, -9, -34,
// 	-35, -30, -20, -22, -10, -10, 6, -25,
// 	-43, -25, -33, -35, -30, 9, 11, -33,
// 	0, 0, 0, 0, 0, 0, 0, 0,
// }
//
// var EarlyKnightPst = [64]int16{
// 	-211, -81, -50, -44, 54, -119, -26, -140,
// 	-79, -39, 79, 39, 30, 75, 4, -12,
// 	-39, 71, 48, 67, 95, 139, 69, 55,
// 	9, 32, 25, 57, 27, 74, 17, 36,
// 	8, 34, 36, 25, 46, 27, 34, 9,
// 	-2, 13, 29, 34, 49, 36, 48, 9,
// 	-4, -23, 14, 24, 24, 40, 19, 17,
// 	-99, 7, -29, -14, 18, 4, 13, 5,
// }
//
// var EarlyBishopPst = [64]int16{
// 	-25, 24, -83, -45, -22, -36, 11, 12,
// 	5, 50, 14, 0, 57, 77, 41, -25,
// 	20, 69, 81, 66, 63, 78, 52, 25,
// 	32, 39, 42, 72, 59, 53, 35, 24,
// 	36, 52, 48, 60, 68, 48, 46, 40,
// 	33, 58, 55, 53, 57, 76, 58, 43,
// 	46, 65, 59, 47, 57, 66, 83, 47,
// 	6, 41, 37, 29, 40, 34, 3, 16,
// }
//
// var EarlyRookPst = [64]int16{
// 	-3, 14, -19, 21, 19, -19, 0, -9,
// 	3, 0, 33, 31, 59, 59, -2, 20,
// 	-36, -10, -5, -5, -29, 24, 43, -17,
// 	-43, -32, -15, 2, -16, 14, -19, -38,
// 	-51, -46, -28, -26, -13, -25, 0, -38,
// 	-52, -35, -32, -33, -18, -7, -15, -34,
// 	-50, -26, -34, -24, -15, 3, -15, -73,
// 	-19, -20, -12, -2, 0, -1, -38, -18,
// }
//
// var EarlyQueenPst = [64]int16{
// 	-64, -33, -12, -17, 36, 33, 32, 12,
// 	-35, -62, -24, -24, -65, 27, -5, 20,
// 	-18, -24, -14, -42, -7, 28, 1, 19,
// 	-39, -38, -38, -49, -33, -21, -34, -27,
// 	-10, -41, -24, -24, -24, -18, -17, -14,
// 	-25, 7, -10, -3, -5, 0, 6, -3,
// 	-25, -2, 19, 11, 17, 24, 7, 16,
// 	13, -4, 9, 25, -2, -8, -10, -31,
// }
//
// var EarlyKingPst = [64]int16{
// 	-52, 95, 93, 48, -45, -13, 44, 46,
// 	97, 33, 25, 63, 19, 4, -21, -61,
// 	35, 38, 53, 8, 26, 60, 56, -14,
// 	-21, -8, 10, -21, -19, -20, -22, -56,
// 	-45, 9, -36, -72, -72, -50, -60, -78,
// 	-6, -15, -27, -58, -59, -46, -21, -37,
// 	14, 17, -10, -61, -40, -15, 10, 17,
// 	-9, 37, 15, -66, -8, -35, 28, 26,
// }
//
// // Endgame
// var LatePawnPst = [64]int16{
// 	0, 0, 0, 0, 0, 0, 0, 0,
// 	169, 150, 134, 103, 116, 104, 155, 191,
// 	87, 85, 61, 34, 18, 24, 61, 73,
// 	23, 4, -7, -28, -19, -12, 1, 9,
// 	20, 9, -1, -9, -9, -7, -2, 4,
// 	4, 0, -8, -5, -4, -7, -14, -12,
// 	15, -1, 5, 3, 11, -8, -13, -10,
// 	0, 0, 0, 0, 0, 0, 0, 0,
// }
//
// var LateKnightPst = [64]int16{
// 	-44, -50, -15, -39, -44, -33, -74, -99,
// 	-30, -14, -47, -18, -28, -51, -34, -61,
// 	-36, -39, -12, -13, -29, -37, -38, -62,
// 	-24, -9, 7, 9, 15, -6, -2, -29,
// 	-27, -24, 0, 13, 2, 4, -6, -24,
// 	-31, -14, -18, 0, -7, -18, -33, -34,
// 	-42, -27, -19, -15, -12, -30, -31, -54,
// 	-23, -57, -31, -15, -28, -27, -58, -74,
// }
//
// var LateBishopPst = [64]int16{
// 	-10, -26, -4, -7, -4, -7, -15, -26,
// 	-11, -16, -3, -15, -15, -22, -15, -13,
// 	-2, -20, -17, -18, -17, -16, -9, -2,
// 	-5, 1, 2, -4, -1, -3, -8, 1,
// 	-12, -9, 2, 5, -10, -2, -14, -12,
// 	-11, -9, 0, 0, 4, -13, -10, -15,
// 	-16, -23, -13, -4, -2, -15, -20, -32,
// 	-19, -9, -18, -4, -10, -12, -4, -13,
// }
//
// var LateRookPst = [64]int16{
// 	8, 2, 15, 3, 5, 13, 6, 7,
// 	8, 12, 2, 3, -17, -6, 12, 3,
// 	12, 9, 2, 4, 5, -6, -8, 2,
// 	15, 9, 15, 0, 7, 6, 2, 15,
// 	15, 19, 16, 12, 4, 8, -3, 5,
// 	12, 12, 7, 12, 2, -2, 4, -2,
// 	10, 4, 12, 12, 1, -3, -4, 15,
// 	4, 11, 9, 0, -3, -4, 10, -14,
// }
//
// var LateQueenPst = [64]int16{
// 	32, 59, 50, 50, 41, 36, 25, 58,
// 	5, 45, 51, 63, 91, 41, 57, 39,
// 	-1, 22, 18, 84, 67, 49, 51, 37,
// 	43, 44, 44, 71, 77, 60, 92, 68,
// 	0, 57, 42, 67, 52, 55, 67, 53,
// 	26, -24, 31, 19, 30, 36, 47, 40,
// 	-4, -9, -20, 0, 4, -1, -15, -12,
// 	-19, -17, -7, -30, 21, -13, 1, -28,
// }
//
// var LateKingPst = [64]int16{
// 	-72, -51, -34, -29, -4, 17, -2, -15,
// 	-29, 1, 1, -3, 6, 31, 16, 20,
// 	3, 5, 5, 6, 4, 30, 29, 15,
// 	-7, 13, 15, 23, 19, 28, 21, 10,
// 	-14, -14, 19, 27, 29, 22, 9, 0,
// 	-20, -8, 9, 22, 24, 18, 3, -3,
// 	-28, -19, 5, 12, 14, 4, -11, -21,
// 	-54, -48, -23, 3, -26, -4, -37, -56,
// }
//
// var MiddlegameBackwardPawnPenalty int16 = 10
// var EndgameBackwardPawnPenalty int16 = 4
// var MiddlegameIsolatedPawnPenalty int16 = 11
// var EndgameIsolatedPawnPenalty int16 = 6
// var MiddlegameDoublePawnPenalty int16 = 2
// var EndgameDoublePawnPenalty int16 = 26
// var MiddlegamePassedPawnAward int16 = 0
// var EndgamePassedPawnAward int16 = 10
// var MiddlegameAdvancedPassedPawnAward int16 = 10
// var EndgameAdvancedPassedPawnAward int16 = 62
// var MiddlegameCandidatePassedPawnAward int16 = 32
// var EndgameCandidatePassedPawnAward int16 = 51
// var MiddlegameRookOpenFileAward int16 = 45
// var EndgameRookOpenFileAward int16 = 0
// var MiddlegameRookSemiOpenFileAward int16 = 14
// var EndgameRookSemiOpenFileAward int16 = 20
// var MiddlegameVeritcalDoubleRookAward int16 = 10
// var EndgameVeritcalDoubleRookAward int16 = 10
// var MiddlegameHorizontalDoubleRookAward int16 = 27
// var EndgameHorizontalDoubleRookAward int16 = 12
// var MiddlegamePawnFactorCoeff int16 = 0
// var EndgamePawnFactorCoeff int16 = 0
// var MiddlegameMobilityFactorCoeff int16 = 6
// var EndgameMobilityFactorCoeff int16 = 3
// var MiddlegameAggressivityFactorCoeff int16 = 1
// var EndgameAggressivityFactorCoeff int16 = 6
// var MiddlegameInnerPawnToKingAttackCoeff int16 = 0
// var EndgameInnerPawnToKingAttackCoeff int16 = 0
// var MiddlegameOuterPawnToKingAttackCoeff int16 = 4
// var EndgameOuterPawnToKingAttackCoeff int16 = 0
// var MiddlegameInnerMinorToKingAttackCoeff int16 = 18
// var EndgameInnerMinorToKingAttackCoeff int16 = 0
// var MiddlegameOuterMinorToKingAttackCoeff int16 = 11
// var EndgameOuterMinorToKingAttackCoeff int16 = 1
// var MiddlegameInnerMajorToKingAttackCoeff int16 = 17
// var EndgameInnerMajorToKingAttackCoeff int16 = 0
// var MiddlegameOuterMajorToKingAttackCoeff int16 = 8
// var EndgameOuterMajorToKingAttackCoeff int16 = 5
// var MiddlegamePawnShieldPenalty int16 = 10
// var EndgamePawnShieldPenalty int16 = 8
// var MiddlegameNotCastlingPenalty int16 = 26
// var EndgameNotCastlingPenalty int16 = 5
// var MiddlegameKingZoneOpenFilePenalty int16 = 35
// var EndgameKingZoneOpenFilePenalty int16 = 0
// var MiddlegameKingZoneMissingPawnPenalty int16 = 16
// var EndgameKingZoneMissingPawnPenalty int16 = 0
// var MiddlegamePawnIslandPenalty int16 = 2
// var EndgamePawnIslandPenalty int16 = 10

// // Middle-game
// var EarlyPawnPst = [64]int16{
// 	0, 0, 0, 0, 0, 0, 0, 0,
// 	97, 128, 65, 103, 97, 124, 17, -31,
// 	-15, -12, 19, 20, 59, 71, 13, -20,
// 	-23, -7, -5, 15, 15, 11, 1, -31,
// 	-34, -23, -12, 5, 8, 3, -9, -34,
// 	-35, -30, -20, -22, -10, -9, 6, -25,
// 	-43, -25, -32, -34, -30, 8, 10, -33,
// 	0, 0, 0, 0, 0, 0, 0, 0,
// }
//
// var EarlyKnightPst = [64]int16{
// 	-210, -82, -47, -45, 54, -119, -25, -139,
// 	-79, -41, 78, 40, 30, 75, 6, -13,
// 	-41, 68, 47, 66, 95, 136, 70, 54,
// 	7, 31, 23, 54, 26, 72, 16, 34,
// 	8, 33, 32, 24, 43, 26, 33, 8,
// 	-2, 12, 27, 34, 47, 33, 46, 8,
// 	-3, -26, 15, 23, 22, 37, 18, 15,
// 	-99, 7, -28, -13, 16, 1, 12, 3,
// }
//
// var EarlyBishopPst = [64]int16{
// 	-25, 22, -79, -42, -19, -37, 13, 7,
// 	3, 48, 13, 1, 56, 76, 41, -26,
// 	19, 67, 80, 65, 60, 78, 54, 24,
// 	31, 39, 42, 71, 58, 52, 35, 23,
// 	36, 51, 47, 59, 68, 48, 46, 40,
// 	33, 54, 54, 53, 55, 74, 56, 44,
// 	45, 64, 57, 45, 57, 63, 81, 45,
// 	5, 39, 36, 29, 39, 34, 2, 17,
// }
//
// var EarlyRookPst = [64]int16{
// 	-1, 14, -17, 22, 22, -17, 2, -7,
// 	4, 1, 33, 31, 59, 61, -3, 21,
// 	-35, -9, -3, -3, -28, 25, 43, -16,
// 	-43, -30, -15, 1, -15, 16, -18, -38,
// 	-51, -43, -28, -26, -12, -27, 0, -38,
// 	-52, -35, -31, -33, -19, -7, -13, -34,
// 	-50, -26, -34, -24, -15, 3, -14, -72,
// 	-19, -20, -12, -2, 0, -1, -37, -18,
// }
//
// var EarlyQueenPst = [64]int16{
// 	-64, -29, -11, -18, 35, 33, 33, 10,
// 	-32, -61, -24, -24, -63, 26, -4, 20,
// 	-18, -24, -13, -40, -8, 29, 2, 17,
// 	-40, -38, -37, -47, -32, -21, -33, -25,
// 	-9, -42, -24, -24, -23, -18, -16, -14,
// 	-24, 6, -9, -3, -6, 0, 7, 0,
// 	-24, -2, 18, 10, 17, 23, 8, 17,
// 	13, -6, 8, 23, -1, -8, -9, -31,
// }
//
// var EarlyKingPst = [64]int16{
// 	-54, 94, 90, 44, -45, -15, 43, 42,
// 	86, 32, 25, 58, 15, -1, -21, -61,
// 	35, 36, 50, 7, 25, 56, 50, -13,
// 	-20, -8, 8, -20, -20, -20, -21, -55,
// 	-46, 10, -37, -71, -71, -50, -59, -77,
// 	-3, -16, -25, -58, -58, -46, -21, -37,
// 	13, 16, -10, -61, -40, -15, 10, 17,
// 	-6, 37, 15, -66, -8, -33, 28, 26,
// }
//
// // Endgame
// var LatePawnPst = [64]int16{
// 	0, 0, 0, 0, 0, 0, 0, 0,
// 	167, 150, 133, 104, 117, 104, 154, 189,
// 	86, 84, 62, 35, 18, 25, 60, 71,
// 	23, 4, -7, -27, -17, -12, 1, 9,
// 	20, 9, -2, -9, -9, -8, -2, 4,
// 	4, 0, -9, -5, -4, -7, -15, -12,
// 	15, -1, 5, 3, 11, -8, -13, -10,
// 	0, 0, 0, 0, 0, 0, 0, 0,
// }
//
// var LateKnightPst = [64]int16{
// 	-44, -50, -17, -40, -44, -34, -74, -98,
// 	-29, -13, -45, -18, -26, -51, -36, -61,
// 	-36, -39, -12, -12, -30, -37, -39, -62,
// 	-25, -8, 7, 8, 15, -6, -1, -28,
// 	-27, -24, 0, 13, 2, 4, -5, -25,
// 	-32, -13, -17, 0, -7, -18, -33, -32,
// 	-44, -27, -19, -15, -12, -28, -31, -53,
// 	-22, -57, -29, -16, -28, -25, -56, -72,
// }
//
// var LateBishopPst = [64]int16{
// 	-10, -25, -6, -9, -6, -7, -15, -24,
// 	-10, -16, -1, -15, -14, -22, -16, -13,
// 	-4, -19, -18, -17, -16, -15, -10, -2,
// 	-6, 1, 2, -3, 0, -3, -8, 1,
// 	-12, -10, 3, 5, -9, -2, -13, -13,
// 	-11, -9, 0, 1, 4, -13, -9, -15,
// 	-17, -24, -12, -3, -2, -14, -19, -31,
// 	-17, -9, -18, -4, -10, -12, -3, -13,
// }
//
// var LateRookPst = [64]int16{
// 	9, 3, 16, 4, 6, 14, 8, 8,
// 	8, 12, 2, 3, -17, -6, 12, 3,
// 	12, 9, 3, 4, 5, -6, -8, 2,
// 	15, 8, 14, 0, 5, 5, 2, 14,
// 	16, 17, 16, 11, 3, 6, -3, 5,
// 	12, 12, 7, 11, 2, -2, 3, -2,
// 	9, 4, 11, 12, 1, -4, -4, 15,
// 	4, 11, 9, 0, -3, -4, 10, -14,
// }
//
// var LateQueenPst = [64]int16{
// 	32, 58, 50, 50, 42, 36, 26, 58,
// 	5, 44, 50, 64, 90, 41, 57, 36,
// 	-1, 24, 19, 84, 66, 48, 51, 34,
// 	43, 44, 44, 69, 76, 60, 91, 68,
// 	0, 56, 43, 66, 51, 55, 66, 50,
// 	26, -24, 30, 17, 30, 36, 46, 37,
// 	-7, -10, -21, 0, 4, -3, -17, -12,
// 	-22, -17, -8, -30, 20, -14, 1, -28,
// }
//
// var LateKingPst = [64]int16{
// 	-72, -54, -35, -30, -5, 14, -3, -16,
// 	-30, 2, -1, -2, 6, 30, 15, 20,
// 	3, 7, 6, 6, 4, 30, 30, 13,
// 	-8, 12, 15, 23, 20, 28, 20, 9,
// 	-14, -13, 20, 28, 29, 23, 8, 0,
// 	-21, -8, 9, 23, 25, 18, 2, -3,
// 	-30, -20, 6, 14, 15, 5, -11, -21,
// 	-53, -48, -23, 4, -26, -4, -37, -56,
// }
//
// var MiddlegameBackwardPawnPenalty int16 = 9
// var EndgameBackwardPawnPenalty int16 = 3
// var MiddlegameIsolatedPawnPenalty int16 = 11
// var EndgameIsolatedPawnPenalty int16 = 6
// var MiddlegameDoublePawnPenalty int16 = 2
// var EndgameDoublePawnPenalty int16 = 26
// var MiddlegamePassedPawnAward int16 = 0
// var EndgamePassedPawnAward int16 = 9
// var MiddlegameAdvancedPassedPawnAward int16 = 9
// var EndgameAdvancedPassedPawnAward int16 = 59
// var MiddlegameCandidatePassedPawnAward int16 = 32
// var EndgameCandidatePassedPawnAward int16 = 50
// var MiddlegameRookOpenFileAward int16 = 45
// var EndgameRookOpenFileAward int16 = 0
// var MiddlegameRookSemiOpenFileAward int16 = 14
// var EndgameRookSemiOpenFileAward int16 = 20
// var MiddlegameVeritcalDoubleRookAward int16 = 10
// var EndgameVeritcalDoubleRookAward int16 = 9
// var MiddlegameHorizontalDoubleRookAward int16 = 26
// var EndgameHorizontalDoubleRookAward int16 = 10
// var MiddlegamePawnFactorCoeff int16 = 0
// var EndgamePawnFactorCoeff int16 = 0
// var MiddlegameMobilityFactorCoeff int16 = 6
// var EndgameMobilityFactorCoeff int16 = 3
// var MiddlegameAggressivityFactorCoeff int16 = 1
// var EndgameAggressivityFactorCoeff int16 = 5
// var MiddlegameInnerPawnToKingAttackCoeff int16 = 0
// var EndgameInnerPawnToKingAttackCoeff int16 = 0
// var MiddlegameOuterPawnToKingAttackCoeff int16 = 3
// var EndgameOuterPawnToKingAttackCoeff int16 = 0
// var MiddlegameInnerMinorToKingAttackCoeff int16 = 17
// var EndgameInnerMinorToKingAttackCoeff int16 = 0
// var MiddlegameOuterMinorToKingAttackCoeff int16 = 10
// var EndgameOuterMinorToKingAttackCoeff int16 = 1
// var MiddlegameInnerMajorToKingAttackCoeff int16 = 16
// var EndgameInnerMajorToKingAttackCoeff int16 = 0
// var MiddlegameOuterMajorToKingAttackCoeff int16 = 7
// var EndgameOuterMajorToKingAttackCoeff int16 = 5
// var MiddlegamePawnShieldPenalty int16 = 10
// var EndgamePawnShieldPenalty int16 = 6
// var MiddlegameNotCastlingPenalty int16 = 25
// var EndgameNotCastlingPenalty int16 = 5
// var MiddlegameKingZoneOpenFilePenalty int16 = 33
// var EndgameKingZoneOpenFilePenalty int16 = 0
// var MiddlegameKingZoneMissingPawnPenalty int16 = 18
// var EndgameKingZoneMissingPawnPenalty int16 = 0

var flip = [64]int16{
	56, 57, 58, 59, 60, 61, 62, 63,
	48, 49, 50, 51, 52, 53, 54, 55,
	40, 41, 42, 43, 44, 45, 46, 47,
	32, 33, 34, 35, 36, 37, 38, 39,
	24, 25, 26, 27, 28, 29, 30, 31,
	16, 17, 18, 19, 20, 21, 22, 23,
	8, 9, 10, 11, 12, 13, 14, 15,
	0, 1, 2, 3, 4, 5, 6, 7,
}

func PSQT(piece Piece, sq Square, isEndgame bool) int16 {
	if isEndgame {
		switch piece {
		case WhitePawn:
			return LatePawnPst[flip[int(sq)]]
		case WhiteKnight:
			return LateKnightPst[flip[int(sq)]]
		case WhiteBishop:
			return LateBishopPst[flip[int(sq)]]
		case WhiteRook:
			return LateRookPst[flip[int(sq)]]
		case WhiteQueen:
			return LateQueenPst[flip[int(sq)]]
		case WhiteKing:
			return LateKingPst[flip[int(sq)]]
		case BlackPawn:
			return LatePawnPst[int(sq)]
		case BlackKnight:
			return LateKnightPst[int(sq)]
		case BlackBishop:
			return LateBishopPst[int(sq)]
		case BlackRook:
			return LateRookPst[int(sq)]
		case BlackQueen:
			return LateQueenPst[int(sq)]
		case BlackKing:
			return LateKingPst[int(sq)]
		}
	} else {
		switch piece {
		case WhitePawn:
			return EarlyPawnPst[flip[int(sq)]]
		case WhiteKnight:
			return EarlyKnightPst[flip[int(sq)]]
		case WhiteBishop:
			return EarlyBishopPst[flip[int(sq)]]
		case WhiteRook:
			return EarlyRookPst[flip[int(sq)]]
		case WhiteQueen:
			return EarlyQueenPst[flip[int(sq)]]
		case WhiteKing:
			return EarlyKingPst[flip[int(sq)]]
		case BlackPawn:
			return EarlyPawnPst[int(sq)]
		case BlackKnight:
			return EarlyKnightPst[int(sq)]
		case BlackBishop:
			return EarlyBishopPst[int(sq)]
		case BlackRook:
			return EarlyRookPst[int(sq)]
		case BlackQueen:
			return EarlyQueenPst[int(sq)]
		case BlackKing:
			return EarlyKingPst[int(sq)]
		}
	}
	return 0
}

func Evaluate(position *Position) int16 {
	board := position.Board
	turn := position.Turn()

	// Compute material balance
	bbBlackPawn := board.GetBitboardOf(BlackPawn)
	bbBlackKnight := board.GetBitboardOf(BlackKnight)
	bbBlackBishop := board.GetBitboardOf(BlackBishop)
	bbBlackRook := board.GetBitboardOf(BlackRook)
	bbBlackQueen := board.GetBitboardOf(BlackQueen)
	bbBlackKing := board.GetBitboardOf(BlackKing)

	bbWhitePawn := board.GetBitboardOf(WhitePawn)
	bbWhiteKnight := board.GetBitboardOf(WhiteKnight)
	bbWhiteBishop := board.GetBitboardOf(WhiteBishop)
	bbWhiteRook := board.GetBitboardOf(WhiteRook)
	bbWhiteQueen := board.GetBitboardOf(WhiteQueen)
	bbWhiteKing := board.GetBitboardOf(WhiteKing)

	blackPawnsCount := int16(0)
	blackKnightsCount := int16(0)
	blackBishopsCount := int16(0)
	blackRooksCount := int16(0)
	blackQueensCount := int16(0)

	whitePawnsCount := int16(0)
	whiteKnightsCount := int16(0)
	whiteBishopsCount := int16(0)
	whiteRooksCount := int16(0)
	whiteQueensCount := int16(0)

	blackCentipawnsMG := int16(0)
	blackCentipawnsEG := int16(0)

	whiteCentipawnsMG := int16(0)
	whiteCentipawnsEG := int16(0)

	whites := board.GetWhitePieces()
	blacks := board.GetBlackPieces()
	all := whites | blacks

	var whiteKingIndex, blackKingIndex int

	// PST for black pawns
	pieceIter := bbBlackPawn
	for pieceIter > 0 {
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		blackPawnsCount++
		blackCentipawnsEG += LatePawnPst[index]
		blackCentipawnsMG += EarlyPawnPst[index]
		pieceIter ^= mask
	}

	// PST for white pawns
	pieceIter = bbWhitePawn
	for pieceIter != 0 {
		whitePawnsCount++
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		whiteCentipawnsEG += LatePawnPst[flip[index]]
		whiteCentipawnsMG += EarlyPawnPst[flip[index]]
		pieceIter ^= mask
	}

	// PST for other black pieces
	pieceIter = bbBlackKnight
	for pieceIter != 0 {
		blackKnightsCount++
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		blackCentipawnsEG += LateKnightPst[index]
		blackCentipawnsMG += EarlyKnightPst[index]
		pieceIter ^= mask
	}

	pieceIter = bbBlackBishop
	for pieceIter != 0 {
		blackBishopsCount++
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		blackCentipawnsEG += LateBishopPst[index]
		blackCentipawnsMG += EarlyBishopPst[index]
		pieceIter ^= mask
	}

	pieceIter = bbBlackRook
	for pieceIter != 0 {
		blackRooksCount++
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		sq := Square(index)
		if blackRooksCount == 1 {
			if board.IsVerticalDoubleRook(sq, bbBlackRook, all) {
				// double-rook vertical
				blackCentipawnsEG += EndgameVeritcalDoubleRookAward
				blackCentipawnsMG += MiddlegameVeritcalDoubleRookAward
			} else if board.IsHorizontalDoubleRook(sq, bbBlackRook, all) {
				// double-rook horizontal
				blackCentipawnsMG += MiddlegameHorizontalDoubleRookAward
				blackCentipawnsEG += EndgameHorizontalDoubleRookAward
			}
		}
		blackCentipawnsEG += LateRookPst[index]
		blackCentipawnsMG += EarlyRookPst[index]
		pieceIter ^= mask
	}

	pieceIter = bbBlackQueen
	for pieceIter != 0 {
		blackQueensCount++
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		blackCentipawnsEG += LateQueenPst[index]
		blackCentipawnsMG += EarlyQueenPst[index]
		pieceIter ^= mask
	}

	pieceIter = bbBlackKing
	for pieceIter != 0 {
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		blackCentipawnsEG += LateKingPst[index]
		blackCentipawnsMG += EarlyKingPst[index]
		blackKingIndex = index
		pieceIter ^= mask
	}

	// PST for other white pieces
	pieceIter = bbWhiteKnight
	for pieceIter != 0 {
		whiteKnightsCount++
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		whiteCentipawnsEG += LateKnightPst[flip[index]]
		whiteCentipawnsMG += EarlyKnightPst[flip[index]]
		pieceIter ^= mask
	}

	pieceIter = bbWhiteBishop
	for pieceIter != 0 {
		whiteBishopsCount++
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		whiteCentipawnsEG += LateBishopPst[flip[index]]
		whiteCentipawnsMG += EarlyBishopPst[flip[index]]
		pieceIter ^= mask
	}

	pieceIter = bbWhiteRook
	for pieceIter != 0 {
		whiteRooksCount++
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		sq := Square(index)
		if whiteRooksCount == 1 {
			if board.IsVerticalDoubleRook(sq, bbWhiteRook, all) {
				// double-rook vertical
				whiteCentipawnsMG += MiddlegameVeritcalDoubleRookAward
				whiteCentipawnsEG += EndgameVeritcalDoubleRookAward
			} else if board.IsHorizontalDoubleRook(sq, bbWhiteRook, all) {
				// double-rook horizontal
				whiteCentipawnsMG += MiddlegameHorizontalDoubleRookAward
				whiteCentipawnsEG += EndgameHorizontalDoubleRookAward
			}
		}
		whiteCentipawnsEG += LateRookPst[flip[index]]
		whiteCentipawnsMG += EarlyRookPst[flip[index]]
		pieceIter ^= mask
	}

	pieceIter = bbWhiteQueen
	for pieceIter != 0 {
		whiteQueensCount++
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		whiteCentipawnsEG += LateQueenPst[flip[index]]
		whiteCentipawnsMG += EarlyQueenPst[flip[index]]
		pieceIter ^= mask
	}

	pieceIter = bbWhiteKing
	for pieceIter != 0 {
		index := bits.TrailingZeros64(pieceIter)
		mask := SquareMask(uint64(index))
		whiteCentipawnsEG += LateKingPst[flip[index]]
		whiteCentipawnsMG += EarlyKingPst[flip[index]]
		whiteKingIndex = index
		pieceIter ^= mask
	}

	pawnFactorMG := int16(16-blackPawnsCount-whitePawnsCount) * MiddlegamePawnFactorCoeff
	pawnFactorEG := int16(16-blackPawnsCount-whitePawnsCount) * EndgamePawnFactorCoeff

	blackCentipawnsMG += blackPawnsCount * BlackPawn.Weight()
	blackCentipawnsMG += blackKnightsCount * (BlackKnight.Weight() - pawnFactorMG)
	blackCentipawnsMG += blackBishopsCount * (BlackBishop.Weight())
	blackCentipawnsMG += blackRooksCount * (BlackRook.Weight() + pawnFactorMG)
	blackCentipawnsMG += blackQueensCount * BlackQueen.Weight()

	blackCentipawnsEG += blackPawnsCount * BlackPawn.Weight()
	blackCentipawnsEG += blackKnightsCount * (BlackKnight.Weight() - pawnFactorEG)
	blackCentipawnsEG += blackBishopsCount * (BlackBishop.Weight())
	blackCentipawnsEG += blackRooksCount * (BlackRook.Weight() + pawnFactorEG)
	blackCentipawnsEG += blackQueensCount * BlackQueen.Weight()

	whiteCentipawnsMG += whitePawnsCount * WhitePawn.Weight()
	whiteCentipawnsMG += whiteKnightsCount * (WhiteKnight.Weight() - pawnFactorMG)
	whiteCentipawnsMG += whiteBishopsCount * (WhiteBishop.Weight())
	whiteCentipawnsMG += whiteRooksCount * (WhiteRook.Weight() + pawnFactorMG)
	whiteCentipawnsMG += whiteQueensCount * WhiteQueen.Weight()

	whiteCentipawnsEG += whitePawnsCount * WhitePawn.Weight()
	whiteCentipawnsEG += whiteKnightsCount * (WhiteKnight.Weight() - pawnFactorEG)
	whiteCentipawnsEG += whiteBishopsCount * (WhiteBishop.Weight())
	whiteCentipawnsEG += whiteRooksCount * (WhiteRook.Weight() + pawnFactorEG)
	whiteCentipawnsEG += whiteQueensCount * WhiteQueen.Weight()

	mobilityEval := Mobility(position, blackKingIndex, whiteKingIndex)

	whiteCentipawnsMG += mobilityEval.whiteMG
	whiteCentipawnsEG += mobilityEval.whiteEG
	blackCentipawnsMG += mobilityEval.blackMG
	blackCentipawnsEG += mobilityEval.blackEG

	rookEval := RookFilesEval(bbBlackRook, bbWhiteRook, bbBlackPawn, bbWhitePawn)
	whiteCentipawnsMG += rookEval.whiteMG
	whiteCentipawnsEG += rookEval.whiteEG
	blackCentipawnsMG += rookEval.blackMG
	blackCentipawnsEG += rookEval.blackEG

	pawnStructureEval := PawnStructureEval(position)
	whiteCentipawnsMG += pawnStructureEval.whiteMG
	whiteCentipawnsEG += pawnStructureEval.whiteEG
	blackCentipawnsMG += pawnStructureEval.blackMG
	blackCentipawnsEG += pawnStructureEval.blackEG

	kingSafetyEval := KingSafety(bbBlackKing, bbWhiteKing, bbBlackPawn, bbWhitePawn,
		position.HasTag(BlackCanCastleQueenSide) || position.HasTag(BlackCanCastleKingSide),
		position.HasTag(WhiteCanCastleQueenSide) || position.HasTag(WhiteCanCastleKingSide),
	)
	whiteCentipawnsMG += kingSafetyEval.whiteMG
	whiteCentipawnsEG += kingSafetyEval.whiteEG
	blackCentipawnsMG += kingSafetyEval.blackMG
	blackCentipawnsEG += kingSafetyEval.blackEG

	phase := TotalPhase -
		whitePawnsCount*PawnPhase -
		blackPawnsCount*PawnPhase -
		whiteKnightsCount*KnightPhase -
		blackKnightsCount*KnightPhase -
		whiteBishopsCount*BishopPhase -
		blackBishopsCount*BishopPhase -
		whiteRooksCount*RookPhase -
		blackRooksCount*RookPhase -
		whiteQueensCount*QueenPhase -
		blackQueensCount*QueenPhase

	phase = (phase*256 + HalfPhase) / TotalPhase

	var evalEG, evalMG int16

	if turn == White {
		evalEG = whiteCentipawnsEG - blackCentipawnsEG
		evalMG = whiteCentipawnsMG - blackCentipawnsMG
	} else {
		evalEG = blackCentipawnsEG - whiteCentipawnsEG
		evalMG = blackCentipawnsMG - whiteCentipawnsMG
	}

	// The following formula overflows if I do not convert to int32 first
	// then I have to convert back to int16, as the function return requires
	// and that is also safe, due to the division
	mg := int32(evalMG)
	eg := int32(evalEG)
	phs := int32(phase)
	taperedEval := int16(((mg * (256 - phs)) + eg*phs) / 256)
	return toEval(taperedEval + Tempo)
}

func RookFilesEval(blackRook uint64, whiteRook uint64, blackPawns uint64, whitePawns uint64) Eval {
	var blackMG, whiteMG, blackEG, whiteEG int16

	blackFiles := FileFill(blackRook)
	whiteFiles := FileFill(whiteRook)

	allPawns := FileFill(blackPawns | whitePawns)

	// open files
	blackRooksNoPawns := blackFiles &^ allPawns
	whiteRooksNoPawns := whiteFiles &^ allPawns

	blackRookOpenFiles := blackRook & blackRooksNoPawns
	whiteRookOpenFiles := whiteRook & whiteRooksNoPawns

	count := int16(bits.OnesCount64(blackRookOpenFiles))
	blackMG += MiddlegameRookOpenFileAward * count
	blackEG += EndgameRookOpenFileAward * count

	count = int16(bits.OnesCount64(whiteRookOpenFiles))
	whiteMG += MiddlegameRookOpenFileAward * count
	whiteEG += EndgameRookOpenFileAward * count

	// semi-open files
	blackRooksNoOwnPawns := blackFiles &^ FileFill(blackPawns)
	whiteRooksNoOwnPawns := whiteFiles &^ FileFill(whitePawns)

	blackRookSemiOpenFiles := (blackRook &^ blackRookOpenFiles) & blackRooksNoOwnPawns
	whiteRookSemiOpenFiles := (whiteRook &^ whiteRookOpenFiles) & whiteRooksNoOwnPawns

	count = int16(bits.OnesCount64(blackRookSemiOpenFiles))
	blackMG += MiddlegameRookSemiOpenFileAward * count
	blackEG += EndgameRookSemiOpenFileAward * count

	count = int16(bits.OnesCount64(whiteRookSemiOpenFiles))
	whiteMG += MiddlegameRookSemiOpenFileAward * count
	whiteEG += EndgameRookSemiOpenFileAward * count

	return Eval{blackMG: blackMG, whiteMG: whiteMG, blackEG: blackEG, whiteEG: whiteEG}
}

func PawnStructureEval(p *Position) Eval {
	var blackMG, whiteMG, blackEG, whiteEG int16

	// passed pawns
	countP, countS := p.CountPassedPawns(Black)
	blackMG += MiddlegamePassedPawnAward * countP
	blackEG += EndgamePassedPawnAward * countP

	blackMG += MiddlegameAdvancedPassedPawnAward * countS
	blackEG += EndgameAdvancedPassedPawnAward * countS

	countP, countS = p.CountPassedPawns(White)
	whiteMG += MiddlegamePassedPawnAward * countP
	whiteEG += EndgamePassedPawnAward * countP

	whiteMG += MiddlegameAdvancedPassedPawnAward * countS
	whiteEG += EndgameAdvancedPassedPawnAward * countS

	// candidate passed pawns
	count := p.CountCandidatePawns(Black)
	blackMG += MiddlegameCandidatePassedPawnAward * count
	blackEG += EndgameCandidatePassedPawnAward * count

	count = p.CountCandidatePawns(White)
	whiteMG += MiddlegameCandidatePassedPawnAward * count
	whiteEG += EndgameCandidatePassedPawnAward * count

	// backward pawns
	count = p.CountBackwardPawns(Black)
	blackMG -= MiddlegameBackwardPawnPenalty * count
	blackEG -= EndgameBackwardPawnPenalty * count

	count = p.CountBackwardPawns(White)
	whiteMG -= MiddlegameBackwardPawnPenalty * count
	whiteEG -= EndgameBackwardPawnPenalty * count

	// isolated pawns
	count = p.CountIsolatedPawns(Black)
	blackMG -= MiddlegameIsolatedPawnPenalty * count
	blackEG -= EndgameIsolatedPawnPenalty * count

	count = p.CountIsolatedPawns(White)
	whiteMG -= MiddlegameIsolatedPawnPenalty * count
	whiteEG -= EndgameIsolatedPawnPenalty * count

	// double pawns
	count = p.CountDoublePawns(Black)
	blackMG -= MiddlegameDoublePawnPenalty * count
	blackEG -= EndgameDoublePawnPenalty * count

	count = p.CountDoublePawns(White)
	whiteMG -= MiddlegameDoublePawnPenalty * count
	whiteEG -= EndgameDoublePawnPenalty * count

	// double pawns
	count = p.CountPawnIslands(Black)
	blackMG -= MiddlegamePawnIslandPenalty * count
	blackEG -= EndgamePawnIslandPenalty * count

	count = p.CountPawnIslands(White)
	whiteMG -= MiddlegamePawnIslandPenalty * count
	whiteEG -= EndgamePawnIslandPenalty * count

	return Eval{blackMG: blackMG, whiteMG: whiteMG, blackEG: blackEG, whiteEG: whiteEG}
}

func kingSafetyPenalty(color Color, side PieceType, ownPawn uint64, allPawn uint64) (int16, int16) {
	var mg, eg int16
	var a_shield, b_shield, c_shield, f_shield, g_shield, h_shield uint64
	if color == White {
		a_shield = WhiteAShield
		b_shield = WhiteBShield
		c_shield = WhiteCShield
		f_shield = WhiteFShield
		g_shield = WhiteGShield
		h_shield = WhiteHShield
	} else {
		a_shield = BlackAShield
		b_shield = BlackBShield
		c_shield = BlackCShield
		f_shield = BlackFShield
		g_shield = BlackGShield
		h_shield = BlackHShield
	}
	if side == King {
		if H_FileFill&allPawn == 0 { // no pawns, super bad
			mg += MiddlegameKingZoneOpenFilePenalty
			eg += EndgameKingZoneOpenFilePenalty
		} else if H_FileFill&ownPawn == 0 { // semi-open file, bad
			mg += MiddlegameKingZoneMissingPawnPenalty
			eg += EndgameKingZoneMissingPawnPenalty
		} else if h_shield&ownPawn == 0 {
			mg += MiddlegamePawnShieldPenalty
			eg += EndgamePawnShieldPenalty
		}

		if G_FileFill&allPawn == 0 { // no pawns, super bad
			mg += MiddlegameKingZoneOpenFilePenalty
			eg += EndgameKingZoneOpenFilePenalty
		} else if G_FileFill&ownPawn == 0 { // semi-open file, bad
			mg += MiddlegameKingZoneMissingPawnPenalty
			eg += EndgameKingZoneMissingPawnPenalty
		} else if g_shield&ownPawn == 0 {
			mg += MiddlegamePawnShieldPenalty
			eg += EndgamePawnShieldPenalty
		}

		if F_FileFill&allPawn == 0 { // no pawns, super bad
			mg += MiddlegameKingZoneOpenFilePenalty
			eg += EndgameKingZoneOpenFilePenalty
		} else if F_FileFill&ownPawn == 0 { // semi-open file, bad
			mg += MiddlegameKingZoneMissingPawnPenalty
			eg += EndgameKingZoneMissingPawnPenalty
		} else if f_shield&ownPawn == 0 {
			mg += MiddlegamePawnShieldPenalty
			eg += EndgamePawnShieldPenalty
		}
	} else {
		if C_FileFill&allPawn == 0 { // no pawns, super bad
			mg += MiddlegameKingZoneOpenFilePenalty
			eg += EndgameKingZoneOpenFilePenalty
		} else if C_FileFill&ownPawn == 0 { // semi-open file, bad
			mg += MiddlegameKingZoneMissingPawnPenalty
			eg += EndgameKingZoneMissingPawnPenalty
		} else if c_shield&ownPawn == 0 {
			mg += MiddlegamePawnShieldPenalty
			eg += EndgamePawnShieldPenalty
		}

		if B_FileFill&allPawn == 0 { // no pawns, super bad
			mg += MiddlegameKingZoneOpenFilePenalty
			eg += EndgameKingZoneOpenFilePenalty
		} else if B_FileFill&ownPawn == 0 { // semi-open file, bad
			mg += MiddlegameKingZoneMissingPawnPenalty
			eg += EndgameKingZoneMissingPawnPenalty
		} else if b_shield&ownPawn == 0 {
			mg += MiddlegamePawnShieldPenalty
			eg += EndgamePawnShieldPenalty
		}

		if A_FileFill&allPawn == 0 { // no pawns, super bad
			mg += MiddlegameKingZoneOpenFilePenalty
			eg += EndgameKingZoneOpenFilePenalty
		} else if A_FileFill&ownPawn == 0 { // semi-open file, bad
			mg += MiddlegameKingZoneMissingPawnPenalty
			eg += EndgameKingZoneMissingPawnPenalty
		} else if a_shield&ownPawn == 0 {
			mg += MiddlegamePawnShieldPenalty
			eg += EndgamePawnShieldPenalty
		}
	}

	return mg, eg
}

func KingSafety(blackKing uint64, whiteKing uint64, blackPawn uint64,
	whitePawn uint64, blackCastleFlag bool, whiteCastleFlag bool) Eval {
	var whiteCentipawnsMG, whiteCentipawnsEG, blackCentipawnsMG, blackCentipawnsEG int16
	allPawn := whitePawn | blackPawn

	if blackKing&BlackKingSideMask != 0 {
		blackCastleFlag = true
		// Missing pawn shield
		mg, eg := kingSafetyPenalty(Black, King, blackPawn, allPawn)
		blackCentipawnsMG -= mg
		blackCentipawnsEG -= eg
	} else if blackKing&BlackQueenSideMask != 0 {
		blackCastleFlag = true
		// Missing pawn shield
		mg, eg := kingSafetyPenalty(Black, Queen, blackPawn, allPawn)
		blackCentipawnsMG -= mg
		blackCentipawnsEG -= eg
	}

	if whiteKing&WhiteKingSideMask != 0 {
		whiteCastleFlag = true
		// Missing pawn shield
		mg, eg := kingSafetyPenalty(White, King, whitePawn, allPawn)
		whiteCentipawnsMG -= mg
		whiteCentipawnsEG -= eg
	} else if whiteKing&WhiteQueenSideMask != 0 {
		whiteCastleFlag = true
		// Missing pawn shield
		mg, eg := kingSafetyPenalty(White, Queen, whitePawn, allPawn)
		whiteCentipawnsMG -= mg
		whiteCentipawnsEG -= eg
	}

	if !whiteCastleFlag {
		whiteCentipawnsMG -= MiddlegameNotCastlingPenalty
		whiteCentipawnsEG -= EndgameNotCastlingPenalty
	}

	if !blackCastleFlag {
		blackCentipawnsMG -= MiddlegameNotCastlingPenalty
		blackCentipawnsEG -= EndgameNotCastlingPenalty
	}
	return Eval{blackMG: blackCentipawnsMG, whiteMG: whiteCentipawnsMG, blackEG: blackCentipawnsEG, whiteEG: whiteCentipawnsEG}
}

func Mobility(p *Position, blackKingIndex int, whiteKingIndex int) Eval {
	board := p.Board
	var whiteCentipawnsMG, whiteCentipawnsEG, blackCentipawnsMG, blackCentipawnsEG int16

	// mobility and attacks
	whitePawnAttacks, whiteMinorAttacks, whiteOtherAttacks := board.AllAttacks(White) // get the squares that are attacked by white
	blackPawnAttacks, blackMinorAttacks, blackOtherAttacks := board.AllAttacks(Black) // get the squares that are attacked by black

	blackKingZone := SquareInnerRingMask[blackKingIndex] | SquareOuterRingMask[blackKingIndex]
	whiteKingZone := SquareInnerRingMask[whiteKingIndex] | SquareOuterRingMask[whiteKingIndex]

	// king attacks are considered later
	whiteAttacks := (whitePawnAttacks | whiteMinorAttacks | whiteOtherAttacks) &^ blackKingZone
	blackAttacks := (blackPawnAttacks | blackMinorAttacks | blackOtherAttacks) &^ whiteKingZone

	wQuietAttacks := bits.OnesCount64(whiteAttacks << 32) // keep hi-bits only
	bQuietAttacks := bits.OnesCount64(blackAttacks >> 32) // keep lo-bits only

	whiteAggressivity := bits.OnesCount64(whiteAttacks >> 32) // keep hi-bits only
	blackAggressivity := bits.OnesCount64(blackAttacks << 32) // keep lo-bits only

	whiteCentipawnsMG += MiddlegameMobilityFactorCoeff * int16(wQuietAttacks)
	whiteCentipawnsEG += EndgameMobilityFactorCoeff * int16(wQuietAttacks)

	blackCentipawnsMG += MiddlegameMobilityFactorCoeff * int16(bQuietAttacks)
	blackCentipawnsEG += EndgameMobilityFactorCoeff * int16(bQuietAttacks)

	whiteCentipawnsMG += MiddlegameAggressivityFactorCoeff * int16(whiteAggressivity)
	whiteCentipawnsEG += EndgameAggressivityFactorCoeff * int16(whiteAggressivity)

	blackCentipawnsMG += MiddlegameAggressivityFactorCoeff * int16(blackAggressivity)
	blackCentipawnsEG += EndgameAggressivityFactorCoeff * int16(blackAggressivity)

	whiteCentipawnsMG +=
		MiddlegameInnerPawnToKingAttackCoeff*int16(bits.OnesCount64(whitePawnAttacks&SquareInnerRingMask[blackKingIndex])) +
			MiddlegameOuterPawnToKingAttackCoeff*int16(bits.OnesCount64(whitePawnAttacks&SquareOuterRingMask[blackKingIndex])) +
			MiddlegameInnerMinorToKingAttackCoeff*int16(bits.OnesCount64(whiteMinorAttacks&SquareInnerRingMask[blackKingIndex])) +
			MiddlegameOuterMinorToKingAttackCoeff*int16(bits.OnesCount64(whiteMinorAttacks&SquareOuterRingMask[blackKingIndex])) +
			MiddlegameInnerMajorToKingAttackCoeff*int16(bits.OnesCount64(whiteOtherAttacks&SquareInnerRingMask[blackKingIndex])) +
			MiddlegameOuterMajorToKingAttackCoeff*int16(bits.OnesCount64(whiteOtherAttacks&SquareOuterRingMask[blackKingIndex]))

	whiteCentipawnsEG +=
		EndgameInnerPawnToKingAttackCoeff*int16(bits.OnesCount64(whitePawnAttacks&SquareInnerRingMask[blackKingIndex])) +
			EndgameOuterPawnToKingAttackCoeff*int16(bits.OnesCount64(whitePawnAttacks&SquareOuterRingMask[blackKingIndex])) +
			EndgameInnerMinorToKingAttackCoeff*int16(bits.OnesCount64(whiteMinorAttacks&SquareInnerRingMask[blackKingIndex])) +
			EndgameOuterMinorToKingAttackCoeff*int16(bits.OnesCount64(whiteMinorAttacks&SquareOuterRingMask[blackKingIndex])) +
			EndgameInnerMajorToKingAttackCoeff*int16(bits.OnesCount64(whiteOtherAttacks&SquareInnerRingMask[blackKingIndex])) +
			EndgameOuterMajorToKingAttackCoeff*int16(bits.OnesCount64(whiteOtherAttacks&SquareOuterRingMask[blackKingIndex]))

	blackCentipawnsMG +=
		MiddlegameInnerPawnToKingAttackCoeff*int16(bits.OnesCount64(blackPawnAttacks&SquareInnerRingMask[whiteKingIndex])) +
			MiddlegameOuterPawnToKingAttackCoeff*int16(bits.OnesCount64(blackPawnAttacks&SquareOuterRingMask[whiteKingIndex])) +
			MiddlegameInnerMinorToKingAttackCoeff*int16(bits.OnesCount64(blackMinorAttacks&SquareInnerRingMask[whiteKingIndex])) +
			MiddlegameOuterMinorToKingAttackCoeff*int16(bits.OnesCount64(blackMinorAttacks&SquareOuterRingMask[whiteKingIndex])) +
			MiddlegameInnerMajorToKingAttackCoeff*int16(bits.OnesCount64(blackOtherAttacks&SquareInnerRingMask[whiteKingIndex])) +
			MiddlegameOuterMajorToKingAttackCoeff*int16(bits.OnesCount64(blackOtherAttacks&SquareOuterRingMask[whiteKingIndex]))

	blackCentipawnsEG +=
		EndgameInnerPawnToKingAttackCoeff*int16(bits.OnesCount64(blackPawnAttacks&SquareInnerRingMask[whiteKingIndex])) +
			EndgameOuterPawnToKingAttackCoeff*int16(bits.OnesCount64(blackPawnAttacks&SquareOuterRingMask[whiteKingIndex])) +
			EndgameInnerMinorToKingAttackCoeff*int16(bits.OnesCount64(blackMinorAttacks&SquareInnerRingMask[whiteKingIndex])) +
			EndgameOuterMinorToKingAttackCoeff*int16(bits.OnesCount64(blackMinorAttacks&SquareOuterRingMask[whiteKingIndex])) +
			EndgameInnerMajorToKingAttackCoeff*int16(bits.OnesCount64(blackOtherAttacks&SquareInnerRingMask[whiteKingIndex])) +
			EndgameOuterMajorToKingAttackCoeff*int16(bits.OnesCount64(blackOtherAttacks&SquareOuterRingMask[whiteKingIndex]))

	return Eval{blackMG: blackCentipawnsMG, whiteMG: whiteCentipawnsMG, blackEG: blackCentipawnsEG, whiteEG: whiteCentipawnsEG}
}

func toEval(eval int16) int16 {
	if eval >= CHECKMATE_EVAL {
		return MAX_NON_CHECKMATE
	} else if eval <= -CHECKMATE_EVAL {
		return -MAX_NON_CHECKMATE
	}
	return eval
}

func min16(x int16, y int16) int16 {
	if x < y {
		return x
	}
	return y
}
