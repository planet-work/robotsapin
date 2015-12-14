package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/planet-work/robotsapin/sapi"
)

type APIDetails struct {
	Version string    `json:"version"`
	Now     time.Time `json:"now"`
}

func NewAPIDetails() *APIDetails {
	ad := APIDetails{Version: sapi.Version(), Now: time.Now()}
	return &ad
}

func AboutAPI(c *gin.Context) {
	c.JSON(http.StatusOK, SapiGinResponse(NewAPIDetails()))
}
