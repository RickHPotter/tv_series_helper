package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"path/filepath"
// 	"regexp"
// 	"strconv"
// 	"strings"
// )
//
// func main() {
// 	sourceFolder := "D:/A THOUSAND LIVES/&VISUAL/TV SERIES/Other/[F] Supernatural [Pahe.in]/new/source/Season 01"
// 	episodeNamesFolder := "D:/A THOUSAND LIVES/&VISUAL/TV SERIES/Other/[F] Supernatural [Pahe.in]/new/dest/Season 01"
//
// 	// Check if the source folder and episode names folder exist
// 	if _, err := os.Stat(sourceFolder); os.IsNotExist(err) {
// 		fmt.Println("Source folder does not exist:", err)
// 		return
// 	}
//
// 	if _, err := os.Stat(episodeNamesFolder); os.IsNotExist(err) {
// 		fmt.Println("Episode names folder does not exist:", err)
// 		return
// 	}
//
// 	episodeNames := make(map[string]string)
// 	reader := bufio.NewReader(os.Stdin)
//
// 	// Read episode names from the episode names folder
// 	err := readEpisodeNames(episodeNamesFolder, episodeNames)
// 	if err != nil {
// 		fmt.Println("Error reading episode names:", err)
// 		return
// 	}
//
// 	// Traverse the source folder and its subfolders
// 	err = filepath.Walk(sourceFolder, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
//
// 		if info.IsDir() {
// 			return nil // Skip directories
// 		}
//
// 		filename := info.Name()
//
// 		matches := regexp.MustCompile(`S(\d+)E(\d+)`).FindStringSubmatch(filename)
//
// 		if len(matches) == 3 {
// 			seasonNumber, _ := strconv.Atoi(matches[1])
// 			episodeNumber, _ := strconv.Atoi(matches[2])
// 			fmt.Println("@@")
// 			fmt.Println(seasonNumber)
// 			fmt.Println(episodeNumber)
//
// 			seasonFolder := fmt.Sprintf("Season %02d", seasonNumber)
// 			fmt.Println(seasonFolder)
//
// 			// // Confirm if user wants to proceed with renaming
// 			// fmt.Printf("Rename %s in %s? (Y/N): ", filename, seasonFolder)
// 			// response, _ := reader.ReadString('\n')
// 			// if response[0] != 'Y' && response[0] != 'y' {
// 			// 	return nil // Skip this file
// 			// }
//
// 			episodeNameKey := fmt.Sprintf("%02d", episodeNumber)
// 			newEpisodeName, exists := episodeNames[episodeNameKey]
// 			fmt.Println(newEpisodeName)
//
// 			if !exists {
// 				fmt.Printf("Episode name not found for episode %02d\n", episodeNumber)
// 				return nil // Skip this file
// 			}
//
// 			newFilename := fmt.Sprintf(
// 				"%02d - %s%s",
// 				episodeNumber,
// 				newEpisodeName,
// 				filepath.Ext(filename),
// 			)
// 			newFilePath := filepath.Join(filepath.Dir(path), newFilename)
//
// 			// Check if the destination file already exists
// 			if _, err := os.Stat(newFilePath); err == nil {
// 				fmt.Printf(
// 					"A file with the same name already exists in %s. Overwrite? (Y/N): ",
// 					seasonFolder,
// 				)
// 				response, _ := reader.ReadString('\n')
// 				if response[0] != 'Y' && response[0] != 'y' {
// 					return nil // Skip this file
// 				}
// 			}
//
// 			err := os.Rename(path, newFilePath)
// 			if err != nil {
// 				fmt.Printf("Error renaming file %s: %v\n", filename, err)
// 			} else {
// 				fmt.Printf("Renamed %s to %s\n", filename, newFilename)
// 			}
// 		}
//
// 		return nil
// 	})
//
// 	if err != nil {
// 		fmt.Println("Error traversing source folder:", err)
// 	}
// }
//
// func readEpisodeNames(episodeNamesFolder string, episodeNames map[string]string) error {
// 	files, err := ioutil.ReadDir(episodeNamesFolder)
// 	if err != nil {
// 		return err
// 	}
//
// 	for _, file := range files {
// 		filename := file.Name()
// 		parts := strings.SplitN(filename, " - ", 2)
// 		name := strings.Split(parts[1], ".") // there might be a . in the name itself, so i need all but the last .*
// 		episodeNumber := fmt.Sprintf("%02s", parts[0])
// 		episodeName := name[all but last part that contains the format]
//
// 		episodeNames[episodeNumber] = episodeName
// 	}
//
// 	return nil
// }
