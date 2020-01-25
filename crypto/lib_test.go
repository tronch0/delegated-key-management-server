package crypto

import (
	"math/big"
	"strings"
	"testing"
)

var (
	exRecord = "Naomi B. Mcbride , Id-36782375"
)

func TestExp(t *testing.T) {
	baseX1, baseY1 := getBasePointOnCurve()

	k, _ := new(big.Int).SetString("115792089210356248762697446949407573529996955224135760342422259061068512044352", 10)

	newX, newY := Mul(baseX1, baseY1, k)

	if strings.EqualFold(newX.String(), "47776904C0F1CC3A9C0984B66F75301A5FA68678F0D64AF8BA1ABCE34738A73E") {
		t.Fatal()
	}
	if strings.EqualFold(newY.String(), "55FFA1184A46A8D89DCE7A9A889B717C7E4D7FBCD72A8CC0CD0878008E0E0323") {
		t.Fatal()
	}
}

func TestDelegatorHappyFlow(t *testing.T) {
	secret := new(big.Int)
	secret.SetString("115792089210356248762697446949407573529996955224135760342422259061068512044352", 10)

	p1x, p1y := getBasePointOnCurve()

	_, _ = Mul(p1x, p1y, secret)

}

func TestClientHappyFlow(t *testing.T) {
	secret, err := GenerateR()
	if err != nil {
		t.Fatal()
	}

	recordB := []byte(exRecord)

	x, y := HashIntoPoint(recordB)

	newX, newY := Mul(x, y, secret)

	oldX, oldY := MulWithInverseK(newX, newY, secret)

	if x.Cmp(oldX) != 0 {
		t.Fatal()
	}

	if y.Cmp(oldY) != 0 {
		t.Fatal()
	}

}

func TestFullHappyFlow(t *testing.T) {
	clientSecret, err := GenerateR()
	if err != nil {
		t.Fatal()
	}

	recordB := []byte(exRecord)

	x, y := HashIntoPoint(recordB)

	xWithClientSecret, yWithClientSecret := Mul(x, y, clientSecret)

	delegatorSecret, err := GenerateR()
	if err != nil {
		t.Fatal()
	}
	xWithClientAndDelegatorSecrets, yWithClientAndDelegatorSecrets := Mul(xWithClientSecret, yWithClientSecret, delegatorSecret)

	xOriginal, yOriginal := MulWithInverseK(xWithClientSecret, yWithClientSecret, clientSecret)

	if x.Cmp(xOriginal) != 0 {
		t.Fatal()
	}

	if y.Cmp(yOriginal) != 0 {
		t.Fatal()
	}

	xWithDelegatorSecret, yWithDelegatorSecret := MulWithInverseK(xWithClientAndDelegatorSecrets, yWithClientAndDelegatorSecrets, clientSecret)

	xDelegator, yDelegator := Mul(x, y, delegatorSecret)

	if xDelegator.Cmp(xWithDelegatorSecret) != 0 || yDelegator.Cmp(yWithDelegatorSecret) != 0 {
		t.Fatal()
	}

}

func getBasePointOnCurve() (x, y *big.Int) {
	x = c.Params().Gx
	y = c.Params().Gy
	return
}
