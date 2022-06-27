package v_util

import "fmt"

type Chi struct {
	Par
}

func (chi *Chi) Chi_echo() {
	fmt.Println("chile ...")
	fmt.Println(chi.Echo())
	fmt.Println("c1")
	fmt.Println("c100")
	fmt.Println("c200")
}
