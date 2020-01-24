package crypto

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

/*
	Delegator:
	- Exp()  (util: string -> byte -> big.int)


	Client:
	- HashToGroup()
	- Exp() - r & 1/r
	- Inverse (from r to 1/r)

*/

// Warning: operations on big.Ints are not constant-time: do not use them
// for cryptography unless you're sure this is not an issue.

var (
	// y^2 = x^3-3x+41058363725152142129326129780047268409114441015993725554835256314039467401291
	c = elliptic.P256()
)

func Exp(x, y *big.Int, k []byte) (newX, newY *big.Int) {
	newX, newY = c.ScalarMult(x, y, k)
	return
}

func InverseK(x, y *big.Int, k *big.Int) (newX, newY *big.Int) {
	inverse := fermatInverse(k, c.Params().N)
	return Exp(x, y, inverse.Bytes())
}

func CalculateInverse(k *big.Int) (kInverse *big.Int) {
	return fermatInverse(k, c.Params().N)
}

func fermatInverse(k, N *big.Int) *big.Int {
	two := big.NewInt(2)
	nMinus2 := new(big.Int).Sub(N, two)
	return new(big.Int).Exp(k, nMinus2, N)
}

func HashIntoPoint(priImage []byte) (x, y *big.Int) {
	t := make([]byte, 32)
	copy(t, priImage)

	for {
		potinalX := computeX(t)
		isOnCurve, potinalY := computeY(potinalX)
		if isOnCurve == true {
			return potinalX, potinalY
		}

		increment(t)
	}
}

func GenerateR() (k *big.Int, err error) {
	params := c.Params()
	k, err = rand.Int(rand.Reader, params.N)
	return
}

func computeX(priImage []byte) *big.Int {
	hash := sha256.Sum256(priImage)
	return new(big.Int).SetBytes(hash[:])
}

func computeY(x *big.Int) (isPointOnCurve bool, y *big.Int) {
	// y² = x³ - 3x + b
	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)

	threeX := new(big.Int).Lsh(x, 1)
	threeX.Add(threeX, x)

	x3.Sub(x3, threeX)
	x3.Add(x3, c.Params().B)

	y = x3.ModSqrt(x3, c.Params().P)
	return y != nil, y
}

func increment(counter []byte) {
	for i := len(counter) - 1; i >= 0; i-- {
		counter[i]++
		if counter[i] != 0 {
			break
		}
	}
}
