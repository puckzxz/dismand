package commands

import (
	"github.com/puckzxz/dismand"
)

func AddCommand(ctx *dismand.Context, args []string) {
	if len(args) != 2 {
		return
	}

	ctx.Reply("This command adds two numbers and return the sum")
}
