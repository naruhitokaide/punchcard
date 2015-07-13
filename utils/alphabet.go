package utils

import (
	"errors"
)

var A = [][]int{[]int{0, 1, 1, 0}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 1, 1, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}}
var B = [][]int{[]int{1, 1, 1, 0}, []int{1, 0, 0, 1}, []int{1, 1, 1, 0}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 1, 1, 0}}
var C = [][]int{[]int{0, 1, 1, 0}, []int{1, 0, 0, 1}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}, []int{1, 0, 0, 1}, []int{0, 1, 1, 0}}
var D = [][]int{[]int{1, 1, 1, 0}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 1, 1, 0}}
var E = [][]int{[]int{1, 1, 1, 1}, []int{1, 0, 0, 0}, []int{1, 1, 1, 0}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}, []int{1, 1, 1, 1}}
var F = [][]int{[]int{1, 1, 1, 1}, []int{1, 0, 0, 0}, []int{1, 1, 1, 0}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}}
var G = [][]int{[]int{1, 1, 1, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 0}, []int{1, 0, 1, 1}, []int{1, 1, 1, 1}}
var H = [][]int{[]int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 1, 1, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}}
var I = [][]int{[]int{1}, []int{1}, []int{1}, []int{1}, []int{1}, []int{1}}
var J = [][]int{[]int{0, 0, 0, 1}, []int{0, 0, 0, 1}, []int{0, 0, 0, 1}, []int{0, 0, 0, 1}, []int{1, 0, 0, 1}, []int{0, 1, 1, 0}}
var K = [][]int{[]int{1, 0, 0, 1}, []int{1, 0, 1, 1}, []int{1, 1, 0, 0}, []int{1, 0, 1, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}}
var L = [][]int{[]int{1, 0, 0, 0}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}, []int{1, 1, 1, 1}}
var M = [][]int{[]int{1, 0, 0, 0, 1}, []int{1, 1, 0, 1, 1}, []int{1, 0, 1, 0, 1}, []int{1, 0, 1, 0, 1}, []int{1, 0, 0, 0, 1}, []int{1, 0, 0, 0, 1}}
var N = [][]int{[]int{1, 1, 0, 0, 1}, []int{1, 1, 0, 0, 1}, []int{1, 0, 1, 0, 1}, []int{1, 0, 1, 0, 1}, []int{1, 0, 0, 1, 1}, []int{1, 0, 0, 1, 1}}
var O = [][]int{[]int{1, 1, 1, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 1, 1, 1}}
var P = [][]int{[]int{1, 1, 1, 1}, []int{1, 0, 0, 1}, []int{1, 1, 1, 1}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}}
var Q = [][]int{[]int{0, 1, 1, 1, 0}, []int{1, 0, 0, 0, 1}, []int{1, 0, 0, 0, 1}, []int{1, 0, 1, 0, 1}, []int{1, 0, 0, 1, 1}, []int{0, 1, 1, 1, 0}}
var R = [][]int{[]int{1, 1, 1, 0}, []int{1, 0, 0, 1}, []int{1, 1, 1, 0}, []int{1, 1, 0, 0}, []int{1, 0, 1, 0}, []int{1, 0, 0, 1}}
var S = [][]int{[]int{1, 1, 1, 1}, []int{1, 0, 0, 0}, []int{1, 1, 1, 1}, []int{0, 0, 0, 1}, []int{0, 0, 0, 1}, []int{1, 1, 1, 1}}
var T = [][]int{[]int{1, 1, 1, 1, 1}, []int{0, 0, 1, 0, 0}, []int{0, 0, 1, 0, 0}, []int{0, 0, 1, 0, 0}, []int{0, 0, 1, 0, 0}, []int{0, 0, 1, 0, 0}}
var U = [][]int{[]int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{1, 0, 0, 1}, []int{0, 1, 1, 0}}
var V = [][]int{[]int{1, 0, 0, 0, 1}, []int{1, 0, 0, 0, 1}, []int{1, 0, 0, 0, 1}, []int{0, 1, 0, 1, 0}, []int{0, 1, 0, 1, 0}, []int{0, 0, 1, 0, 0}}
var W = [][]int{[]int{1, 0, 0, 0, 1}, []int{1, 0, 0, 0, 1}, []int{1, 0, 1, 0, 1}, []int{1, 0, 1, 0, 1}, []int{1, 1, 1, 1, 1}, []int{0, 1, 0, 1, 0}}
var X = [][]int{[]int{1, 0, 0, 0, 1}, []int{0, 1, 0, 1, 0}, []int{0, 0, 1, 0, 0}, []int{0, 0, 1, 0, 0}, []int{0, 1, 0, 1, 0}, []int{1, 0, 0, 0, 1}}
var Y = [][]int{[]int{1, 0, 0, 0, 1}, []int{0, 1, 0, 1, 0}, []int{0, 0, 1, 0, 0}, []int{0, 0, 1, 0, 0}, []int{0, 0, 1, 0, 0}, []int{0, 0, 1, 0, 0}}
var Z = [][]int{[]int{1, 1, 1, 1}, []int{0, 0, 0, 1}, []int{0, 0, 1, 0}, []int{0, 1, 0, 0}, []int{1, 0, 0, 0}, []int{1, 1, 1, 1}}

var SPACE = [][]int{[]int{0}, []int{0}, []int{0}, []int{0}, []int{0}, []int{0}}

// TranslateLetter returns the bitmap like [][]int of the given letter.
func TranslateLetter(letter string) ([][]int, error) {
	var err error
	switch {
	case letter == "a":
		return A, err
	case letter == "b":
		return B, err
	case letter == "c":
		return C, err
	case letter == "d":
		return D, err
	case letter == "e":
		return E, err
	case letter == "f":
		return F, err
	case letter == "g":
		return G, err
	case letter == "h":
		return H, err
	case letter == "i":
		return I, err
	case letter == "j":
		return J, err
	case letter == "k":
		return K, err
	case letter == "l":
		return L, err
	case letter == "m":
		return M, err
	case letter == "n":
		return N, err
	case letter == "o":
		return O, err
	case letter == "p":
		return P, err
	case letter == "q":
		return Q, err
	case letter == "r":
		return R, err
	case letter == "s":
		return S, err
	case letter == "t":
		return T, err
	case letter == "u":
		return U, err
	case letter == "v":
		return V, err
	case letter == "w":
		return W, err
	case letter == "x":
		return X, err
	case letter == "y":
		return Y, err
	case letter == "z":
		return Z, err
	case letter == " ":
		return SPACE, err
	default:
		return nil, errors.New("The letter must be lowercase or a space: [a-z ]{1}")
	}
}
