package main

import (
	"github.com/gin-gonic/gin"
	"github.com/planet-work/robotsapin/sapi"
	"net/http"
)

// @Title SensorsGet
// @Description Get the sensors values
// @Accept  json
// @Success 200 {object} sapi.SensorsStatus
// @Failure 406 {object} error "Bad bad bad"
// @Router /topper/ [get]
func SensorsGet(c *gin.Context) {
	sl, err := sapi.Sensors()
	if err == nil {
		c.JSON(http.StatusOK, SapiGinResponse(sl))
		return
	}
	SapiGinError(c, err)
}
