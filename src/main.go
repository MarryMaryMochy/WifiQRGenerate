package main

import (
	"github.com/MarryMaryMochy/WifiQRGenerator/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	route.Loader(r)
	r.Run(":8080")
}
