package common

import (
	"log"
	"os"
)

var logger *log.Logger

func SetLog(content string){
	logger = log.New(os.Stdout, "INFO: ", log.LstdFlags)
}