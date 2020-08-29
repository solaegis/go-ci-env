package cienv_test

import (
	"testing"

	"github.com/suzuki-shunsuke/go-ci-env/cienv"
)

func TestPlatform(t *testing.T) {
	platform := cienv.Get()
	if platform == nil {
		return
	}

	t.Run("Platform.CI", func(t *testing.T) {
		if platform.CI() == "" {
			t.Error("platform.CI() is empty")
		}
	})

	t.Run("Platform.RepoOwner", func(t *testing.T) {
		if owner := platform.RepoOwner(); owner != "suzuki-shunsuke" {
			t.Error("RepoOwner = "+owner, ", wanted suzuki-shunsuke")
		}
	})

	t.Run("Platform.RepoName", func(t *testing.T) {
		if repo := platform.RepoName(); repo != "go-ci-env" {
			t.Error("RepoName = "+repo, ", wanted go-ci-env")
		}
	})

	t.Run("Platform.SHA1", func(t *testing.T) {
		if platform.SHA1() == "" {
			t.Error("platform.SHA1() is empty")
		}
	})

	t.Run("Platform.PRNumber", func(t *testing.T) {
		if !platform.IsPR() {
			return
		}
		num, err := platform.PRNumber()
		if err != nil {
			t.Error(err)
			return
		}
		if num == 0 {
			t.Error("PRNumber() == 0")
		}
	})
}
