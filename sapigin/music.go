package main

import (
	"net/http"
	//"strconv"

	"github.com/gin-gonic/gin"
	//	"github.com/gin-gonic/gin/binding"
	"github.com/planet-work/robotsapin/sapi"
)

// @Title MusicListGet
// @Description Retrieve the list of available music
// @Accept  json
// @Success 200 {object} sapi.Organization
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/ [get]
func MusicListGet(c *gin.Context) {
	sl, err := sapi.MusicList()
	if err == nil {
		c.JSON(http.StatusOK, SapiGinArrayResponse(sl))
		return
	}
	SapiGinError(c, err)
}

// @Title MusicGet
// @Description Play song
// @Accept  json
// @Success 200 {object} sapi.Song
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/:filename [post]
func MusicPost(c *gin.Context) {
	filename := c.Params.ByName("filename")
	err := sapi.MusicPlay(filename)
	if err == nil {
		c.JSON(http.StatusOK, SapiGinArrayResponse("OK"))
		return
	}
	SapiGinError(c, err)
}
