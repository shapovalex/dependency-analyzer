package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.Equal(t, "https://api.github.com/repos/django/django", GetGithubReposApiUrl("https://github.com/django/django"))
	assert.Equal(t, "https://api.github.com/repos/django/django", GetGithubReposApiUrl("http://github.com/django/django"))
	assert.Equal(t, "https://api.github.com/repos/django/django", GetGithubReposApiUrl("github.com/django/django"))
}
