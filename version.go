package main

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed templates/html/version.html
var versionTemplate string

func handleVersion(w http.ResponseWriter, r *http.Request) {
	type SiteInfo struct {
		CommitId string
		Name     string
		Version  string
	}

	tmpl, _ := template.New("version.html").Parse(versionTemplate)
	// Return (write) the version to the response body
	tmpl.Execute(w, SiteInfo{
		CommitId: COMMIT_ID,
		Name:     "YASKM",
		Version:  APP_VERSION,
	})
}
