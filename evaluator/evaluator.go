// Package evaluator provides types and functions related to
// the evaluation of golisp expressions.
package evaluator

// An Env represents an association list.
type Env map[any]Expr

// ExprType represents the type of an expression.
type ExprType string

// These constants define the possible expression types.
const (
	Symbol ExprType = "symbol"
	Number ExprType = "number"
	String ExprType = "string"
	Char   ExprType = "char"
	Bool   ExprType = "bool"
	Vector ExprType = "vector"
)

// An Expr represents an expression, which may represent a deeply
// nested series of additional expressions.
type Expr struct {
	Type     ExprType
	Operator any
	Operand1 *Expr
	Operand2 *Expr
}

// unknown is a special type used when no result can be returned.
var unknown = Expr{}
