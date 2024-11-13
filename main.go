package main

import (
	"bluebell/controller"
	"bluebell/router"
	"fmt"
)

func main() {

	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err:%v\n", err)
		return
	}

	r := router.SetupRouter()
	err := r.Run(fmt.Sprintf(":%d", 8080))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
