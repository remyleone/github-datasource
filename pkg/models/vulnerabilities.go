package models

// ListVulnerabilitiesOptions is provided when listing volnerabilities in a repository
type ListVulnerabilitiesOptions struct {
	// Repository is the name of the repository being queried (ex: grafana)
	Repository string `json:"repository"`

	// Owner is the owner of the repository (ex: grafana)
	Owner string `json:"owner"`

	// Query searches x by name and description
	Query string `json:"query"`
}
