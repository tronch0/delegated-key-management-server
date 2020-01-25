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

	return crypto.Mul(pX, pY, c.r)
}

func (c *Client) Unhide(x *big.Int, y *big.Int) (newX *big.Int, newY *big.Int) {
	rInverse := crypto.CalculateInverse(c.r)

	return crypto.Mul(x, y, rInverse)
}
