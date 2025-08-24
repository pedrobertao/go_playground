package main

import (
	"fmt"
	"unique"
)

func unique_example() {

	v1 := "Pedro"
	v2 := "Pedro"
	u1 := unique.Make(v1)
	u2 := unique.Make(v2)
	fmt.Println("u1 == u2 if v1 == v2", u1 == u2)
}
