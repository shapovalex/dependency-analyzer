package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func FindLicenseFile(repoUrl string) string {
	if strings.Contains(repoUrl, "github.com") {
		repoContent := GetRepoContent(repoUrl)
		for _, contentElement := range repoContent {
			if strings.Contains(contentElement.Name, "LICENSE") {
				return contentElement.Html_url
			}
		}
	}
	return "Unable to find license file"
}

func GetRepoContent(repoUrl string) []ContentResponseElement {
	client := &http.Client{}
	api_url := GetGithubReposApiUrl(repoUrl) + "/contents"
	req, err := http.NewRequest("GET", api_url, nil)
	if err != nil {
		fmt.Println("Unable to get repoContent, error during creating http request")
		return []ContentResponseElement{}
	}
	if os.Getenv("GITHUB_USER") != "" && os.Getenv("GITHUB_TOKEN") != "" {
		req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_TOKEN"))
	}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to get repoContent, error during http request")
		return []ContentResponseElement{}
	}
	if response.StatusCode != 200 {
		fmt.Println("Unable to get repoContent, invalid response code ", response.StatusCode, " ", api_url)
		return []ContentResponseElement{}
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Unable to get repoContent, error during reading data")
		return []ContentResponseElement{}
	}

	var content []ContentResponseElement
	err = json.Unmarshal(body, &content)
	if err != nil {
		fmt.Println("Unable to get repoContent", repoUrl, ". Error during parsing json", err.Error())
		return []ContentResponseElement{}
	}

	return content
}

func GetGithubReposApiUrl(repoUrl string) string {
	gitHubAddress := "github.com"
	addressStartIndex := strings.Index(repoUrl, gitHubAddress)
	repoUrl = repoUrl[addressStartIndex:]

	for strings.Count(repoUrl, "/") > 2 {
		repoUrl = repoUrl[:strings.LastIndex(repoUrl, "/")]
	}

	return "https://api." + repoUrl[:len(gitHubAddress)] + "/repos" + repoUrl[len(gitHubAddress):]
}
