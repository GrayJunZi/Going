package main

import (
	"fmt"
	"myapp/types"
	"myapp/util"
)

func main() {
	user := types.User{
		Username: util.GetUsername(),
		Age:      util.GetNumber(),
	}

	fmt.Printf("%s is %d years old.\n", user.Username, user.Age)
}
