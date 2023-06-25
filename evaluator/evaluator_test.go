package evaluator

import (
	"fmt"
	"testing"
)

func TestEvaluate(t *testing.T) {
	exp := Expr{
		Type:     Symbol,
		Operator: "'test",
	}

	env := Env{
		"'test": Expr{
			Type:     Number,
			Operator: 123,
		},
	}

	res, err := Evaluate(exp, env)
	fmt.Println(res)
	fmt.Println(err)
}
