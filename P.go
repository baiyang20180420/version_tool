package v_util

import "fmt"

type Par struct {
}

func (par *Par) Echo() bool {
	fmt.Println("parent ...")
	fmt.Println("p1 ...")
	return true
}
