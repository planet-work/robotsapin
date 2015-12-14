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
	ad := APIDetails{Version: papi.Version(), Now: time.Now()}
	return &ad
}

func AboutMe(c *gin.Context) {
	uc := c.MustGet("uc").(*papi.UserContext)
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"type": "available", "attributes": gin.H{"whoami": uc.ContactID, "roles": uc.Roles, "tenants": uc.AvailableTenants(), "accounts": uc.AvailableAccounts(), "zones": uc.AvailableZones(), "contacts": uc.AvailableContacts()}}})
}

func AboutAPI(c *gin.Context) {
	c.JSON(http.StatusOK, SapiGinResponse(NewAPIDetails()))
}
