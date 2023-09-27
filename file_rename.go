package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func CheckIfFolderExists(folder string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		fmt.Println("Source folder does not exist:", err)
		return
	}
	fmt.Println("It exists!")
}

func RenameFolderEpisodes(folder string, episode_names map[string]string, regexp_pattern string) {
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		fmt.Println("Source folder does not exist:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		parentDir := filepath.Dir(path)
		if !strings.EqualFold(parentDir, folder) {
			return nil
		}

		filename := info.Name()

		matches := regexp.MustCompile(regexp_pattern).FindStringSubmatch(filename)

		if len(matches) < 2 || len(matches) > 3 {
			return nil
		}

		episode_number, _ := strconv.Atoi(matches[len(matches)-1])

		// fmt.Printf("Rename %s? (Y/N): ", filename)
		// response, _ := reader.ReadString('\n')
		// if response[0] != 'Y' && response[0] != 'y' {
		// 	return nil
		// }

		episode_name_key := fmt.Sprintf("%02d", episode_number)
		new_episode_name, exists := episode_names[episode_name_key]

		if !exists {
			fmt.Printf("Episode name not found for episode %02d\n", episode_number)
			return nil
		}

		new_filename := fmt.Sprintf(
			"%s%s",
			new_episode_name,
			filepath.Ext(filename),
		)
		new_file_path := filepath.Join(filepath.Dir(path), new_filename)

		if _, err := os.Stat(new_file_path); err == nil {
			fmt.Printf(
				"A file with the same name (%s) already exists in %s. Overwrite? (Y/N): ",
				new_file_path,
				folder,
			)
			response, _ := reader.ReadString('\n')
			if response[0] != 'Y' && response[0] != 'y' {
				return nil
			}
		}

		if err := os.Rename(path, new_file_path); err != nil {
			fmt.Printf("Error renaming file %s: %v\n", filename, err)
		} else {
			fmt.Printf("Renamed %s to %s\n", filename, new_filename)
		}

		return nil
	})
	if err != nil {
		fmt.Println("Error traversing source folder:", err)
	}
}
