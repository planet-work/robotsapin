package main

import (
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	//	"github.com/gin-gonic/gin/binding"
	"github.com/planet-work/robotsapin/sapi"
)

// @Title DisplayListGet
// @Description Retrieve the list of available music
// @Accept  json
// @Success 200 {object} sapi.Organization
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/ [get]
func DisplayList(c *gin.Context) {
	sl, err := sapi.DisplayList()
	if err == nil {
		c.JSON(http.StatusOK, SapiGinArrayResponse(sl))
		return
	}
	SapiGinError(c, err)
}

// @Title DisplayGet
// @Description Play song
// @Accept  json
// @Success 200 {object} sapi.Song
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/:filename [get]
func DisplayImage(c *gin.Context) {
	filename := c.Params.ByName("filename")
	err := sapi.DisplayImage(filename)
	if err == nil {
		c.JSON(http.StatusOK, SapiGinArrayResponse("OK"))
		return
	}
	SapiGinError(c, err)
}

func DisplayData(c *gin.Context) {
	data := c.Params.ByName("data")
	err := sapi.DisplayData(data)
	if err == nil {
		c.JSON(http.StatusOK, SapiGinArrayResponse("OK"))
		return
	}
	SapiGinError(c, err)
}
