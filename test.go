package main

import (
	"fmt"

	"github.com/myxtype/filecoin-client/util"
)

func main1() {

	util.GenerateKey()

	enstr := util.Encrypt("f1brsbja6ew5fexkycgh2p7q4sd4htib2hbvnedei")
	fmt.Println(enstr)
	fmt.Println("==========", util.Decrypt(enstr))
	enstr1 := util.Encrypt("w200207231631")
	fmt.Println(enstr1)
	fmt.Println("==========", util.Decrypt(enstr1))
	enstr2 := util.Encrypt("0.002")
	fmt.Println(enstr2)
	fmt.Println("==========", util.Decrypt(enstr2))

}
