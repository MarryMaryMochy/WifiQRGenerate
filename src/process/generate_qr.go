package process

import (
	"image/png"
	"net/http"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
)

func QRGenerateRoute(c *gin.Context) {
	ssid := c.PostForm("ssid")
	password := c.PostForm("password")
	if ssid == "" || password == "" {
		c.HTML(http.StatusBadRequest, "bad_request.html", gin.H{})
	}
	WifiQRGenerate(ssid, password)
	c.Redirect(http.StatusSeeOther, "/dl/qr")
}

func WifiQRGenerate(ssid string, password string) {
	qrCode, _ := qr.Encode("WIFI:S:{"+ssid+"};T:WPA;P:{"+password+"};;", qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	file, _ := os.Create("wifiqr.png")
	defer file.Close()

	png.Encode(file, qrCode)
}
