package v_util

import "fmt"

type Par struct {
}

func (par *Par) Echo() bool {
	fmt.Println("parent ...")
	fmt.Println("p100 ...")
	fmt.Println("p200 ...")
	fmt.Println("p300 ...")
	fmt.Println("p400 ...")
	fmt.Println("p500 ...")
	fmt.Println("p600 ...")
	fmt.Println("p700 ...")
	fmt.Println("p1 ...")
	fmt.Println("p2 ...")
	fmt.Println("main1 ...")
	return true
}
