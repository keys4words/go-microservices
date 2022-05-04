package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"src/api/clients/restclient"
	"src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(token string) string {
	return fmt.Sprintf(headerAuthorizationFormat, token)
}

func CreateRepo(token string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(token))
	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Printf("error when trying to create a new repo at github: %s", err.Error())
		return nil, github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorRepsonse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorRepsonse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}
	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf(fmt.Sprintf("error when trying to unmarshal successful creation of new repo in github: %s", err.Error()))
		return nil, &github.GithubErrorRepsonse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error unmarshaling success  ",
		}
	}
	return &result, nil
}
