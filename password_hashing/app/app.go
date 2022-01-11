package app

import (
	"fmt"
	"password_hashing/logger"
)

func Start(port string) {
	logger.InfoLogger.Println(fmt.Sprintf("Starting server on localhost:%s ...", port))
}
