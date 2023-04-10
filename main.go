package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Get the problem number from the command line argument
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run script.go <problem number>")
		return
	}
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid problem number:", os.Args[1])
		return
	}
	// Use the leetcode API to get the problem details by number
	resp, err := http.Get("https://leetcode.com/api/problems/algorithms/?ids=1")
	if err != nil {
		fmt.Println("Error getting problem details:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	// Parse the JSON response and get the title
	var data struct {
		Questions []struct {
			Details struct {
				ID   int    `json:"question_id"`
				Slug string `json:"question__title_slug"`
			} `json:"stat"`
		} `json:"stat_status_pairs"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Error parsing JSON response:", err)
	}
	if len(data.Questions) == 0 {
		log.Fatalf("API didn't return any problems. They must have changed something. Time to debug! Raise an issue!")
	}
	// Map responses
	idToSlug := map[int]string{}
	for _, question := range data.Questions {
		idToSlug[question.Details.ID] = question.Details.Slug
	}
	slug := idToSlug[num]
	if slug == "" {
		log.Fatalf("API returned a mapping, but it didn't have the problem with id %d. Sorry!", num)
	}

	fmt.Printf("https://leetcode.com/problems/%s/\n", slug)
}
