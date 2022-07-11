package main

import (
	"fmt"

	"github.com/myxtype/filecoin-client/util"
)

func main1() {

	util.GenerateKey()

	enstr := util.Encrypt("f1yg5adjeoao4avmnn6msze2ym7uvcnoxx63unszi")
	fmt.Println(enstr)
	fmt.Println("==========", util.Decrypt(enstr))
	enstr1 := util.Encrypt("f1g6ckkx7pbj2g6wnmns55ubz5jlexvq6eksw7apa")
	fmt.Println(enstr1)
	fmt.Println("==========", util.Decrypt(enstr1))
	enstr2 := util.Encrypt("0.01")
	fmt.Println(enstr2)
	fmt.Println("==========", util.Decrypt(enstr2))

}
