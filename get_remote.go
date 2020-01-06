package integrations

import "os"

// GitRemote represents, well... a git remote
type GitRemote struct {
	Host    string
	Project string
}

// GetGitRemote returns a GitRemote instance based on the environment variables found during runtime
func GetGitRemote() GitRemote {
	project, isGitHub := os.LookupEnv("GITHUB_REPOSITORY")

	if isGitHub {
		return GitRemote{
			// Assuming no self-hosted GitHub, not even sure if actions works there tbh
			Host:    "github.com",
			Project: project,
		}
	}

	// Otherwise it's GitLab
	return GitRemote{
		Host:    os.Getenv("CI_SERVER_HOST"),
		Project: os.Getenv("CI_PROJECT_PATH"),
	}
}

// GetRemoteURL builds an URL pointing to the specific project on the remote provider
func (remote GitRemote) GetRemoteURL() string {
	return "https://" + remote.Host + "/" + remote.Project
}
