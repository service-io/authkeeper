package entity

import "strings"

type Mode string

func (p Mode) String() string {
	return string(p)
}

const (
	AndMode Mode = "AND"
	OrMode  Mode = "OR"
	DftMode      = AndMode
)

type Cond struct {
	Col string
	Op  string
	Val any
}

type Pred struct {
	Mode  Mode
	Preds []*Pred
	Conds []*Cond
}

type Order struct {
	Col string
	Asc bool
}

func Desc(col string) *Order {
	return &Order{
		Col: col,
		Asc: false,
	}
}

func Asc(col string) *Order {
	return &Order{
		Col: col,
		Asc: true,
	}
}

func OfCond(col string, op string, val any) *Cond {
	return &Cond{
		Col: col,
		Op:  op,
		Val: val,
	}
}

func Of(p Mode, s []*Pred, c []*Cond) *Pred {
	return &Pred{
		Mode:  p,
		Preds: s,
		Conds: c,
	}
}

func And(ps ...*Pred) *Pred {
	return Of(AndMode, ps, nil)
}

func Or(ps ...*Pred) *Pred {
	return Of(OrMode, ps, nil)
}

func Once(c ...*Cond) *Pred {
	return Of(DftMode, nil, c)
}

func (pred *Pred) Render() (string, []any) {
	if pred == nil {
		return "", nil
	}
	plain, values := pred.parsePred()
	//if strings.HasPrefix(plain, "(") {
	//	plain = plain[1 : len(plain)-1]
	//}
	return plain, values
}

func (pred *Pred) parsePred() (string, []any) {
	pl := len(pred.Preds)
	if pl == 0 {
		cl := len(pred.Conds)
		if cl == 0 {
			return "", nil
		}
		var sb = make([]string, cl)
		var values []any
		for i, cond := range pred.Conds {
			if cond.Val != nil {
				values = append(values, cond.Val)
			}
			s := cond.Col + " " + cond.Op + " ?"
			sb[i] = s
		}
		return strings.Join(sb, " "+pred.Mode.String()+" "), values
	}
	var sb = make([]string, pl)
	var values []any
	for i, p := range pred.Preds {
		s, v := p.parsePred()
		if len(v) != 0 {
			values = append(values, v...)
		}
		sb[i] = "(" + s + ")"
	}
	return strings.Join(sb, " "+pred.Mode.String()+" "), values
}
