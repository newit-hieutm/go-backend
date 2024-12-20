package initialize

import (
	"fmt"

	"github.com/newit-hieutm/go-backend/global"
)

func Run() {
	InitConfig()
	fmt.Println("config", global.ConfigGlobal)

	InitLogger()
	InitMysql()
	InitRouter()
}
