package entity

import (
	"fmt"
	"testing"
)

func TestAccount_BEI(t *testing.T) {
	a := &Account{}
	bei := a.BEI()
	bei.Insert(a.FID())
	fmt.Println(bei.InsertFields())
	fmt.Println(a.FAge().Name)
}

func TestAccount_Configure(t *testing.T) {
	a := &Account{}
	a.Configure(
		func(bei *BaseEntity[Account]) {
			bei.Insert(a.FID(), a.FName(), a.FBirthday(), a.FAge())
			bei.WithValues(1, "tabuyos", "2023-11-01", 25)
			bei.WithValues(3, "aaron", "2023-11-01", 27)
		},
	)
	fmt.Println(a.BEI().InsertFields())
	fmt.Println(a.BEI().InsertSql(2))
}
