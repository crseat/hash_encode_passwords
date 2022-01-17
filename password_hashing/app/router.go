package app

import (
	"encoding/json"
	"net/http"
	"password_hashing/errs"
	"password_hashing/logger"
	"path"
	"strings"
	"sync"
	"time"
)

type Router struct {
	PasswordHandler *PasswordHandler
	WaitGroup       *sync.WaitGroup
	quitChan        *chan bool
}

//define routes
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var head string
	var startTime time.Time

	//Check for invalid method.
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		invalidMethodError(w)
	}

	//Keep track of number of post request for stats
	if r.Method == http.MethodPost {
		router.PasswordHandler.passwordService.IncTotal()
		startTime = time.Now()
	}

	//get endpoint
	head, r.URL.Path = shiftPath(r.URL.Path)
	//define routes
	switch head {
	case "hash":
		//check for id
		var identifier string
		identifier, r.URL.Path = shiftPath(r.URL.Path)
		if identifier != "" {
			router.PasswordHandler.FindBy(w, identifier)
		} else {
			router.PasswordHandler.NewPassword(w, r, startTime, router.WaitGroup)
		}
	case "stats":
		//check for invalid data ex: localhost:8000/stats/anyData
		if existsInvalidData(w, r.URL.Path) {
			return
		} else {
			// Call stats function
			router.PasswordHandler.GetStats(w)
		}
	case "shutdown":
		//Shutdown gracefully
		shutdown(*router.quitChan)
	default:
		logger.DebugLogger.Println("Attempted invlaid endpoint = ", head)
		invalidEndpointError(w)
	}
}

// shiftPath splits the given path into the first segment (head) and  the rest (tail).
// For example, "/foo/bar/baz" gives "foo", "/bar/baz".
//source: https://blog.merovius.de/2017/06/18/how-not-to-use-an-http-router.html
func shiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func shutdown(quit chan bool) {
	quit <- true
}

func invalidMethodError(w http.ResponseWriter) {
	appError := errs.NewValidationError("Method is not supported")
	writeResponse(w, http.StatusNotFound, appError.AsMessage())
}

func invalidEndpointError(w http.ResponseWriter) {
	appError := errs.NewValidationError("Please provide a valid endpoint")
	writeResponse(w, http.StatusNotFound, appError.AsMessage())
}

func existsInvalidData(w http.ResponseWriter, endpoint string) (err bool) {
	var testData string
	testData, endpoint = shiftPath(endpoint)
	if testData != "" {
		invalidEndpointError(w)
		return true
	}
	return false
}

// writeResponse formats all http responses to client into json
func writeResponse(writer http.ResponseWriter, code int, data interface{}) {
	// We need to define the header here or the json/xml response will come across as plain text
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		panic(err)
	}
}
