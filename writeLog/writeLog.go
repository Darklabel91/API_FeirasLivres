package writeLog

import (
	"log"
	"os"
)

//WriteLog function to create log.txt
func WriteLog() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}
