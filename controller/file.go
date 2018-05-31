package controller

import (
	"io"
	"os"

	"github.com/kataras/iris"
)

// FileUpload 文件上传
func FileUpload(ctx iris.Context) {
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	defer file.Close()
	fname := info.Filename
	out, err := os.OpenFile("./uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.JSON(iris.Map{"code": "-1", "msg": err})
		return
	}
	defer out.Close()
	io.Copy(out, file)
	ctx.JSON(iris.Map{"code": "200", "msg": err, "url": "/uploads/" + fname})
}
