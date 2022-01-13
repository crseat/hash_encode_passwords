//Package app provides logic for starting the server and wiring together the user side, domain, and server side
package app

import (
	"fmt"
	"password_hashing/domain"
	"password_hashing/logger"
	"password_hashing/service"
)

//Start the server on the given port.
func Start(port string) {
	logger.InfoLogger.Println(fmt.Sprintf("Starting server on localhost:%s ...", port))

	//wiring
	newHashRepo := make(map[int]string)
	hashRepository := domain.NewHashRepository(newHashRepo)
	passwordHandler := PasswordHandler{service: service.NewPasswordService(hashRepository)}

	//starting the server
}
