// Package Sapin API (Christmas Tree API).
//
// the purpose of this application is to control a Christmas Tree
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: api.sapin.io
//     BasePath: /v1
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Frederic VANNIRE<f.vanniere@planet-work.com> https://www.planet-work.com/
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

// @APIVersion 1.0.0
// @APITitle sAPI
// @APIDescription Sapin's Application Programming Interface
// @Contact f.vanniere@planet-work.com
// @TermsOfServiceUrl https://api.sapin.io/gcu
// @License BSD
// @LicenseUrl https://api.sapin.io/licence
// @BasePath /doc/v0
// @SubApi Play music [/music]
// @SubApi Display images [/display]
// @SubApi Top star [/star]
import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/planet-work/robotsapin/sapi"
	"github.com/pmylund/go-cache"
)

var Engine *gin.Engine
var logD *log.Logger
var logI *log.Logger
var logE *log.Logger
var emptyList = make([]string, 0)
var settings ServerSettings
var failedLogins = cache.New(1*time.Hour, 1*time.Minute)

type ServerSettings struct {
	Port int `json:"port"`
}

func SapiGinError(c *gin.Context, err error) {
	var id string
	var title string
	var code string
	var e []gin.H
	s := http.StatusInternalServerError
	switch x := err.(type) {
	case *sapi.AuthFailedError:
		// TODO: IPv6 bruteforce prevention
		ip := c.Request.Header.Get("X-Real-Ip")
		var failures int
		if cFailures, found := failedLogins.Get(ip); found {
			failures = cFailures.(int)
			failedLogins.IncrementInt(ip, 1)
		} else {
			failures = 1
			failedLogins.Set(ip, 1, cache.DefaultExpiration)
		}
		sleepTime := time.Duration(failures) * time.Duration(failures) * time.Second
		logD.Println(failures, sleepTime)
		time.Sleep(sleepTime)
		if sleepTime > 60*time.Second {
			s = 429
			code = "failures.too.many"
			title = "Grats: locked out for 1 hour"
		} else {
			s = http.StatusUnauthorized
			code = "access.unauthorized"
			title = "Invalid credentials"
		}
		/*
			case *sapi.QueryFailedError:
				s = http.StatusInternalServerError
				code = "dev.failure"
				title = "Our dba skills sux"
			case *sapi.ResourceForbiddenError:
				s = http.StatusForbidden
				code = "resource.forbidden"
				title = err.Error()
			case *sapi.ResourceNotFoundError, *sapi.RbacForbiddenError:
				s = http.StatusNotFound
				code = "resource.not_found"
				title = err.Error()
			case *sapi.ResourceDuplicateError:
				s = http.StatusConflict
				code = "resource.duplicate"
				title = err.Error()
			case *sapi.ResourceValidationError, *json.SyntaxError:
				s = http.StatusBadRequest
				code = "resource.invalid"
				title = err.Error()
		*/
	default:
		logD.Printf("%T not handled (details: %s)", x, err)
	}
	if s != http.StatusInternalServerError {
		e = []gin.H{gin.H{"status": s, "code": code, "title": title}}
	} else {
		e = []gin.H{gin.H{"id": id, "status": s, "code": code, "title": title}}
	}
	c.JSON(s, gin.H{"errors": e})
}

func objectHash(d interface{}) gin.H {
	v := reflect.ValueOf(d)
	var r gin.H
	switch x := v.Interface().(type) {
	case *APIDetails:
		r = gin.H{"id": "api", "type": "apidetails", "attributes": x}
	case *sapi.SapiStatus:
		r = gin.H{"id": "status", "type": "status", "attributes": x}
	case *sapi.Song:
		r = gin.H{"id": x.Filename, "type": "song", "attributes": x}
	case *sapi.Image:
		r = gin.H{"id": x.Filename, "type": "image", "attributes": x}
	case *sapi.LedSequence:
		r = gin.H{"id": x.Id, "type": "ledsequence", "attributes": x}
	default:
		logE.Println("unknown object: ", x)
	}
	return r
}

func SapiGinResponse(d interface{}) gin.H {
	return gin.H{"data": objectHash(d)}
}

func SapiGinArrayResponse(d interface{}) gin.H {
	s := reflect.ValueOf(d)
	if s.Kind() != reflect.Slice {
		panic("d is a non-slice type")
	}
	l := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		l[i] = objectHash(s.Index(i).Interface())
	}
	return gin.H{"data": l}
}

func OptionsHandler(c *gin.Context) {
	c.String(http.StatusOK, "")
}

func IndexHandler(c *gin.Context) {
	c.String(http.StatusOK, "<a href=\"https://sapin.io\">https://sapin.io/</a>")
}

func StatusHandler(c *gin.Context) {
	s, _ := sapi.GetStatus()
	logD.Println(s)
	c.JSON(http.StatusOK, SapiGinResponse(s))
}

func HealthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "Sapin: OK")
	/*
		if sapi.PingSapin != nil {
			c.String(http.StatusOK, "PgConnection: OK")
		} else {
			c.String(http.StatusOK, "PgConnection: KO")
		}
	*/
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin")
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			//c.Next()
		}
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		th := c.Request.Header.Get("authorization")
		if th == "" || strings.HasPrefix(th, "Bearer ") == false {
			logE.Println("no or badly formatted token")
			SapiGinError(c, &sapi.AuthFailedError{})
			c.Abort()
			return
		}
		tid := strings.TrimPrefix(th, "Bearer ")
		t, err := sapi.TokenQuery(tid)
		if err != nil {
			logE.Println(err)
			SapiGinError(c, &sapi.AuthFailedError{})
			c.Abort()
			return
		}
		failedLogins.Delete(c.Request.Header.Get("X-Real-Ip"))
		uc := sapi.NewUserContext(t)
		//uc.SetRight(c.Request.RequestURI, c.Request.Method)
		c.Set("user_id", t.UserID)
		c.Set("uc", uc)
		//pre
		c.Next()
		//post
	}
}

func initLogger(debugHandler, infoHandler, errorHandler io.Writer) {
	logD = log.New(debugHandler, "[DEBUG] sapi/", log.Lshortfile)
	logI = log.New(infoHandler, "[INFO] sapi/", log.Lshortfile)
	logE = log.New(errorHandler, "[ERROR] sapi/", log.Lshortfile)
}

func resetLogger() {
	sapi.DetectLogMode()
	switch sapi.LogMode() {
	case sapi.DebugLogMode:
		initLogger(os.Stdout, os.Stdout, os.Stderr)
	case sapi.InfoLogMode:
		initLogger(ioutil.Discard, os.Stdout, os.Stderr)
	case sapi.ErrorLogMode:
		initLogger(ioutil.Discard, ioutil.Discard, os.Stderr)
	}
}

func init() {
	resetLogger()
	usr, err := user.Current()
	if err != nil {
		logD.Fatal(err)
	}
	configDir := usr.HomeDir + "/.config/sapi"
	file, err := os.Open(configDir + "/settings.json")
	if err == nil {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&settings)
		if err != nil {
			logD.Fatal("parsing settings.json failed: ", err)
		}
	} else {
		logD.Fatal(err)
	}
}

func init() {
	Engine = gin.Default()
	Engine.Use(CORSMiddleware())
	index := Engine.Group("/")
	{
		index.GET("", IndexHandler)
		index.GET("healthcheck", HealthCheckHandler)
		index.GET("status", StatusHandler)
	}
	about := Engine.Group("/about")
	//about.Use(AuthRequired())
	{
		about.GET("/api", AboutAPI)
	}
	music := Engine.Group("/music")
	//music.Use(AuthRequired())
	{
		//music.POST("/:filename", MusicPost)
		music.GET("/", MusicListGet)
		music.POST("/:filename", MusicPost)
	}
	display := Engine.Group("/display")
	//	display.Use(AuthRequired())
	{
		display.GET("/", DisplayList)
		display.POST("/:filename", DisplayPost)
		display.POST("", DisplayPostData)
	}
	topper := Engine.Group("/topper")
	{
		topper.GET("/", TopperList)
		topper.GET("/:seqId", TopperGet)
		//topper.POST("", DisplayData)
	}
	apidoc := Engine.Group("/api-doc")
	apidoc.GET("", ApiDoc)
	/*
		login := Engine.Group("/auth/token")
		{
			login.POST("", TokenHandler)
		}
		logout := Engine.Group("/auth/logout")
		logout.Use(AuthRequired())
		{
			logout.GET("", LogoutHandler)
		}
	*/
}

func main() {
	sapi.Setup()
	var portString = fmt.Sprintf("localhost:%v", settings.Port)
	Engine.Run(portString)
}
