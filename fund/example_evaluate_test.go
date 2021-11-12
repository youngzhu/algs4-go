package fund_test

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/youngzhu/algs4-go/fund"
)

// Arithmetic expression evaluation.
// Evaluate() is a stack client that evaluates fully parenthesized arithmetic
// expressions. It uses Dijkstra's 2-stack algorithm:
// - Push operands onto the operand stack
// - Push operators onto the operator stack
// - Ignore left parentheses
// - On encountering a right parenthesis, pop an operator, pop the requisite 
//   number of operands, and push onto the operand stack the result of applying
//   that operator to those operands.

func Evaluate(exp string) float64 {
	operators := fund.NewStack()
	operands := fund.NewStack()

	ss := strings.Fields(exp)

	for _, s := range ss {
		// fmt.Println(s)

		if s == "(" {
			// ignore
		} else if isOperator(s) {
			operators.Push(s)
		} else if s == ")" {
			op := operators.Pop()
			v := operands.Pop().(float64)

			if op == "+" {
				v = operands.Pop().(float64) + v
			} else if op == "-" {
				v = operands.Pop().(float64) - v
			} else if op == "*" {
				v = operands.Pop().(float64) * v
			} else if op == "/" {
				v = operands.Pop().(float64) / v
			}

			operands.Push(v)

		} else {
			operand, err := strconv.ParseFloat(s, 64)
			if err != nil {
				panic("invalid operand: " + s)
			}
			operands.Push(operand)
		}
	}

	return operands.Pop().(float64)
}

func isOperator(s string) bool {
	switch s {
	case "+", "-", "*", "/":
		return true
	}
	return false
}

func ExampleEvaluate() {
	exp := "( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )"

	fmt.Printf("%.1f", Evaluate(exp))

	// Output:
	// 101.0
}