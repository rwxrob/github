// Copyright 2022 github Robert Muhlestein.
// SPDX-License-Identifier: Apache-2.0

// Package github provides the Bonzai command branch of the same name.
package github

import (
	"log"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
)

// main branch
var Cmd = &Z.Cmd{

	Name:      `github`,
	Summary:   `common GitHub API requests`,
	Version:   `v0.0.1`,
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	License:   `Apache-2.0`,
	Source:    `git@github.com:rwxrob/github.git`,
	Issues:    `github.com/rwxrob/github/issues`,

	Commands: []*Z.Cmd{
		help.Cmd, vars.Cmd, conf.Cmd, // common
		release,
	},

	Description: `
		The {{cmd .Name}} branch contains commands for common GitHub API requests.`,
}

// branch
var release = &Z.Cmd{
	Name:     `release`,
	Aliases:  []string{`rel`},
	Summary:  `requests for GitHub releases`,
	Commands: []*Z.Cmd{help.Cmd, latestRelease},
}

// branch
var latestRelease = &Z.Cmd{
	Name:    `latest`,
	Aliases: []string{`last`},
	Summary: `get latest release`,

	Description: `
		The {{cmd .Name}} branch contains commands to deal with the latest
		GitHub release of a repo matching this executable ({{exe exename}}).
	`,

	Commands: []*Z.Cmd{help.Cmd, latestReleaseVersion},
}

// leaf
var latestReleaseVersion = &Z.Cmd{
	Name:     `version`,
	Aliases:  []string{`vers`},
	Commands: []*Z.Cmd{help.Cmd},
	Summary:  `request latest GitHub release version for {{exe exename}}`,
	Call: func(x *Z.Cmd, _ ...string) error {
		// TODO infer the release URL from the x.Source
		// latest := github.LatestVersion(x.Source)
		log.Print("would fetch latest version of this binary")
		return nil
	},
}
