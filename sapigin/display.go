package main

import (
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	//	"github.com/gin-gonic/gin/binding"
	"github.com/planet-work/robotsapin/sapi"
)

// @Title DisplayList
// @Description Retrieve the list of available pictures
// @Accept  json
// @Success 200 {object} sapi.Image
// @Failure 406 {object} error "Bad bad bad"
// @Router /display/ [get]
func DisplayList(c *gin.Context) {
	sl, err := sapi.DisplayList()
	if err == nil {
		c.JSON(http.StatusOK, SapiGinArrayResponse(sl))
		return
	}
	SapiGinError(c, err)
}

// @Title DisplayPost
// @Description Display giver image from parameters
// @Accept  json
// @Success 200 {object} sapi.Song
// @Failure 406 {object} error "Bad bad bad"
// @Router /display/:filename [post]
func DisplayPost(c *gin.Context) {
	filename := c.Params.ByName("filename")
	img, err := sapi.DisplayImage(filename)
	if err == nil {
		c.JSON(http.StatusOK, SapiGinResponse(img))
		return
	}
	SapiGinError(c, err)
}

// @Title DisplayPostData
// @Description Display image data from paratemerts
// @Accept  json
// @Success 200 {object} string
// @Failure 406 {object} error "Bad bad bad"
// @Router /display/ [post]
func DisplayPostData(c *gin.Context) {
	data := c.Params.ByName("data")
	img, err := sapi.DisplayData(data)
	if err == nil {
		c.JSON(http.StatusOK, SapiGinResponse(img))
		return
	}
	SapiGinError(c, err)
}
