package main

import (
	docopt "github.com/docopt/docopt-go"
	"github.com/josa42/git-interactive/commands"
	stringutils "github.com/josa42/go-stringutils"
)

func main() {
	usage := stringutils.TrimLeadingTabs(`
		Usage:
		  git-interactive commit
		  git-interactive rebase

		Options:
		  -h --help          Show this screen.
		  --version          Show version.
  `)

	arguments, _ := docopt.Parse(usage, nil, true, "Git Cleanup 0.1.0", false)

	cmdType := ""
	cmdTypes := []string{"commit", "rebase"}

	for _, key := range cmdTypes {
		if arguments[key] == true {
			cmdType = key
		}
	}

	switch cmdType {
	case "commit":
		commands.Commit()
	case "rebase":
		commands.Rebase()
	}
}
