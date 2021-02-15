package main

import (
	"github.com/gogf/gf/frame/g"
	_ "lqc.com/tree/router"
)

// @title       tree project
// @version     1.0
// @schemes     http

func main() {
	server := g.Server()
	server.Run()
}
