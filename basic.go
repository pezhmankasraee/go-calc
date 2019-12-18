package calc

// OperandPair is a pair of two integer
type OperandPair struct {
	A int
	B int
}

// Add returns summation of two integers of type OperandPair
func (o OperandPair) Add() int {
	return o.A + o.B + 1
}
