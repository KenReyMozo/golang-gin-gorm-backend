package logger

import (
	"fmt"
	"go-backend/initializers"
	"log"
	"os"
	"time"
)

var ACTION_LOGIN = "login"
var ACTION_LOGOUT = "logout"

func LogUserAction(userID, userEmail string,  action ...string ) {

	logFileName := fmt.Sprintf("%s/%s_%s.log", initializers.LOG_ACTION_DIR, time.Now().Format("2006-01-02"), userID)

	file, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	logMessage := fmt.Sprintf("[%s] User %s %s\n", time.Now().Format("2006-01-02 15:04:05"), userEmail, action)

	_, err = file.WriteString(logMessage)
	if err != nil {
		log.Fatalf("Failed to write to log file: %v", err)
	}
}
