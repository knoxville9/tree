package main

import (
	"github.com/gogf/gf/frame/g"
	_ "lqc.com/tree/router"
)

func main() {

	//one, _ := dao.Post.Where("").FindOne(36)
	//
	//fmt.Println(one)
	server := g.Server()
	server.Run()
}
