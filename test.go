package main

import (
	"fmt"

	"github.com/myxtype/filecoin-client/util"
)

func main1() {
	testStr := "t1otzxmzqqrlnpfzczhb4tx6c4glwzlozqx5lalja"
	enstr := util.Encrypt(testStr)
	fmt.Println(enstr)
	destr := util.Decrypt(enstr)
	fmt.Println(destr)
	fmt.Println("==========")
	username := util.Encrypt("admin")
	fmt.Println(username)
	fmt.Println("==========")
	password := util.Encrypt("1234567")
	fmt.Println(password)

}
