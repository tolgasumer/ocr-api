package controllers

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract/v2"
)

var (
	imgexp = regexp.MustCompile("^image")
)

// FileUpload ...
func FileUpload(c *gin.Context) {

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer file.Close()

	tempBuffer := bytes.NewBuffer(nil)
	if _, err = io.Copy(tempBuffer, file); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	defer tempBuffer.Reset() // TODO: bu tam olarak bufferi silmiyor

	client := gosseract.NewClient()
	defer client.Close()

	client.SetImageFromBytes(tempBuffer.Bytes())

	client.Languages = []string{"eng"}
	if langs := c.Request.FormValue("languages"); langs != "" {
		client.Languages = strings.Split(langs, ",")
	}
	if whitelist := c.Request.FormValue("whitelist"); whitelist != "" {
		client.SetWhitelist(whitelist)
	}

	var out string
	switch c.Request.FormValue("format") {
	case "hocr":
		out, err = client.HOCRText()
		//r.render.EscapeHTML = false
	default:
		out, err = client.Text()
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": strings.Trim(out, c.Request.FormValue("trim")),
	})
}
