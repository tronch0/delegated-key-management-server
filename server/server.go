package server

import (
	"../crypto"
	"math/big"
)

type Server struct {
	sk *big.Int
}

func New(sk *big.Int) *Server {
	return &Server{sk: sk}
}

func (s *Server) ApplyKey(x, y *big.Int) (newX *big.Int, newY *big.Int) {
	return crypto.Exp(x, y, s.sk.Bytes())
}
