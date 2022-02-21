package main

import (
	"file-manager/internal/security"
	"file-manager/internal/viewer"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static/", "./static/")

	r.LoadHTMLGlob("./templates/**/*")

	security.RegisterMiddleware(r)
	viewer.RegisterHandler(r)

	r.Run()
}
