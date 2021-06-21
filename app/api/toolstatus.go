package api

import (
	"context"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Toolstatus = toolApi{}

type toolApi struct{}

// Index is a demonstration route handler for output "Hello World!".
func (*toolApi) Index(r *ghttp.Request) {
	v := g.View()
	v.AddPath("./template/tooling-check-web")
	result, err := v.Parse(context.TODO(), "index.html", g.Map{})
	if err != nil {
		r.Response.WritelnExit(err.Error())
	}
	r.Response.Write(result)
}
