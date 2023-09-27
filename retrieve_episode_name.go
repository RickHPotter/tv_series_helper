package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a struct to represent the JSON response
type EpisodeInfo struct {
	Name          string `json:"name"`
	SeasonNumber  int    `json:"season_number"`
	EpisodeNumber int    `json:"episode_number"`
}

func RetrieveEpisodeNames(
	api_key string,
	tv_id int,
	season_number int,
	starting_episode_number int,
) map[string]string {
	episode_map := make(map[string]string)

	episode_number := starting_episode_number

	for {
		apiUrl := fmt.Sprintf(
			"https://api.themoviedb.org/3/tv/%d/season/%d/episode/%d?api_key=%s&language=en-US",
			tv_id,
			season_number,
			episode_number,
			api_key,
		)

		response, err := http.Get(apiUrl)
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Println("Error: API request returned status code", response.StatusCode)
			fmt.Println("Current Season Number:", season_number)
			fmt.Println("Current Episode Number:", episode_number)
			break
		}

		var episode EpisodeInfo
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&episode); err != nil {
			fmt.Println("Error decoding JSON response:", err)
			break
		}

		filename := fmt.Sprintf("%02d - %s", episode.EpisodeNumber, episode.Name)
		episode_map[fmt.Sprintf("%02d", episode.EpisodeNumber)] = filename

		episode_number++
	}

	return episode_map
}
