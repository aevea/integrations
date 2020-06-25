package integrations

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentRef(t *testing.T) {
	// Github action specific environment variable
	os.Setenv("GITHUB_REF", "github-develop")

	actionCompareBranch := GetCurrentRef()

	assert.Equal(t, "github-develop", actionCompareBranch)

	os.Setenv("GITHUB_REF", "")

	// Gitlab specific environment variable
	os.Setenv("CI_COMMIT_REF_NAME", "gitlab-develop")

	gitlabCompareBranch := GetCurrentRef()

	assert.Equal(t, "gitlab-develop", gitlabCompareBranch)

	os.Setenv("CI_COMMIT_REF_NAME", "")

	// Drone specific environment variable
	os.Setenv("DRONE_COMMIT_REF", "drone-develop")

	droneCompareBranch := GetCurrentRef()

	assert.Equal(t, "drone-develop", droneCompareBranch)

	os.Setenv("DRONE_COMMIT_REF", "")

	// Should default to empty if no conditions are satisfied
	defaultMaster := GetCurrentRef()

	assert.Equal(t, "", defaultMaster)
}
