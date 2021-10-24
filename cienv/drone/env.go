package drone

import (
	"fmt"
	"strconv"
)

type Client struct {
	Getenv func(string) string
}

func (client Client) CI() string {
	return "drone"
}

func (client Client) Match() bool {
	return client.Getenv("DRONE") != ""
}

func (client Client) RepoOwner() string {
	return client.Getenv("DRONE_REPO_OWNER")
}

func (client Client) RepoName() string {
	return client.Getenv("DRONE_REPO_NAME")
}

func (client Client) Ref() string {
	return client.Getenv("DRONE_COMMIT_REF")
}

func (client Client) Tag() string {
	return client.Getenv("DRONE_TAG")
}

func (client Client) Branch() string {
	return client.Getenv("DRONE_SOURCE_BRANCH")
}

func (client Client) PRBaseBranch() string {
	return client.Getenv("DRONE_TARGET_BRANCH")
}

func (client Client) SHA() string {
	return client.Getenv("DRONE_COMMIT_SHA")
}

func (client Client) IsPR() bool {
	return client.Getenv("DRONE_PULL_REQUEST") != ""
}

func (client Client) PRNumber() (int, error) {
	pr := client.Getenv("DRONE_PULL_REQUEST")
	if pr == "" {
		return 0, nil
	}
	b, err := strconv.Atoi(pr)
	if err == nil {
		return b, nil
	}
	return 0, fmt.Errorf("DRONE_PULL_REQUEST is invalid. It failed to parse DRONE_PULL_REQUEST as an integer: %w", err)
}
