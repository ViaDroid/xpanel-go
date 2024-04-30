package utils

import "github.com/beego/beego/v2/server/web"

type Response interface {
	Success()
}

type AbStruct[T Response] struct{}

type CdStruct[T Response] struct{}

type ResponseHelper struct{}

func (ResponseHelper) Success(c *web.Controller, msg string) {
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.JSONResp(map[string]any{
		"ret": 1,
		"msg": msg,
	})
}

func (ResponseHelper) SuccessWithData(c *web.Controller, msg string, data any) {
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.JSONResp(map[string]any{
		"ret":  1,
		"msg":  msg,
		"data": data,
	})
}

func (ResponseHelper) SuccessWithDataEtag(c *web.Controller, data any) {
	c.Ctx.Output.Header("Content-Type", "application/json")
	etag := ""
	if etag == c.Ctx.Input.Header("If-None-Match") {
		c.Abort("304")
	}
	c.Ctx.Request.Response.Header.Add("ETag", etag)
	c.JSONResp(data)
}

func (h ResponseHelper) Error(c *web.Controller, msg string) {
	c.Ctx.Output.Header("Content-Type", "application/json")

	c.JSONResp(map[string]any{
		"ret": 0,
		"msg": msg,
	})

	// c.Ctx.WriteString(msg)
	// c.Abort("200")
}

func (h ResponseHelper) ErrorWithData(c *web.Controller, msg string, data any) {
	c.Ctx.Output.Header("Content-Type", "application/json")
	c.JSONResp(map[string]any{
		"ret":  0,
		"msg":  msg,
		"data": data,
	})
}
