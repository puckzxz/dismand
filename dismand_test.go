package dismand

import (
	"testing"

	"github.com/andersfylling/disgord"
)

func Test_cmd_Description(t *testing.T) {
	c := &cmd{}
	c.Description("test description")
	if c.description != "test description" {
		t.Error("Description was not set")
	}
}

func Test_cmd_Example(t *testing.T) {
	c := &cmd{}
	c.Example("test example")
	if c.example != "test example" {
		t.Error("Example was not set")
	}
}

func Test_cmd_Group(t *testing.T) {
	c := &cmd{}
	c.Group("test group")
	if c.group != "test group" {
		t.Error("Group was not set")
	}
}

func Test_DismandNew(t *testing.T) {
	d := New(nil, nil)

	if d == nil {
		t.Error("New didn't return a Dismand pointer")
	}
}

func Test_Disgoman_RegisterDefaults(t *testing.T) {

	d := &Dismand{}
	ptr := d.RegisterDefaults()
	if ptr == nil {
		t.Error("RegisterDefaults returned a nil pointer")
	}

	_, ok := commands["ping"]
	if !ok {
		t.Error("Commands did not contain the default ping command")
	}
}

func TestConfig_messageIsInAllowedChannel(t *testing.T) {
	d := &Dismand{}
	ptr := d.RegisterDefaults()
	if ptr == nil {
		t.Error("RegisterDefaults returned a nil pointer")
	}

	c := Config{}

	c.AllowedChannels = []uint64{123, 456}

	if c.messageIsInAllowedChannel(678) {
		t.Error("Failed to detect unallowed channel")
	}

	if !c.messageIsInAllowedChannel(123) {
		t.Error("Failed to detech allowed channel")
	}
}

func Test_cmd_MinPermission(t *testing.T) {
	c := cmd{}

	c.MinPermission(disgord.PermissionAll)

	if c.minPerm != disgord.PermissionAll {
		t.Error("Failed to set minimum permission")
	}
}

func Test_cmd_AllowedChannels(t *testing.T) {
	c := cmd{}

	c.AllowedChannels([]uint64{123, 456})

	if c.allowedChannels == nil {
		t.Error("Failed to set allowed channels")
	}

	if c.allowedChannels[0] != 123 {
		t.Error("Failed to set allowed channels")
	}

	if c.allowedChannels[1] != 456 {
		t.Error("Failed to set allowed channels")
	}
}
