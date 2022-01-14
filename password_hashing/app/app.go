//Package app provides logic for starting the server and wiring together the user side, domain, and server side
package app

import (
	"fmt"
	"log"
	"net/http"
	"password_hashing/domain"
	"password_hashing/logger"
	"password_hashing/service"
)

//Start the server on the given port.
func Start(port string) {
	logger.InfoLogger.Println(fmt.Sprintf("Starting server on localhost:%s ...", port))

	//wiring
	newHashRepo := make(map[int64]string)
	hashRepository := domain.NewHashRepository(newHashRepo)
	passwordHandler := PasswordHandler{service: service.NewPasswordService(hashRepository)}

	//Create Router object
	router := &Router{
		PasswordHandler: &passwordHandler,
	}
	address := "localhost"
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
