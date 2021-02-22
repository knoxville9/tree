package main

import (
	"github.com/gogf/gf/frame/g"
	_ "lqc.com/tree/router"
)

func main() {
	s := g.Server()
	s.Run()
}
