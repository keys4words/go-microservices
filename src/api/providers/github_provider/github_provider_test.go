package github_provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthorizationHeader(t *testing.T) {
	header := getAuthorizationHeader("abc")
	assert.EqualValues(t, "token abc", header)
}
