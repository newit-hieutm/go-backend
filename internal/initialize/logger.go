package initialize

import (
	"fmt"

	"github.com/newit-hieutm/go-backend/global"
	"github.com/newit-hieutm/go-backend/pkg/loggers"
)

func InitLogger() {
	global.Logs = loggers.InitLogger()

	fmt.Println("Logs:", global.Logs)
	fmt.Println("Initialized logger successfully!")
}
