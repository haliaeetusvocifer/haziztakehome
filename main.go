package main

import (
	"sportteam/controllers"
	"sportteam/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()

	//Teams
	r.POST("/teams", controllers.TeamsCreate)
	r.GET("/teams", controllers.TeamsGetAll)
	r.GET("/teams/:id", controllers.TeamsGetATeam)
	r.PUT("/teams/:id", controllers.TeamsUpdate)
	r.DELETE("/teams/:id", controllers.TeamsDelete)

	//Players
	r.POST("/players", controllers.PlayersCreate)
	r.GET("/players", controllers.PlayersGetAll)
	r.GET("/players/:id", controllers.PlayersGetplayer)
	r.GET("/players/:id/players", controllers.PlayersGetTeamplayer)
	r.PUT("/players/:id", controllers.PlayersUpdate)
	r.DELETE("/players/:id", controllers.PlayersDelete)

	//Service Sync
	r.POST("/syncservices/:teamName", controllers.FetchTeamAndPlayers)

	r.Run()
}
