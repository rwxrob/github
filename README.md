# 🌳 GitHub Common API Requests

[![GoDoc](https://godoc.org/github.com/rwxrob/github?status.svg)](https://godoc.org/github.com/rwxrob/github)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)

Just a collection of my personal favorite and most needed GitHub API
queries. Most people will want to use the amazing [gh GitHub
CLI](https://github.com/cli/cli) instead (which can be extended for most
things, just not these).

## Install

This command can be installed as a standalone program or composed into a
Bonzai command tree.

Standalone

```
go install github.com/rwxrob/github/github@latest
```

Composed

```go
package z

import (
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/github"
)

var Cmd = &Z.Cmd{
	Name:     `z`,
	Commands: []*Z.Cmd{help.Cmd, github.Cmd},
}
```

## Tab Completion

To activate bash completion just use the `complete -C` option from your
`.bashrc` or command line. There is no messy sourcing required. All the
completion is done by the program itself.

```
complete -C github github
```

If you don't have bash or tab completion check use the shortcut
commands instead.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source
code of the application. See the source or run the program with help to
access it.

## Design Considerations

* **Subsets of full structs and interface accessors**

  Like the Kubernetes Kind project, often we don't need the entire
  struct as defined by the (changing) GitHub API itself. We don't want
  to get into rewriting the entire API (others have already done that).
  We just need isolated pieces of different structs and can
  incrementally add the things as we need. And we'll take unusual steps
  to wrap everything in idiomatic Go access functions so that code
  doesn't break when the structs change rapidly, because they will,
  mostly because of the choice not to implement the entire API from the
  beginning..

* **GitHub `Client` struct in `pkg`**

  It's probable that someone might want more than one GitHub client in
  memory at the same time. For example, while at work. One client might
  be connected to the external github.com site while another is pointing
  at one of potentially many internal GitHub Enterprise instances.
  Creating a full interface seems unnecessary, however. If and when that
  need arises we can do that. Only the methods of `github.client` will
  be expose, however, to avoid tight coupling (use `github.NewClient()` or
  `github.Client`, a default instance).
