package github_provider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAuthorizatoinHeader(t *testing.T) {
	header := getAuthorizationHeader("test1234")
	fmt.Println(header)
	assert.EqualValues(t, header, "token test1234")
}
