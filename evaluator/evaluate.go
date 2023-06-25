package evaluator

import "fmt"

// Evaluate accepts an expression as well as an association list
// representing the environment, then handles the operation(s).
func Evaluate(exp Expr, env Env) (Expr, error) {
	if isAtom(exp) {
		switch exp.Type {
		case Symbol:
			return lookup(exp.Operator, env)
		case Number, String, Char, Bool, Vector:
			return exp, nil
		default:
			return exp, fmt.Errorf("cannot evaluate %v", exp)
		}
	} else {
		// ...
	}
	return unknown, nil
}

// isAtom returns true if the expression represents an atom.
func isAtom(exp Expr) bool {
	return exp.Operand1 == nil && exp.Operand2 == nil
}

// lookup attempts to resolve an identifier with its binding.
func lookup(id any, env Env) (Expr, error) {
	if exp, ok := env[id]; ok {
		return exp, nil
	}
	return unknown, fmt.Errorf("no such binding: %v", id)
}
