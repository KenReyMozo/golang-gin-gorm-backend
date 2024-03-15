package initializers

import (
	"fmt"
	"os"
)

var LOG_DIR = "logs"
var LOG_ACTION_DIR = "logs-action"

func SetupDirs() {
	if err := os.MkdirAll(LOG_DIR, os.ModePerm); err != nil {
		panic(fmt.Sprintf("failed to create logs directory: %v", err))
	}

	if err := os.MkdirAll(LOG_ACTION_DIR, os.ModePerm); err != nil {
		panic(fmt.Sprintf("failed to create logs action directory: %v", err))
	}
}