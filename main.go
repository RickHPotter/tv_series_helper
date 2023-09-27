package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// get the key on the site API
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file.")
	}

	api_key, exists := os.LookupEnv("API_KEY")
	if !exists {
		panic("API_KEY environment variable not found or not set.")
	}

	regex_pattern := `S(\d+)E(\d+)`
	// regex_pattern := `(\d+)`

	for i := 1; i <= 4; i++ {
		folder := fmt.Sprintf(
			"/mnt/d/A THOUSAND LIVES/&VISUAL/TV SERIES/British/[F] A Bit Of Fry & Laurie/Season %02d",
			i,
		)

		// gotta check the site to get the correct tv_id
		tv_id := 192
		season_number := i
		starting_episode_number := 1

		episode_names := RetrieveEpisodeNames(
			api_key,
			tv_id,
			season_number,
			starting_episode_number,
		)

		RenameFolderEpisodes(folder, episode_names, regex_pattern)
	}
}
