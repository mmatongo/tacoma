package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// create a struct that holds information to be displayed when type is called
type GitHub struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
	Url         string `json:"html_url"`
	Hireable    bool   `json:"hireable"`
}

func getUserInfo(username string) *GitHub {
	url := "https://api.github.com/users/" + username
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var user GitHub
	json.Unmarshal(body, &user)
	return &user
}

func getUsername() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	} else {
		fmt.Print("Enter a GitHub username: ")
		var username string
		fmt.Scanln(&username)
		return username
	}
}

func drawBox(text string) {
	lines := strings.Split(text, "\n")
	max := 0
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}
	fmt.Println("┌" + strings.Repeat("─", max) + "┐")
	for _, line := range lines {
		fmt.Println("│" + line + strings.Repeat(" ", max-len(line)) + "│")
	}
	fmt.Println("└" + strings.Repeat("─", max) + "┘")
}

func main() {
	username := getUsername()
	drawBox(fmt.Sprintf("Name: %s\nPublic Repos: %d\nFollowers: %d\nFollowing: %d\nHireable: %t\nURL: %s", getUserInfo(username).Name, getUserInfo(username).PublicRepos, getUserInfo(username).Followers, getUserInfo(username).Following, getUserInfo(username).Hireable, getUserInfo(username).Url))
}
