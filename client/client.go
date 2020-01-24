package client

import (
	"../crypto"
	"math/big"
)

type Client struct {
	r *big.Int
}

func New() (*Client, error) {
	r, err := crypto.GenerateR()
	if err != nil {
		return nil, err
	}

	return &Client{r: r}, nil
}

func (c *Client) Hide(data []byte) (x *big.Int, y *big.Int) {
	pX, pY := crypto.HashIntoPoint(data)

	return crypto.Exp(pX, pY, c.r.Bytes())
}

func (c *Client) Unhide(x *big.Int, y *big.Int) (newX *big.Int, newY *big.Int) {
	rInverse := crypto.CalculateInverse(c.r)

	return crypto.Exp(x, y, rInverse.Bytes())
}
