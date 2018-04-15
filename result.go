package main

import (
	"github.com/kataras/iris"
)

type Result struct {
	code int 
	data interface{}
	msg string 
}

func (r *Result) success(ctx iris.Context) {
  ctx.JSON(iris.Map{"code": 200, "data": r.data, "msg":"success"})
}

func (r *Result) failed(ctx iris.Context) {
  ctx.JSON(iris.Map{"code":r.code , "data": nil, "msg":r.msg})
}



