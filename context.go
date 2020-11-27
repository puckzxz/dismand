package dismand

import (
	"errors"
	"fmt"

	"github.com/andersfylling/disgord"
)

var (
	ERR_ROLE_NOT_FOUND = errors.New("Role not found")
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
	perms, err := c.Client.Guild(c.Message.GuildID).Member(c.Message.Author.ID).GetPermissions()
	if err != nil {
		return false, err
	}

	return perms.Contains(permission), nil
}

func (c *Context) MemberHasRole(role *disgord.Role) (bool, error) {
	member, err := c.Client.Guild(c.Message.GuildID).Member(c.Message.Author.ID).Get()

	if err != nil {
		return false, err
	}

	for _, r := range member.Roles {
		if r == role.ID {
			return true, nil
		}
	}

	return false, nil
}

func (c *Context) GetRoleByName(name string) (*disgord.Role, error) {
	roles, err := c.Client.Guild(c.Message.GuildID).GetRoles()

	if err != nil {
		return nil, err
	}

	for _, r := range roles {
		if r.Name == name {
			return r, nil
		}
	}

	return nil, ERR_ROLE_NOT_FOUND
}
