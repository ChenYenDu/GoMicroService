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

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))
	response, err := restclient.Post(urlCreateRepo, request, headers)

	if err != nil {
		log.Printf("Error when trying to create new repo in github: %s", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)

	// read byte from the response body
	if err != nil {
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid response body",
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errResponse github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GithubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "Invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	var result github.CreateRepoResponse

	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Printf("Error when trying to unmarshal create repo successful response: %s", err.Error())
		return nil, &github.GithubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error when trying to unmarshal github create repo response",
		}
	}

	return &result, nil
}
