package github

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
)

var Cmd = &Z.Cmd{
	Name:     `github`,
	Version:  `v0.2.0`,
	Aliases:  []string{`gh`},
	Summary:  `github utilities`,
	Commands: []*Z.Cmd{help.Cmd, latestCmd},
}

var latestCmd = &Z.Cmd{
	Name:    `latest`,
	Summary: `latest release name`,
	Description: `
		The {{cmd .Name}} command returns the name of the
		latest release for the specified GitHub user/org and repo (ex:
		docker/compose).
	`,

	NumArgs:  1,
	Commands: []*Z.Cmd{help.Cmd},
	Call: func(_ *Z.Cmd, args ...string) error {
		latest, err := Latest(args[0])
		if err != nil {
			return err
		}
		term.Print(latest)
		return nil
	},
}
