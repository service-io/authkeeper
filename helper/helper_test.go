// Package helper
// @author tabuyos
// @since 2023/8/4
// @description helper
package helper

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"testing"
	"time"
)

func TestHelper(t *testing.T) {
	var i int = 123
	var ii *int = &i

	d := *ii

	ii = nil

	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", ii)
}

func TestGenToken(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("nihaofdsafdsafdsa"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(hash))

	fmt.Println(len(string(hash)))
	fmt.Println(len(hash))
	fmt.Println(len("WlUuYXFZZ2UvZW1xZTllNzdhVXRRWGRWJGpZQ0MkaUUwTzhxTk4zOWx3V1NoUDFjJC50eGJBMjkyUnov"))

	fmt.Println(base64.StdEncoding.EncodeToString(hash))
	fmt.Println(base64.StdEncoding.EncodeToString(hash))
	Shuffle(hash)
	fmt.Println(base64.StdEncoding.EncodeToString(hash))
	Shuffle(hash)
	fmt.Println(base64.StdEncoding.EncodeToString(hash))
	token := base64.StdEncoding.EncodeToString(hash)

	decodeString, err := base64.StdEncoding.DecodeString(token)
	fmt.Println(string(decodeString))

	fmt.Println(GenerateToken())
}

const (
	defaultTokenLen int = 128
)

func GenerateToken() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	runes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, defaultTokenLen)

	for i := range b {
		b[i] = runes[r.Intn(len(runes))]
	}
	return string(b)
}
