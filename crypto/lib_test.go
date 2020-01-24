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

	k := new(big.Int)
	k.SetString("115792089210356248762697446949407573529996955224135760342422259061068512044352", 10)
	newX, newY := Exp(baseX1, baseY1, k.Bytes())

	if strings.EqualFold(newX.String(), "47776904C0F1CC3A9C0984B66F75301A5FA68678F0D64AF8BA1ABCE34738A73E") {
		t.Fatal()
	}
	if strings.EqualFold(newY.String(), "55FFA1184A46A8D89DCE7A9A889B717C7E4D7FBCD72A8CC0CD0878008E0E0323") {
		t.Fatal()
	}
}

func TestDelegatorHappyFlow(t *testing.T) {
	// gets requests contains strings of hashes (represents points)
	// convert to bytes
	// mount to big.int x,y
	// START

	//k = 115792089210356248762697446949407573529996955224135760342422259061068512044352
	//x = 47776904C0F1CC3A9C0984B66F75301A5FA68678F0D64AF8BA1ABCE34738A73E
	//y = 55FFA1184A46A8D89DCE7A9A889B717C7E4D7FBCD72A8CC0CD0878008E0E0323

	secret := new(big.Int)
	secret.SetString("115792089210356248762697446949407573529996955224135760342422259061068512044352", 10)

	p1x, p1y := getBasePointOnCurve()

	// raise the point in exp(secret)
	_, _ = Exp(p1x, p1y, secret.Bytes())

	// END
	// serialize result to bytes
	// bytes to string
	// return response
}

func TestClientHappyFlow(t *testing.T) {
	// gets requests contains records
	// START
	secret, err := GenerateR()
	if err != nil {
		t.Fatal()
	}

	recordB := []byte(exRecord)

	// transform records to hashes (points on the curve (hash to group))
	x, y := HashIntoPoint(recordB)

	// raise the point in exp(secret)
	newX, newY := Exp(x, y, secret.Bytes())

	// send result to delegator

	// raise the result of the delegator to exp(1/secret)
	oldX, oldY := InverseK(newX, newY, secret)

	if x.Cmp(oldX) != 0 {
		t.Fatal()
	}

	if y.Cmp(oldY) != 0 {
		t.Fatal()
	}

	// great success!
}

func TestFullHappyFlow(t *testing.T) {
	clientSecret, err := GenerateR()
	if err != nil {
		t.Fatal()
	}

	recordB := []byte(exRecord)

	// transform records to hashes (points on the curve (hash to group))
	x, y := HashIntoPoint(recordB)
	//fmt.Printf("Record: %s\n",exRecord)
	//fmt.Printf("Record bytes: %x\n",recordB)
	//fmt.Printf("Record's point: x=%s, y=%s \n",x.String(),y.String())

	// raise the point in exp(clientSecret)
	xWithClientSecret, yWithClientSecret := Exp(x, y, clientSecret.Bytes())

	// send result to delegator
	delegatorSecret, err := GenerateR()
	if err != nil {
		t.Fatal()
	}
	xWithClientAndDelegatorSecrets, yWithClientAndDelegatorSecrets := Exp(xWithClientSecret, yWithClientSecret, delegatorSecret.Bytes())

	// raise the result of the delegator to exp(1/clientSecret)
	xOriginal, yOriginal := InverseK(xWithClientSecret, yWithClientSecret, clientSecret)

	if x.Cmp(xOriginal) != 0 {
		t.Fatal()
	}

	if y.Cmp(yOriginal) != 0 {
		t.Fatal()
	}

	xWithDelegatorSecret, yWithDelegatorSecret := InverseK(xWithClientAndDelegatorSecrets, yWithClientAndDelegatorSecrets, clientSecret)

	xDelegator, yDelegator := Exp(x, y, delegatorSecret.Bytes())

	if xDelegator.Cmp(xWithDelegatorSecret) != 0 || yDelegator.Cmp(yWithDelegatorSecret) != 0 {
		t.Fatal()
	}

}

func getBasePointOnCurve() (x, y *big.Int) {
	x = c.Params().Gx
	y = c.Params().Gy
	return
}