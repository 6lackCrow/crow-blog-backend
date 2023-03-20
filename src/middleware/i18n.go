package middleware

import "github.com/kataras/iris/v12"

func I18nMiddleware(ctx iris.Context) {
	language := ctx.GetHeader("language")
	ctx.Values().Set("language", language)
	ctx.Next()
}
