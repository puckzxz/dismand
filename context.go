package dismand

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

type Context struct {
	Message *disgord.Message
	Client  *disgord.Client
	Session disgord.Session
}

func (c *Context) Reply(msg string) (*disgord.Message, error) {
	return c.Client.SendMsg(c.Message.ChannelID, fmt.Sprintf("<@%s>, %s", c.Message.Author.ID, msg))
}

func (c *Context) MemberHasPermission(permission disgord.PermissionBit) (bool, error) {
	perms, err := c.Client.Guild(c.Message.GuildID).GetMemberPermissions(c.Message.Author.ID)
	if err != nil {
		return false, err
	}

	return perms.Contains(permission), nil
}
