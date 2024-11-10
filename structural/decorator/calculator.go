package main

type MathOperation interface {
	Execute(a, b float32) float32
}

type calculator struct {
	operations []MathOperation
}

func (c *calculator) Execute(a, b float32) float32 {
	for _, o := range c.operations {
		a = o.Execute(a, b)
	}
	return a
}

func NewCalculator(operations []MathOperation) *calculator {
	return &calculator{
		operations: operations,
	}
}

type sum struct{}

func (sum) Execute(a, b float32) float32 {
	return a + b
}

func NewSum() MathOperation {
	return &sum{}
}

type subtract struct{}

func (subtract) Execute(a, b float32) float32 {
	return a - b
}

func NewSubtract() MathOperation {
	return &subtract{}
}

type divide struct{}

func (divide) Execute(a, b float32) float32 {
	return a / b
}

func NewDivide() MathOperation {
	return &divide{}
}

type multiply struct{}

func (multiply) Execute(a, b float32) float32 {
	return a * b
}

func NewMultiply() MathOperation {
	return &multiply{}
}
