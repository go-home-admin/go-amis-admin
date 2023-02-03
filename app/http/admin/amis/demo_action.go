package amis

import (
	gin "github.com/gin-gonic/gin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/protobuf"
)

// Demo   测试
func (receiver *Controller) Demo(req *admin.DemoRequest, ctx http.Context) (*admin.DemoResponse, error) {
	// TODO 这里写业务
	return &admin.DemoResponse{
		Title: "test json",
		Body: &protobuf.Any{
			B: []byte("{\n    \"mode\": \"horizontal\",\n    \"type\": \"form\",\n    \"body\": [\n      {\n        \"type\": \"input-text\",\n        \"label\": \"文本框\",\n        \"name\": \"text\",\n        \"size\": \"md\"\n      },\n      {\n        \"type\": \"input-password\",\n        \"label\": \"<a href='./password'>密码</a>\",\n        \"name\": \"password\",\n        \"size\": \"md\"\n      }\n    ]\n  }"),
		},
	}, nil
}

// GinHandleDemo gin原始路由处理
// http.Post(/amis/tabs)
func (receiver *Controller) GinHandleDemo(ctx *gin.Context) {
	req := &admin.DemoRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Demo(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
