package keyword

import (
	"metis/model/3.0/iris/internal/constant"
	"strings"
)

type Keyword string

var lowercase bool

func RegistryCase(lc bool) {
	lowercase = lc
}

func Of(k string) Keyword {
	return Keyword(k)
}

func (k Keyword) String() string {
	return k.Literal()
}

func (k Keyword) Pretty() string {
	return constant.Space.Literal() + k.Literal() + constant.Space.Literal()
}

func (k Keyword) Literal() string {
	if lowercase {
		return strings.ToLower(string(k))
	}
	return string(k)
}

func (k Keyword) LowerCase() Keyword {
	return Of(strings.ToLower(k.String()))
}

func (k Keyword) UpperCase() Keyword {
	return k
}

const (
	As        Keyword = "AS"
	Join      Keyword = "JOIN"
	LeftJoin  Keyword = "LEFT JOIN"
	RightJoin Keyword = "RIGHT JOIN"
	Left      Keyword = "LEFT"
	Right     Keyword = "RIGHT"
	Inner     Keyword = "INNER"
	Outer     Keyword = "OUTER"
	Distinct  Keyword = "DISTINCT"
	From      Keyword = "FROM"
	Select    Keyword = "SELECT"
	Delete    Keyword = "DELETE"
	Update    Keyword = "UPDATE"
	Insert    Keyword = "INSERT"
	On        Keyword = "ON"
	Set       Keyword = "SET"
	Into      Keyword = "INTO"
	Value     Keyword = "VALUE"
	Values    Keyword = "VALUES"
	Asterisk  Keyword = "ASTERISK"
	Where     Keyword = "WHERE"
	And       Keyword = "AND"
	Or        Keyword = "OR"
	Group     Keyword = "GROUP"
	GroupBy   Keyword = "GROUP BY"
	Having    Keyword = "HAVING"
	Order     Keyword = "ORDER"
	OrderBy   Keyword = "ORDER BY"
	By        Keyword = "BY"
	Asc       Keyword = "ASC"
	Desc      Keyword = "DESC"
	Limit     Keyword = "LIMIT"
	Like      Keyword = "Like"
	Offset    Keyword = "OFFSET"
	Count     Keyword = "COUNT"
	Is        Keyword = "IS"
	Not       Keyword = "NOT"
	Null      Keyword = "NULL"
)
