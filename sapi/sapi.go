package sapi

import (
	"encoding/json"
	"fmt"
	"io"
	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"strings"
	//"time"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB
var ConfigDir string
var Settings APISettings
var logD *log.Logger
var logI *log.Logger
var logE *log.Logger
var Status SapiStatus

func Version() string {
	out, err := exec.Command("git", "describe").Output()
	if err != nil {
		log.Fatal(err)
	}
	v := strings.TrimSuffix(string(out), "\n")

	return v
}

type Me struct {
	Username string `json:"username"`
}

type SapiStatus struct {
	Version string        `json:"version"`
	Display DisplayStatus `json:"display"`
	Topper  ToperStatus   `json:"topper"`
	Music   MusicStatus   `json:"music"`
	Sensors SensorsStatus `json:"sensors"`
}

type APISettings struct {
	DbPath        string `json:"db_path"`
	MusicDir      string `json:"music_dir"`
	PictureDir    string `json:"pictures_dir"`
	AdminPassword string `json:"admin_password"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title>list</title></head><body>")
	fmt.Fprintf(w, "<p>hello</p>")
	fmt.Fprintf(w, "</body></html>")
	fmt.Println(r)
}

func GetStatus() (*SapiStatus, error) {
	return &Status, nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//var sessid string
	//var id string
	var pwd string
	//var tk Token
	err := r.ParseForm()
	if err != nil {
		logD.Fatal(err)
	}
	if r.FormValue("contact_id") != "" {
		fmt.Printf("%v\n", r.FormValue("contact_id"))
		//id = r.FormValue("contact_id")
	}
	if r.FormValue("password") != "" {
		fmt.Printf("%v\n", r.FormValue("password"))
		pwd = r.FormValue("password")
	}

	if pwd == Settings.AdminPassword {
		out := "xxxxx"
		fmt.Printf("%v\n", out)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(out))
	} else {
		fmt.Println("wrong")
	}
}

func isSessionActive(sessid string) bool {
	/*
		var id string
			err := db.QueryRow("SELECT id FROM session WHERE id = $1", sessid).Scan(&id)
			if err != nil {
				fmt.Printf("wrong: %v\n", err)
				return false
			}*/
	return true
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	me := Me{
		Username: "you",
	}
	out, err := json.Marshal(me)
	if err != nil {
		logD.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

type handler func(w http.ResponseWriter, r *http.Request)

func JSONError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Pw-Api-Token, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Origin")
	var empty []byte
	w.Write(empty)
}

func BasicAuth(pass handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		/*auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)
		  if len(auth) != 2 || auth[0] != "Basic" {
		      http.Error(w, "bad syntax", http.StatusBadRequest)
		      return
		  }
		  payload, _ := base64.StdEncoding.DecodeString(auth[1])
		  pair := strings.SplitN(string(payload), ":", 2)*/
		fmt.Printf("%v\n", r.Header)
		fmt.Printf("%v\n", r.Header.Get("X-Pw-Token"))
		if isSessionActive(r.Header.Get("X-Pw-Token")) {
			/*if len(pair) != 2 || !Validate(pair[0], pair[1]) {
			    http.Error(w, "authorization failed", http.StatusUnauthorized)
			    return
			}*/
			pass(w, r)
		} else {
			JSONError(w, `{"error":"authorization failed"}`, http.StatusUnauthorized)
			return
		}
	}
}

func Validate(username, password string) bool {
	if username == "username" && password == "password" {
		return true
	}
	return false
}

type LogRecord struct {
	http.ResponseWriter
	status int
}

func (r *LogRecord) Write(p []byte) (int, error) {
	return r.ResponseWriter.Write(p)
}

func (r *LogRecord) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		record := &LogRecord{
			ResponseWriter: w,
			status:         200,
		}
		handler.ServeHTTP(record, r)
		logD.Println(r.RemoteAddr, r.Method, r.URL, record.status)
	})
}

func PingDb() error {
	return nil
}

func Setup() {
	var err error
	DetectLogMode()
	usr, err := user.Current()
	if err != nil {
		logE.Fatal(err)
	}
	ConfigDir = usr.HomeDir + "/.config/sapi"
	file, err := os.Open(ConfigDir + "/settings.json")
	if err == nil {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&Settings)
		if err != nil {
			logE.Fatal("parsing settings.json failed: ", err)
		}
		dburi := fmt.Sprintf("sqlite:%s application_name=sapi", Settings.DbPath)
		logD.Println(dburi)
		db, err = sqlx.Open("sqlite3", dburi)
	} else {
		logE.Fatal(err)
	}
}

func initLogger(debugHandler, infoHandler, errorHandler io.Writer) {
	logD = log.New(debugHandler, "[DEBUG] sapi/", log.Lshortfile)
	logI = log.New(infoHandler, "[INFO] sapi/", log.Lshortfile)
	logE = log.New(errorHandler, "[ERROR] sapi/", log.Lshortfile)
}
