package main

import (
	"fmt"

	"github.com/spf13/cast"
)

type dynamic interface{}

func main() {
	var d1 dynamic

	d1 = "name"
	fmt.Println(d1)

	d1 = 10
	fmt.Println(d1)

	switch d1.(type) {
	case string:
		fmt.Println("variable is string")
	case int:
		d1 = "string"
		fmt.Println("variable is integer")
		
		intValue := cast.ToInt(d1)
		fmt.Printf("type %T\n", intValue)
	}
}
