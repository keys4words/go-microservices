package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "Golang introduction",
		Description: "my project to work with github api",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     false,
	}

	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))

	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target, request)

	// assert.EqualValues(t,
	// `{"name":"Golang introduction","description":"my project to work with github api","homepage":"https://github.com","private":true,"has_issues":false,"has_projects":false,"has_wiki":false}`,
	// string(bytes))
}
