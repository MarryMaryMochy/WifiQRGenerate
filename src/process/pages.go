package process

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GeneratePage(c *gin.Context) {
	c.HTML(http.StatusOK, "make_qr_form.html", gin.H{})
}

func DownloadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "finish.html", gin.H{})
}
