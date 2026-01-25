package main

import (
	"fmt"
	"go_intro/modules/modules/types"
	"go_intro/modules/modules/util"
)

func main() {
	user := types.User{
		Username: util.GetUsername(),
		Age:      util.GetAge(),
	}
	fmt.Printf("The user is: %+v", user)
}
