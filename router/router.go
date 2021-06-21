package router

import (
	"sreTool/app/api"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/toolstatus", func(group *ghttp.RouterGroup) {
		group.ALL("/", api.Toolstatus)

	})

}
