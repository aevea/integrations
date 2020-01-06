package integrations

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGitRemote(t *testing.T) {
	var gitRemote GitRemote

	os.Setenv("GITHUB_REPOSITORY", "outillage/integrations")
	gitRemote = GetGitRemote()

	assert.Equal(t, "github.com", gitRemote.Host)
	assert.Equal(t, "outillage/integrations", gitRemote.Project)

	os.Clearenv()

	os.Setenv("CI_SERVER_HOST", "gitlab.com")
	os.Setenv("CI_PROJECT_PATH", "outillage/integrations")
	gitRemote = GetGitRemote()

	assert.Equal(t, "gitlab.com", gitRemote.Host)
	assert.Equal(t, "outillage/integrations", gitRemote.Project)
}
