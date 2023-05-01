// generated by Textmapper; DO NOT EDIT

package test

import (
	"github.com/inspirer/textmapper/parsers/test/token"
)

const tmNumClasses = 37

type mapRange struct {
	lo         rune
	hi         rune
	defaultVal uint8
	val        []uint8
}

func mapRune(c rune) int {
	lo := 0
	hi := len(tmRuneRanges)
	for lo < hi {
		m := lo + (hi-lo)/2
		r := tmRuneRanges[m]
		if c < r.lo {
			hi = m
		} else if c >= r.hi {
			lo = m + 1
		} else {
			i := int(c - r.lo)
			if i < len(r.val) {
				return int(r.val[i])
			}
			return int(r.defaultVal)
		}
	}
	return 1
}

// Latin-1 characters.
var tmRuneClass = []uint8{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 4, 5, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 6, 1, 1, 7, 1, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 18, 1, 1, 1, 19, 1, 1, 20, 20, 20, 20,
	20, 20, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 22, 23, 24, 25, 1, 26, 1, 20, 20, 20, 20, 27, 28, 21, 21, 21, 21, 21,
	21, 21, 21, 29, 21, 30, 21, 31, 32, 33, 21, 21, 21, 21, 21, 34, 1, 35, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}

const tmRuneClassLen = 256
const tmFirstRule = -5

var tmRuneRanges = []mapRange{
	{8232, 8234, 36, nil},
}

var tmStateMap = []int{
	0, 48,
}

var tmToken = []token.Token{
	1, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 1, 35, 36, 38, 1, 38, 38,
	2,
}

var tmLexerAction = []int8{
	-6, -5, 47, 47, -5, 47, 46, 41, 40, 39, 38, -5, 37, 36, 34, 31, 28, 26, 25,
	-5, 24, 24, 17, 16, 15, 14, 13, 24, 24, 24, 24, 24, 3, 24, 2, 1, -5, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -1, -9, -9, 24, -9, -9, 24, 24,
	24, -9, -9, -9, 24, 4, 24, 24, 24, 24, 24, 24, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -1, -9, -9, 24, -9, -9, 24, 24, 24,
	-9, -9, -9, 24, 24, 24, 24, 24, 5, 24, 24, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -1, -9, -9, 24, -9, -9, 24, 24, 24, -9,
	-9, -9, 24, 24, 24, 24, 24, 24, 6, 24, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -2, -9, -9, 24, -9, -9, 24, 24, 24, -9, -9,
	-9, 24, 24, 7, 24, 24, 24, 24, 24, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -1, -9, -9, 24, -9, -9, 24, 24, 24, -9, -9, -9,
	24, 24, 24, 8, 24, 24, 24, 24, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -1, -9, -9, 24, -9, -9, 24, 24, 24, -9, -9, -9, 24,
	24, 24, 9, 24, 24, 24, 24, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -2, -9, -9, 24, -9, -9, 24, 24, 24, -9, -9, -9, 24, 24,
	24, 24, 24, 24, 24, 24, -9, -9, -9, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, 10, -5, -5, 24, -5, 11, 24, 24, 24, -5, -5, -5, 24, 24, 24,
	24, 24, 24, 24, 24, -5, -5, -5, -42, -42, -42, -42, -42, -42, -42, -42, -42,
	-42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42,
	-42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -42, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 12, -5, -5, 24, -5, -5, 24, 24,
	24, -5, -5, -5, 24, 24, 24, 24, 24, 24, 24, 24, -5, -5, -5, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 23, 23, 23, -5, 18, -5,
	-5, 23, 23, 23, 23, 23, 23, 23, -5, -5, -5, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, 19, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, 20, -40, -40, 20, -40, -40, -40, -40, -40, -40, 20, 20, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, 21, -40, -40, 21, -40, -40, -40,
	-40, -40, -40, 21, 21, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, 22,
	-40, -40, 22, -40, -40, -40, -40, -40, -40, 22, 22, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, 23, -40, -40, 23, -40, -40, -40, -40, -40, -40, 23,
	23, -40, -40, -40, -40, -40, -40, -40, -40, -39, -39, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, 23, 23,
	23, -39, 18, -39, -39, 23, 23, 23, 23, 23, 23, 23, -39, -39, -39, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -1, -9, -9, 24, -9, -9, 24, 24,
	24, -9, -9, -9, 24, 24, 24, 24, 24, 24, 24, 24, -9, -9, -9, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, 27, -10, -10, 27, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, 26, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 30, -5, -5, -5, -5, 29,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -8, 29, 29, -8, 29, -8, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, 29, -8,
	-43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -43, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -3, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 33, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, 35, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -5, -5, 41, 41, 41, 41, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, 42, -5, -5, -5, -5, -5, -5, 43, -5, -5, 43, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 44, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 44, 44, 44, 44, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	45, -5, -5, -5, -5, -5, -5, -4, -36, -36, -4, -36, -36, -36, -36, -36, -36,
	-36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36,
	-36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, 54, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 51, 53, 53, 53,
	53, 49, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	53, 53, 53, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, 50, -47,
	-47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47,
	-47, -47, -47, -47, -47, -47, -47, -47, -47, -45, -45, -45, -45, -45, -45,
	-45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45,
	-45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45,
	-45, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47,
	-47, -47, 52, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47,
	-47, -47, -47, -47, -47, -47, -47, -47, -46, -46, -46, -46, -46, -46, -46,
	-46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46,
	-46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46,
	-47, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, -47, 53, 53, 53, 53, -47, 53,
	53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53,
	-44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44,
	-44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44,
	-44, -44, -44, -44, -44, -44, -44,
}

var tmBacktracking = []int{
	4, 12, // in Identifier
	4, 10, // in Identifier
	20, 32, // in '.'
	31, 43, // in multiline
}
