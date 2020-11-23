package main

import (
	"os"

	"github.com/andersfylling/disgord"
	"github.com/puckzxz/dismand"
	"github.com/puckzxz/dismand/example/commands"
)

func main() {
	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("TOKEN"),
		RejectEvents: []string{
			disgord.EvtPresenceUpdate,
			disgord.EvtGuildMemberAdd,
			disgord.EvtGuildMemberUpdate,
			disgord.EvtGuildMemberRemove,
		},
	})

	d := dismand.New(client, &dismand.Config{
		Prefix: "!",
	}).RegisterDefaults()

	defer client.Gateway().StayConnectedUntilInterrupted()

	d.On("add", commands.AddCommand).
		Description("Adds two numbers").
		Example("add one two").
		Group("Math")

	client.Gateway().MessageCreate(d.MessageHandler)
}
