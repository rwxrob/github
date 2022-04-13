package github_test

import (
	"fmt"

	github "github.com/rwxrob/github/pkg"
)

func ExampleClient_defaults() {
	gh := github.NewClient()
	fmt.Println(gh.Host())
	fmt.Println(gh.APIVersion())
	fmt.Println(gh.APIURL(`rwxrob/web`))
	// Output:
	// github.com
	// v3
	// https://api.github.com/rwxrob/web
}

func ExampleClient_APIURL() {
	gh := github.NewClient()
	fmt.Println(gh.APIURL(`rwxrob/web`))
	gh.SetHost(`github.example.com`)
	fmt.Println(gh.APIURL(`rwxrob/web`))
	// Output:
	// https://api.github.com/rwxrob/web
	// https://github.example.com/api/v3/rwxrob/web
}

func ExampleClient_Repo() {
	gh := github.NewClient()
	repo, err := gh.Repo(`rwxrob/web`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(repo["name"])
	// Output:
	// web
}
