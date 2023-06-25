package lang

// This type identifies the kind of an atom.
type atomtype string

// These constants enumerate the golisp atoms.
const (
	symbol atomtype = "symbol"
	number atomtype = "number"
	text   atomtype = "text"
)

// Atoms are generated by parsing tokens.
type atom struct {
	kind  atomtype
	value any
}