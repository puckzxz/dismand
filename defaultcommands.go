package dismand

import (
	"fmt"
)

func pingPong(ctx *Context, args []string) {
	ctx.Reply("Pong!")
}

func help(ctx *Context, args []string) {
	switch len(args) {
	case 0:
		msg := "```\n"
		counter := 1
		for n, c := range commands {
			msg += fmt.Sprintf("%d. %s - %s\n\t%s\n\t%s\n\n", counter, n, c.group, c.description, c.example)
			counter++
		}
		msg += "\n```"
		ctx.Reply(msg)
	case 1:
		cmdName := args[0]
		if c, found := commands[cmdName]; found {
			msg := "```\n"
			msg += fmt.Sprintf("%s - %s\n\t%s\n\t%s", cmdName, c.group, c.description, c.example)
			msg += "\n```"
			ctx.Reply(msg)
		}
	default:
		return
	}
}

func enableCommand(ctx *Context, args []string) {
	if len(args) != 1 {
		return
	}

	cmdName := args[0]

	if c, found := commands[cmdName]; found {
		c.enabled = true
		ctx.Reply(fmt.Sprintf("Enabled `%s`", cmdName))
	}
}

func disableCommand(ctx *Context, args []string) {
	if len(args) != 1 {
		return
	}

	cmdName := args[0]

	if cmdName == "disable" {
		ctx.Reply("You can't disable the `disable` command!")
		return
	}

	if c, found := commands[cmdName]; found {
		c.enabled = false
		ctx.Reply(fmt.Sprintf("Disabled `%s`", cmdName))
	}
}
