package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "a golang introduction repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	// Marshal takes an input interface and attempt to create a valid json string
	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))

	assert.EqualValues(
		t,
		`{"name":"golang introduction","description":"a golang introduction repo","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`,
		string(bytes),
	)

	var target CreateRepoRequest

	// Unmarshal takes an bytes array and *pointer* that we're trying to fill using this json
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)
	assert.Equal(t, target.Name, request.Name)
	assert.Equal(t, target.Description, request.Description)
	assert.Equal(t, target.Private, request.Private)
	assert.Equal(t, target.Homepage, request.Homepage)
	assert.Equal(t, target.HasIssues, request.HasIssues)
	assert.Equal(t, target.HasProjects, request.HasProjects)
	assert.Equal(t, target.HasWiki, request.HasWiki)
}
