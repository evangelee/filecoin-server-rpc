package main

import (
	"fmt"

	"github.com/myxtype/filecoin-client/util"
)

func main() {

	util.GenerateKey()

	testStr := "f1g6ckkx7pbj2g6wnmns55ubz5jlexvq6eksw7apa"
	enstr := util.Encrypt(testStr)
	fmt.Println(enstr)
	fmt.Println("==========")
	destr := util.Decrypt(enstr)
	fmt.Println(destr)
	fmt.Println("==========")
	username := util.Encrypt("admin")
	fmt.Println(username)
	fmt.Println("==========")
	password := util.Encrypt("1234567")
	fmt.Println(password)

}
