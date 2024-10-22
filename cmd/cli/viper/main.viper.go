package main

import (
	"fmt"
	"os"

	"github.com/newit-hieutm/go-backend/configs"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()

	v.AddConfigPath("./configs")
	env := os.Getenv("GO_APP_ENV")
	if env == "" {
		env = "local" // default to development
	}
	v.SetConfigName(env)
	v.SetConfigType("yml")

	// Read the configuration file
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file: %v\n", err)
		return
	}

	c := configs.Config{}
	err := v.Unmarshal(&c)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	fmt.Println(c.Security.Jwt.TokenSecret)
}
