package dismand

import (
	"strings"

	"github.com/andersfylling/disgord"
)

var commands = make(map[string]*cmd)

type command func(*Context, []string)

type cmd struct {
	c           command
	description string
	group       string
	example     string
	enabled     bool
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

type Dismand struct {
	commands map[string]*cmd
	client   *disgord.Client
	cfg      *Config
}

// Config contains settings for Dismand
type Config struct {
	Prefix string
}

// New returns a Dismand instance
func New(client *disgord.Client, cfg *Config) *Dismand {
	return &Dismand{
		client:   client,
		cfg:      cfg,
		commands: commands,
	}
}

// RegisterDefaults registers the default Dismand commands
//
// - Ping
func (d *Dismand) RegisterDefaults() *Dismand {
	commands["ping"] = &cmd{
		c:           pingPong,
		description: "Ping Pong",
		example:     "ping",
		group:       "Default",
		enabled:     true,
	}
	commands["help"] = &cmd{
		c:           help,
		description: "Shows information about a command",
		example:     "help ping",
		group:       "Default",
		enabled:     true,
	}
	commands["enable"] = &cmd{
		c:           enableCommand,
		description: "Enables a command",
		example:     "enable ping",
		group:       "Default",
		enabled:     true,
	}
	commands["disable"] = &cmd{
		c:           disableCommand,
		description: "Disables a command",
		example:     "disable ping",
		group:       "Default",
		enabled:     true,
	}
	return d
}

func (d *Dismand) On(command string, handler command) *cmd {
	c := &cmd{
		c:           handler,
		description: "No description provided",
		example:     "No example provided",
		group:       "None",
		enabled:     true,
	}
	d.commands[command] = c
	return c
}

func (d *Dismand) MessageHandler(s disgord.Session, evt *disgord.MessageCreate) {
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
		if cmd.enabled {
			cmd.c(ctx, args)
		}
	}
}
