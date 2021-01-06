package component

import (
	"github.com/kaylyu/ministore/context"
	"github.com/kaylyu/ministore/models"
	"github.com/kaylyu/ministore/util"
)

type Core struct {
	ctx    *context.Context
	client *util.Client
}

func NewCore(ctx *context.Context) *Core {
	return &Core{ctx, util.NewHttpClient(ctx)}
}

//POST
func (c *Core) HttpPostJson(path string, request models.IAPIRequest, response interface{}, headers ...map[string]string) (err error) {
	bComponentAccessToken := len(request.GetComponentAccessToken()) != 0
	if bComponentAccessToken {
		c.ctx.Config.ComponentAccessToken = request.GetComponentAccessToken()
	}
	//附加数据
	body, err := c.client.HttpPostJson(path, util.JsonEncode(request), bComponentAccessToken, headers...)
	if err != nil {
		return
	}
	c.ctx.Logger.Debug("body", body)
	err = util.JsonDecode(body, response)
	return
}

//上传文件
func (c *Core) Upload(path, filename string, response interface{}) (err error) {
	body, err := c.client.Upload(path, filename)
	if err != nil {
		return
	}
	c.ctx.Logger.Debug("body", body)
	err = util.JsonDecode(body, response)
	return
}
