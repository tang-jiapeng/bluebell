package main

import (
	"bluebell/router"
	"fmt"
)

func main() {
	r := router.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", 8080))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
