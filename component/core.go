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
	//附加数据
	body, err := c.client.HttpPostJson(path, util.JsonEncode(request), headers...)
	if err != nil {
		return
	}
	c.ctx.Logger.Debug("body", body)
	////转换
	//j, err := gjson.DecodeToJson(body)
	//if err != nil {
	//	return
	//}
	//
	////转换
	//err = j.ToStruct(response)
	//if err != nil {
	//	return
	//}
	err = util.JsonDecode(body, response)
	return
}
