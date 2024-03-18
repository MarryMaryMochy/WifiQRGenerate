package route

import (
	"github.com/MarryMaryMochy/WifiQRGenerator/process"
	"github.com/gin-gonic/gin"
)

func Loader(r *gin.Engine) {
	Get(r)
	Post(r)
}

func Get(r *gin.Engine) {
	r.GET("/", process.GeneratePage)
	r.GET("/dl/qr", process.DownloadPage)
}

func Post(r *gin.Engine) {
	r.POST("/generate/qr", process.QRGenerateRoute)
}
