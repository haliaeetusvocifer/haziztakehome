package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"sportteam/initializers"
	"sportteam/models"

	"github.com/gin-gonic/gin"
)

func FetchTeamAndPlayers(c *gin.Context) {
	//get string from url
	teamName := c.Param("teamName")

	apiURL := fmt.Sprintf("API_URL", teamName)

	// Create an HTTP client
	client := &http.Client{}

	// Make the API request
	resp, err := client.Get(apiURL)
	if err != nil {
		log.Fatal("Failed to establish connection")
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response body")
	}

	// Parse the API response
	var apiResponse struct {
		Teams []struct {
			ID      string `json:"idTeam"`
			Name    string `json:"strTeam"`
			Logo    string `json:"strTeamLogo"`
			Players []struct {
				ID             string `json:"idPlayer"`
				FirstName      string `json:"strPlayer"`
				LastName       string `json:"strPlayerAlternate"`
				ProfilePicture string `json:"strCutout"`
				Position       string `json:"strPosition"`
				Nationality    string `json:"strNationality"`
				DateOfBirth    string `json:"dateBorn"`
				Height         string `json:"strHeight"`
				MarketValue    string `json:"strSigning"`
			} `json:"strPlayer"`
		} `json:"teams"`
	}

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Fatal("Failed, try again")
	}

	// Check if the team exists in the API response
	if len(apiResponse.Teams) == 0 {
		log.Fatal("Failed, Team not found")
	}

	teamData := apiResponse.Teams[0]

	// Check if the team already exists in the database
	team := models.Team{Name: teamData.Name}
	err = initializers.DB.FirstOrCreate(&team, models.Team{Name: teamData.Name}).Error
	if err != nil {
		log.Fatal("Team exist")
	}

	// Update the team's logo
	team.Logo = teamData.Logo
	err = initializers.DB.Save(&team).Error
	if err != nil {
		log.Fatal(" ")
	}

	// Fetch existing players for the team from the database
	var existingPlayers []models.Player
	err = initializers.DB.Where("team_id = ?", team.ID).Find(&existingPlayers).Error
	if err != nil {
		log.Fatal(" ")
	}

	// Create a map of existing player IDs for quick lookup
	existingPlayerIDs := make(map[string]bool)
	for _, player := range existingPlayers {
		playerIDStr := strconv.FormatUint(uint64(player.ID), 10)
		existingPlayerIDs[playerIDStr] = true
		// existingPlayerIDs[player.ID] = true
	}

	var players []*models.Player
	for _, playerData := range teamData.Players {
		player := &models.Player{
			FirstName:      playerData.FirstName,
			LastName:       playerData.LastName,
			ProfilePicture: playerData.ProfilePicture,
			Position:       playerData.Position,
			Nationality:    playerData.Nationality,
			// DateOfBirth:    playerData.DateOfBirth,
			// DateOfBirth:    strconv.FormatFloat(playerData.DateOfBirth, 'f', -1, 64),
			// Height:         strconv.FormatFloat(playerData.Height, 'f', -1, 64),
			// MarketValue:    playerData.MarketValue,
			// TeamID:         team.ID,
		}

		// Check if the player already exists in the database
		if existingPlayerIDs[playerData.ID] {
			// If the player exists, update their details
			err = initializers.DB.Model(&models.Player{}).Where("id = ?", player.ID).Updates(player).Error
			if err != nil {
				log.Fatal(" ")
			}
		} else {
			// If the player doesn't exist, create a new player record
			err = initializers.DB.Create(player).Error
			if err != nil {
				log.Fatal(" ")
			}
		}

		players = append(players, player)
	}

	return
}
