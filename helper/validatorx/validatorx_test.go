// Package validatorx
// @author tabuyos
// @since 2023/8/24
// @description validatorx
package validatorx

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	fmt.Printf(time.Now().Format("2006-01-02 15:04:05"))
}

func TestValidator_WithRequiredAndPanic(t *testing.T) {
	var rr = RegisterRequest{
		Username: "a",
		// Nickname: "a",
		Email:    "a",
		Password: "a",
		Age:      12,
	}
	validatorx := NewValidator()

	validatorx.WithRequiredAndPanic(rr.Username, "名称")
	validatorx.WithRequiredAndPanic(rr.Nickname, "昵称")
	validatorx.WithRequiredAndPanic(rr.Email, "邮箱")
	validatorx.WithRequiredAndPanic(rr.Password, "密码")

}

func TestValidator_WithRequired(t *testing.T) {
	var rr = RegisterRequest{
		Username: "a",
		Nickname: "a",
		Email:    "a",
		Password: "a",
		Age:      12,
	}
	validatorx := NewValidator()

	err := errors.Join(
		validatorx.WithRequired(rr.Username, "名称"),
		validatorx.WithRequired(rr.Nickname, "昵称"),
		validatorx.WithRequired(rr.Email, "邮箱"),
		validatorx.WithRequired(rr.Password, "密码"),
	)

	if err != nil {
		fmt.Println(err)
	}

}

type name string

func (n *name) String() string {
	return "tabuyos"
}

func TestSwitch(t *testing.T) {
	var n any = name("ok")
	switch v := n.(type) {
	case string:
		fmt.Println(v)
	default:
		fmt.Println("not ok")
	}
}

func TestValidateRequiredField(t *testing.T) {

	var rr = RegisterRequest{
		Username: "a",
		Nickname: "a",
		// Password: "a",
		Email: "a",
		Age:   12,
	}
	errs := ValidateRequiredField(rr.Username, rr.Password, rr.Email, rr.Nickname)
	if errs != nil {
		for _, err := range errs {
			validationErrors, ok := err.(validator.ValidationErrors)
			if !ok {
				continue
			}

			for _, ve := range validationErrors {
				// 列出效验出错字段的信息
				fmt.Println("Namespace: ", ve.Namespace())
				fmt.Println("Field: ", ve.Field())
				fmt.Println("StructNamespace: ", ve.StructNamespace())
				fmt.Println("StructField: ", ve.StructField())
				fmt.Println("Tag: ", ve.Tag())
				fmt.Println("ActualTag: ", ve.ActualTag())
				fmt.Println("Kind: ", ve.Kind())
				fmt.Println("Type: ", ve.Type())
				fmt.Println("Value: ", ve.Value())
				fmt.Println("Param: ", ve.Param())
				fmt.Println()
			}
		}
	} else {
		fmt.Println("ok")
	}

}

func TestSingleField(t *testing.T) {
	var rr RegisterRequest = RegisterRequest{
		Username: "a",
		Nickname: "a",
		// Password: "a",
		Email: "a",
		Age:   12,
	}
	validatorx := NewValidator()
	err2 := ValidateRequiredField(rr.Username, rr.Password, rr.Email, rr.Nickname)
	if err2 != nil {
		fmt.Println(reflect.TypeOf(err2))
		fmt.Println(err2)
	} else {
		fmt.Println("ok")
	}

	var boolTest bool
	err := validatorx.Var(boolTest, "required")
	if err != nil {
		fmt.Println(err)
	}
	stringTest := ""
	err = validatorx.Var(stringTest, "required")
	if err != nil {
		fmt.Println(err)
	}

	err = validatorx.Var(nil, "required")
	if err != nil {
		fmt.Println(err)
	}

	emailTest := "test@126.com"
	err = validatorx.Var(emailTest, "email")
	if err != nil {
		fmt.Println(err)
	} else {
		// 输出： success。 说明验证成功
		fmt.Println("success")
	}

	emailTest2 := "test.126.com"
	errs := validatorx.Var(emailTest2, "required,email")
	if errs != nil {
		// 输出: Key: "" Error:Field validation for "" failed on the "email" tag。验证失败
		fmt.Println(errs)
	}
}

type User struct {
	FirstName string     `validate:"required"`
	LastName  string     `validate:"required"`
	Age       uint8      `validate:"gte=0,lte=130"`
	Email     string     `validate:"required,email"`
	Addresses []*Address `validate:"required,dive,required"`
}

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func TestStruct(t *testing.T) {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName: "Badger",
		LastName:  "Smith",
		Age:       135,
		Email:     "Badger.Smith@gmail.com",
		Addresses: []*Address{address},
	}

	validatorx := NewValidator()
	err := validatorx.Struct(user)
	if err != nil {
		fmt.Println("=== error msg ====")
		fmt.Println(err)

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			fmt.Println(err)
			return
		}

		fmt.Println("\r\n=========== error field info ====================")
		for _, err := range err.(validator.ValidationErrors) {
			// 列出效验出错字段的信息
			fmt.Println("Namespace: ", err.Namespace())
			fmt.Println("Field: ", err.Field())
			fmt.Println("StructNamespace: ", err.StructNamespace())
			fmt.Println("StructField: ", err.StructField())
			fmt.Println("Tag: ", err.Tag())
			fmt.Println("ActualTag: ", err.ActualTag())
			fmt.Println("Kind: ", err.Kind())
			fmt.Println("Type: ", err.Type())
			fmt.Println("Value: ", err.Value())
			fmt.Println("Param: ", err.Param())
			fmt.Println()
		}
	}
}

func TestSlice(t *testing.T) {
	sliceOne := []string{"123", "onetwothree", "myslicetest", "four", "five"}

	validatorx := NewValidator()
	err := validatorx.Var(sliceOne, "max=15,dive,min=4")
	if err != nil {
		fmt.Println(err)
	}

	var sliceTwo []string
	err = validatorx.Var(sliceTwo, "min=4,dive,required")
	if err != nil {
		fmt.Println(err)
	}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Age      uint8  `json:"age" binding:"gte=1,lte=120"`
}

func TestGin(t *testing.T) {

	router := gin.Default()
	router.POST("register", Register)
	// err := router.Run(":9999")
	// if err != nil {
	// 	return
	// }
}

func Register(c *gin.Context) {
	var r RegisterRequest
	err := c.ShouldBindJSON(&r)
	if err != nil {
		of := reflect.TypeOf(err)
		fmt.Println(of.String())

		fmt.Println("register failed")
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	// 验证 存储操作省略.....
	fmt.Println("register success")
	c.JSON(http.StatusOK, "successful")
}
