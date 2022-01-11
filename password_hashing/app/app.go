//Package app provides logic for starting the server and wiring together the user side, domain, and server side
package app

import (
	"fmt"
	"password_hashing/logger"
)

//Start the server on the given port.
func Start(port string) {
	logger.InfoLogger.Println(fmt.Sprintf("Starting server on localhost:%s ...", port))
}
