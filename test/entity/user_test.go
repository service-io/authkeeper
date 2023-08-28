package entity

import (
	"fmt"
	"strings"
	"testing"
)

func TestUser_InsertFields(t *testing.T) {
	u := &User{}
	u.Insert(u.FID(), u.FName())
	fields, ph := u.InsertFields()
	fmt.Println(fields)
	fmt.Println(ph)
}

func TestUser_UpdateFields(t *testing.T) {
	u := &User{}
	u.Update(u.FID(), u.FName())
	fields := u.UpdateFields()
	fmt.Println(fields)
}

func TestUser_SelectFS(t *testing.T) {
	u := &User{}
	u.Select(u.FID(), u.FName())
	fs := u.SelectFS()
	var fields []string
	for _, f := range fs {
		fields = append(fields, f.Name)
	}
	fmt.Println(fields)
}

func TestUser_SelectFields(t *testing.T) {
	u := &User{}
	fields := u.SelectFields()
	fmt.Println(fields)
}

func TestUser_ReturnFields(t *testing.T) {
	u := &User{}
	e := &User{}
	fns := u.ReturnFields()
	fmt.Println(fns)

	var values []any
	for _, fn := range fns {
		values = append(values, fn(e))
	}
	fmt.Println(values)
}

func TestUser_ReturnFields1(t *testing.T) {
	u := &User{}
	e := &User{}
	u.Select(u.FID(), u.FName(), u.FAge())
	fns := u.ReturnFields()

	fmt.Println(fns)

	var values []any
	for _, fn := range fns {
		values = append(values, fn(e))
	}
	fmt.Println(values)
}

func TestUser_Where(t *testing.T) {
	u := &User{}
	u.Select(u.FID(), u.FName(), u.FAge())
	u.WithWhere(Once(OfCond(u.FID().Name, "=", 1), OfCond(u.FName().Name, "=", "tabuyos")))
	u.WithOrder(Asc(u.FName().Name))
	selectFields := u.SelectFields()
	returnFields := u.ReturnFields()
	plain, values := u.Where()
	order := u.Order()
	fmt.Println(plain)
	fmt.Println(values)
	fmt.Println(selectFields)
	fmt.Println(returnFields)

	var sqlBuilder strings.Builder

	sqlBuilder.WriteString("SELECT ")
	sqlBuilder.WriteString(strings.Join(selectFields, ", "))
	sqlBuilder.WriteString(" FROM ")
	sqlBuilder.WriteString(u.TableName())
	sqlBuilder.WriteString(" WHERE ")
	sqlBuilder.WriteString(plain)
	sqlBuilder.WriteString(" ORDER BY ")
	sqlBuilder.WriteString(strings.Join(order, ", "))
	sqlBuilder.WriteString(";")

	fmt.Println(sqlBuilder.String())

}

func TestUser_Order(t *testing.T) {
	u := &User{}
	plain, values := u.Where()
	fmt.Println(plain, values)
}
