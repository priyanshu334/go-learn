package main 

import (
	"fmt"
)
type person struct{
	firstname string
	lastnmae string
}

func main (){
	h:=person{
		firstname: "h",
		lastnmae: "k",
	}
	fmt.Print(h);
	userData := map[string]int{
		"h":1,
	}
	fmt.Println(userData);
}