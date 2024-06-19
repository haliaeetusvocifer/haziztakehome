package controllers

import (
	"sportteam/initializers"
	"sportteam/models"

	"github.com/gin-gonic/gin"
)

// create a team
func TeamsCreate(c *gin.Context) {
	//get data off req body
	var body struct {
		Name    string
		Logo    string
		Sport   string
		League  string
		Founded string
		Stadium string
		Coach   string
	}

	c.Bind(&body)

	//create teams
	// using array of object to create multiple rows. See gorm documentation for more
	teams := []*models.Team{
		{
			Name:    body.Name,
			Logo:    body.Logo,
			Sport:   body.Sport,
			League:  body.League,
			Founded: body.Founded,
			Stadium: body.Stadium,
			Coach:   body.Coach,
		},
	}

	result := initializers.DB.Create(teams) //pass a slice to insert multiple row

	// check for errors
	if result.Error != nil {
		c.Status(400)
		return
	}

	//Return it
	c.JSON(200, gin.H{
		"team": teams,
	})
}

// Get Teams
func TeamsGetAll(c *gin.Context) {

	//get the teams
	// Get all records
	var teams []models.Team
	initializers.DB.Find(&teams)

	//respond with all the teams
	c.JSON(200, gin.H{
		"team": teams,
	})
}

// Get a team
func TeamsGetATeam(c *gin.Context) {

	//get id from url
	id := c.Param("id")

	//get a single team
	var team models.Team
	initializers.DB.First(&team, id)

	//respond with the them
	c.JSON(200, gin.H{
		"team": team,
	})

}

// Update a Team record
func TeamsUpdate(c *gin.Context) {

	// Get the id from url
	id := c.Param("id")

	//get the data off req body
	var body struct {
		Name    string
		Logo    string
		Sport   string
		League  string
		Founded string
		Stadium string
		Coach   string
	}

	c.Bind(&body)

	//find the team we are updating
	var team models.Team
	initializers.DB.First(&team, id)

	// update the team
	initializers.DB.Model(&team).Updates(models.Team{
		Name:    body.Name,
		Logo:    body.Logo,
		Sport:   body.Sport,
		League:  body.League,
		Founded: body.Founded,
		Stadium: body.Stadium,
		Coach:   body.Coach,
	})

	// respond with it
	c.JSON(200, gin.H{
		"team": team,
	})
}

func TeamsDelete(c *gin.Context) {

	//get the id off the url
	id := c.Param("id")

	//delete the team
	initializers.DB.Delete(&models.Team{}, id)

	//respond with it
	c.Status(200)

}
