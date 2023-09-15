package constant

type Symbolic string

const (
	Space            Symbolic = " "
	CommaSpace       Symbolic = ", "
	Dot              Symbolic = "."
	Equal            Symbolic = "="
	LeftParentheses  Symbolic = "("
	RightParentheses Symbolic = ")"
)

func (s Symbolic) Pretty() string {
	return Space.Literal() + s.Literal() + Space.Literal()
}

func (s Symbolic) Literal() string {
	return string(s)
}
