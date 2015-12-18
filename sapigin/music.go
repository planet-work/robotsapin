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
// @Success 200 {object} sapi.Song
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

// @Title MusicPost
// @Description Play song
// @Accept  json
// @Success 200 {object} sapi.Song
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/:filename [post]
func MusicPost(c *gin.Context) {
	filename := c.Params.ByName("filename")
	err := sapi.MusicPlay(filename)
	if err == nil {
		c.JSON(http.StatusOK, "OK")
		return
	}
	SapiGinError(c, err)
}

// @Title MusicPutStop
// @Description Stop Music
// @Accept  json
// @Success 200 {object} sapi.Song
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/stop [put]
func MusicPutStop(c *gin.Context) {
	err := sapi.MusicStop()
	if err == nil {
		c.JSON(http.StatusOK, "OK")
		return
	}
	SapiGinError(c, err)
}

// @Title MusicPutPause
// @Description Pause Music
// @Accept  json
// @Success 200 {object} sapi.Song
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/pause [put]
func MusicPutPause(c *gin.Context) {
	err := sapi.MusicPause()
	if err == nil {
		c.JSON(http.StatusOK, "OK")
		return
	}
	SapiGinError(c, err)
}

// @Title MusicPutVolumeUp
// @Description VolumeUp
// @Accept  json
// @Success 200 {object} sapi.Song
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/volume+ [put]
func MusicPutVolumeUp(c *gin.Context) {
	newvolume, err := sapi.MusicVolumeUp()
	if err == nil {
		c.JSON(http.StatusOK, newvolume)
		return
	}
	SapiGinError(c, err)
}

// @Title MusicPutVolumeDown
// @Description Volume Down
// @Accept  json
// @Success 200 {object} sapi.Song
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/volume- [put]
func MusicPutVolumeDown(c *gin.Context) {
	newvolume, err := sapi.MusicVolumeDown()
	if err == nil {
		c.JSON(http.StatusOK, newvolume)
		return
	}
	SapiGinError(c, err)
}
