package main

import (
	"github.com/j3yzz/encryptor/database"
	"github.com/j3yzz/encryptor/routers"
	"github.com/j3yzz/encryptor/settings"
)

func main() {
	/// Init Settings
	settings.Init()

	/// Init Database
	database.Connect("root:@tcp(localhost:3306)/encryptor?parseTime=true")
	database.Migrate()

	/// Init Router
	router := routers.InitRoutes()
	router.Run(":4000")
}
