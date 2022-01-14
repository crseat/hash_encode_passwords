package app

import (
	"encoding/json"
	"net/http"
	"password_hashing/errs"
	"password_hashing/logger"
	"path"
	"strings"
)

type Router struct {
	PasswordHandler *PasswordHandler
	StatsHandler    *StatsHandler
}

//define routes
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var head string
	var invalid_data string
	//get endpoint
	head, r.URL.Path = shiftPath(r.URL.Path)

	//check for invalid extra data
	invalid_data, r.URL.Path = shiftPath(r.URL.Path)
	logger.DebugLogger.Println("invalid_data = ", invalid_data)
	if invalid_data != "" {
		invalidEndpointError(w)
	} else {
		//define routes
		switch head {
		case "hash":
			router.PasswordHandler.NewPassword(w, r)
		case "stats":
			//serveContact(w, r)
		default:
			logger.InfoLogger.Println("Attempted endpoint = ", head)
			invalidEndpointError(w)
		}
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

func invalidEndpointError(w http.ResponseWriter) {
	appError := errs.NewValidationError("Please provide a valid endpoint")
	writeResponse(w, http.StatusNotFound, appError.AsMessage())
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
