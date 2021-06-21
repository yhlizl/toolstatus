package main

import (
	_ "sreTool/boot"
	_ "sreTool/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()

}
