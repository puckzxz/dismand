package dismand

import (
	"strings"

	"github.com/andersfylling/disgord"
)

var commands = make(map[string]*cmd)

type Command func(*Context, []string)

type cmd struct {
	c           Command
	description string
	group       string
	example     string
}

func (c *cmd) Description(desc string) *cmd {
	c.description = desc
	return c
}

func (c *cmd) Group(group string) *cmd {
	c.group = group
	return c
}

func (c *cmd) Example(example string) *cmd {
	c.example = example
	return c
}

type Disgoman struct {
	commands map[string]*cmd
	client   *disgord.Client
	cfg      *Config
}

type Config struct {
	Prefix string
}

func New(client *disgord.Client, cfg *Config) *Disgoman {
	return &Disgoman{
		client:   client,
		cfg:      cfg,
		commands: commands,
	}
}

func (d *Disgoman) RegisterDefaults() *Disgoman {
	commands["ping"] = &cmd{
		c: func(ctx *Context, args []string) {
			ctx.Reply("Pong!")
		},
		description: "Ping Pong",
		example:     "ping",
		group:       "Default",
	}
	return d
}

func (d *Disgoman) On(command string, handler Command) *cmd {
	c := &cmd{
		c:           handler,
		description: "No description provided",
		example:     "No example provided",
		group:       "None",
	}
	d.commands[command] = c
	return c
}

func (d *Disgoman) MessageHandler(s disgord.Session, evt *disgord.MessageCreate) {
	msg := evt.Message
	if !strings.HasPrefix(msg.Content, d.cfg.Prefix) {
		return
	}

	msg.Content = strings.TrimPrefix(msg.Content, d.cfg.Prefix)

	values := strings.Split(msg.Content, " ")

	commandName := values[0]

	args := values[1:]

	ctx := &Context{
		Client:  d.client,
		Session: s,
		Message: msg,
	}

	if cmd, ok := d.commands[commandName]; ok {
		cmd.c(ctx, args)
	}
}
