package main

import (
	"github.com/amanverasia/bitmagnet/internal/dev/app"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app.New().Run()
}
