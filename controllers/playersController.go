package controllers

import (
	"sportteam/initializers"
	"sportteam/models"
	"time"

	"github.com/gin-gonic/gin"
)

// create a Player
func PlayersCreate(c *gin.Context) {
	//get data off req body
	var body struct {
		FirstName      string
		LastName       string
		ProfilePicture string
		Team           string
		TeamID         string
		Position       string
		Nationality    string
		DateOfBirth    time.Time
		Height         float32
		MarketValue    float32
	}

	c.Bind(&body)

	//create players
	// using array of object to create multiple rows. See gorm documentation for more
	players := []*models.Player{
		{
			FirstName:      body.FirstName,
			LastName:       body.LastName,
			ProfilePicture: body.ProfilePicture,
			Team:           body.Team,
			TeamID:         body.TeamID,
			Position:       body.Position,
			Nationality:    body.Nationality,
			DateOfBirth:    body.DateOfBirth,
			Height:         body.Height,
			MarketValue:    body.MarketValue,
		},
	}

	result := initializers.DB.Create(players) //pass a slice to insert multiple row

	// check for errors
	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"player": players,
	})
}

// Get Player
func PlayersGetAll(c *gin.Context) {

	//get the players
	// Get all records
	var players []models.Player
	initializers.DB.Find(&players)

	//respond with all the teams
	c.JSON(200, gin.H{
		"player": players,
	})
}

// Get a player
func PlayersGetplayer(c *gin.Context) {

	//get id from url
	id := c.Param("id")

	//get a single player
	var player models.Player
	initializers.DB.First(&player, id)

	//respond with the player
	c.JSON(200, gin.H{
		"player": player,
	})

}

// Get all players for a team
func PlayersGetTeamplayer(c *gin.Context) {

	//get id from url
	id := c.Param("id")

	// Get all matched records
	var players []models.Player
	initializers.DB.Where("Team <> ?", id).Find(&players)

	//respond with the player
	c.JSON(200, gin.H{
		"player": players,
	})

}

// Update a player record
func PlayersUpdate(c *gin.Context) {

	// Get the id from url
	id := c.Param("id")

	//get the data off req body
	var body struct {
		FirstName      string
		LastName       string
		ProfilePicture string
		Team           string
		TeamID         string
		Position       string
		Nationality    string
		DateOfBirth    time.Time
		Height         float32
		MarketValue    float32
	}

	c.Bind(&body)

	//find the player we are updating
	var player models.Player
	initializers.DB.First(&player, id)

	// update the player
	initializers.DB.Model(&player).Updates(models.Player{
		FirstName:      body.FirstName,
		LastName:       body.LastName,
		ProfilePicture: body.ProfilePicture,
		Team:           body.Team,
		TeamID:         body.TeamID,
		Position:       body.Position,
		Nationality:    body.Nationality,
		DateOfBirth:    body.DateOfBirth,
		Height:         body.Height,
		MarketValue:    body.MarketValue,
	})

	// respond with it
	c.JSON(200, gin.H{
		"player": player,
	})
}

func PlayersDelete(c *gin.Context) {

	//get the id off the url
	id := c.Param("id")

	//delete the player
	initializers.DB.Delete(&models.Player{}, id)

	//respond with it
	c.Status(200)

}
