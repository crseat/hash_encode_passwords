package app

import (
	"encoding/json"
	"net/http"
	"password_hashing/domain"
	"password_hashing/service"
	"path"
	"strings"
)

type Router struct {
}

//define routes
func serve(w http.ResponseWriter, r *http.Request) {

	//wiring
	hashRepository := domain.NewHashRepository()
	passwordHandler := PasswordHandlers{service: service.NewPasswordService(hashRepository)}

	//define routes
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	switch head {
	case "hash":
		passwordHandler.service.NewHash()
	case "stats":
		serveContact(w, r)
	default:
		writeResponse(w, http.StatusNotFound, head)
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

// writeResponse formats all http responses to client into json
func writeResponse(writer http.ResponseWriter, code int, data string) {
	// We need to define the header here or the json/xml response will come across as plain text
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		panic(err)
	}
}
