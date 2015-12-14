// API RESTful JSON partiellement conforme à la spec jsonapi.org
// Réponses : {"data":{"id":96,"type":"mailboxes", "attributes":{...}}
// Erreurs : {"errors":[{"status":"409","code":"object.duplicate","title":"Objet en double"...}]}
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
)

var Engine *gin.Engine
var logD *log.Logger
var logI *log.Logger
var logE *log.Logger
var emptyList = make([]string, 0)
var settings ServerSettings

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
	c.String(http.StatusOK, "Wilkommen !")
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
		uc.SetRight(c.Request.RequestURI, c.Request.Method)
		c.Set("contact_id", t.ContactID)
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
	}
	about := Engine.Group("/about")
	about.Use(AuthRequired())
	{
		about.GET("/me", AboutMe)
		about.GET("/api", AboutAPI)
	}
	music := Engine.Group("/music")
	music.Use(AuthRequired())
	{
		music.POST("", OrganizationPost)
		music.GET("/", OrganizationListGet)
		music.GET("/:organization", OrganizationGet)
	}
	display := Engine.Group("/display")
	display.Use(AuthRequired())
	{
		display.POST("", DisplayPost)
		display.GET("/", DisplayListGet)
		display.DELETE("/", DisplayClear)
	}
	apidoc := Engine.Group("/api-doc")
	apidoc.GET("", ApiDoc)
	login := Engine.Group("/auth/token")
	{
		login.POST("", TokenHandler)
	}
	logout := Engine.Group("/auth/logout")
	logout.Use(AuthRequired())
	{
		logout.GET("", LogoutHandler)
	}
}

func main() {
	sapi.Setup()
	var portString = fmt.Sprintf("localhost:%v", settings.Port)
	Engine.Run(portString)
}
