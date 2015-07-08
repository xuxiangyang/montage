package main

import (
	"fmt"
	"github.com/xuxiangyang/montage/models/user"
)

func main() {
	fmt.Println(user.Fetch(1))
}
