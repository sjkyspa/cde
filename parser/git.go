package parser

import (
	"fmt"

	docopt "github.com/sjkyspa/stacks/Godeps/_workspace/src/github.com/docopt/docopt-go"
	"github.com/sjkyspa/stacks/client/cmd"
)

// Git routes git commands to their specific function.
func Git(argv []string) error {
	usage := `
Valid commands for git:

git:remote          Adds git remote of application to repository

Use 'deis help [command]' to learn more.
`

	switch argv[0] {
	case "git:remote":
		return gitRemote(argv)
	case "git":
		fmt.Print(usage)
		return nil
	default:
		PrintUsage()
		return nil
	}
}

func gitRemote(argv []string) error {
	usage := `
Adds git remote of application to repository

Usage: cde git:remote <app> [options]

Options:
  app
    the uniquely identifiable name for the application.
  -r --remote=REMOTE
    name of remote to create. [default: cde]
`

	args, err := docopt.Parse(usage, argv, true, "", false, true)

	if err != nil {
		return err
	}

	return cmd.GitRemote(safeGetValue(args, "<app>"), args["--remote"].(string))
}
