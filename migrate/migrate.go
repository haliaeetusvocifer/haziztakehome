package main

import (
	"sportteam/initializers"
	"sportteam/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Team{})
	initializers.DB.AutoMigrate(&models.Player{})
}
