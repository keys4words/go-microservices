package github

type CreateRepoRequest struct {
	Name        string
	Description string
	Homepage    string
	Private     bool
	HasIssues   bool
	HasWiki     bool
}

func CreateRepo() {

}
