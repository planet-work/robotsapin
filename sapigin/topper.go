package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	//	"github.com/gin-gonic/gin/binding"
	"github.com/planet-work/robotsapin/sapi"
)

// @Title TopperList
// @Description Retrieve the list of available topper sequences
// @Accept  json
// @Success 200 {object} sapi.TopperSequence
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/ [get]
func TopperList(c *gin.Context) {
	sl, err := sapi.TopperList()
	if err == nil {
		c.JSON(http.StatusOK, SapiGinArrayResponse(sl))
		return
	}
	SapiGinError(c, err)
}

// @Title Topper
// @Description Play song
// @Accept  json
// @Success 200 {object} sapi.Song
// @Failure 406 {object} error "Bad bad bad"
// @Router /music/:filename [get]
func TopperGet(c *gin.Context) {
	var seqId, speed int
	var err error
	seqId, err = strconv.Atoi(c.Params.ByName("seqId"))
	if err != nil {
		seqId = 1
	}
	speed, err = strconv.Atoi(c.Params.ByName("speed"))
	if err != nil {
		speed = 100
	}

	err = sapi.TopperSetSequence(seqId, speed)
	if err == nil {
		c.JSON(http.StatusOK, SapiGinArrayResponse("OK"))
		return
	}
	SapiGinError(c, err)
}
