package integrations

import "os"

// GitRemote represents, well... a git remote
type GitRemote struct {
	host    string
	project string
}

// GetGitRemote returns a GitRemote instance based on the environment variables found during runtime
func GetGitRemote() GitRemote {
	project, isGitHub := os.LookupEnv("GITHUB_REPOSITORY")

	if isGitHub {
		return GitRemote{
			// Assuming no self-hosted GitHub, not even sure if actions works there tbh
			host:    "github.com",
			project: project,
		}
	}

	// Otherwise it's GitLab
	return GitRemote{
		host:    os.Getenv("CI_SERVER_HOST"),
		project: os.Getenv("CI_PROJECT_PATH"),
	}
}

// GetRemoteURL builds an URL pointing to the specific project on the remote provider
func (remote GitRemote) GetRemoteURL() string {
	return "https://" + remote.host + "/" + remote.project
}

// Host returns the URL of the git host
func (remote GitRemote) Host() string {
	return remote.host
}

// Project returns the name of the project
func (remote GitRemote) Project() string {
	return remote.project
}