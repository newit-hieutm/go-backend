package main

import (
	"github.com/newit-hieutm/go-backend/internal/routers"
)

func main() {
  r := routers.NewRouter()

  r.Run(":8888")
}