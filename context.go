package dismand

import (
	"errors"
	"fmt"

	"github.com/andersfylling/disgord"
)

var (
	// ErrRoleNotFound is returned when the searched role is not found
	ErrRoleNotFound = errors.New("role not found")
)

// Context contains the recieved message, the Disgord client, and the active session
type Context struct {
	Message *disgord.Message
	Client  *disgord.Client
	Session disgord.Session
}

// Reply will reply to the user that sent the message
//
// It WILL mention the user!
func (c *Context) Reply(msg string) (*disgord.Message, error) {
	return c.Client.SendMsg(c.Message.ChannelID, fmt.Sprintf("<@%s>, %s", c.Message.Author.ID, msg))
}

// MemberHasPermission will check if the user has the specified permission
func (c *Context) MemberHasPermission(permission disgord.PermissionBit) (bool, error) {
	perms, err := c.Client.Guild(c.Message.GuildID).Member(c.Message.Author.ID).GetPermissions()
	if err != nil {
		return false, err
	}

	return perms.Contains(permission), nil
}

// MemberHasRole will check if the user has the specified role
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

// GetRoleByName will try to find the role in the server with the same name.
// If multiple roles have the same name it will return the first role.
// Returns ErrRoleNotFound if the role was not found.
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

	return nil, ErrRoleNotFound
}

// GetRoleByID will find the role in the server with the same ID.
// Returns ErrRoleNotFound if the role was not found.
func (c *Context) GetRoleByID(id disgord.Snowflake) (*disgord.Role, error) {
	roles, err := c.Client.Guild(c.Message.GuildID).GetRoles()

	if err != nil {
		return nil, err
	}

	for _, r := range roles {
		if r.ID == id {
			return r, nil
		}
	}

	return nil, ErrRoleNotFound
}
