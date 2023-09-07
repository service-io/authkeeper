// Package keyword
// @author tabuyos
// @since 2023/9/6
// @description beiold
package keyword

import "strings"

type Keyword string

var uppercase = true

func RegistryCase(uc bool) {
	uppercase = uc
}

func Of(k string) *Keyword {
	var kw = Keyword(k)
	return &kw
}

func (k Keyword) String() string {
	return k.Literal()
}

func (k Keyword) Literal() string {
	if len(k) == 0 {
		return ""
	}
	if uppercase {
		return strings.ToUpper(string(k))
	}
	return strings.ToLower(string(k))
}

func (k Keyword) LowerCase() *Keyword {
	return Of(strings.ToLower(k.String()))
}

func (k Keyword) UpperCase() *Keyword {
	return Of(strings.ToUpper(k.String()))
}

const (
	As       Keyword = "AS"
	Join     Keyword = "JOIN"
	Left     Keyword = "LEFT"
	Right    Keyword = "RIGHT"
	Inner    Keyword = "INNER"
	Outer    Keyword = "OUTER"
	Distinct Keyword = "DISTINCT"
	From     Keyword = "FROM"
	Select   Keyword = "SELECT"
	Delete   Keyword = "DELETE"
	Update   Keyword = "UPDATE"
	Insert   Keyword = "INSERT"
	On       Keyword = "ON"
	Set      Keyword = "SET"
	Into     Keyword = "INTO"
	Value    Keyword = "VALUE"
	Values   Keyword = "VALUES"
	Asterisk Keyword = "ASTERISK"
	Where    Keyword = "WHERE"
	And      Keyword = "AND"
	Or       Keyword = "OR"
	Group    Keyword = "GROUP"
	Having   Keyword = "HAVING"
	Order    Keyword = "ORDER"
	By       Keyword = "BY"
	Asc      Keyword = "ASC"
	Desc     Keyword = "DESC"
	Limit    Keyword = "LIMIT"
	Like     Keyword = "Like"
	Offset   Keyword = "OFFSET"
	Count    Keyword = "COUNT"
	Is       Keyword = "IS"
	Not      Keyword = "NOT"
	Null     Keyword = "NULL"
)
