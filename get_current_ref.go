package integrations

import "os"

// GetCurrentRef will return the Reference provided by gitlab, github or drone.
func GetCurrentRef() string {
	gitlabRef := os.Getenv("CI_COMMIT_REF_NAME")

	if gitlabRef != "" {
		return gitlabRef
	}

	githubRef := os.Getenv("GITHUB_REF")

	if githubRef != "" {
		return githubRef
	}

	droneRef := os.Getenv("DRONE_COMMIT_REF")

	if droneRef != "" {
		return droneRef
	}

	return ""
}
