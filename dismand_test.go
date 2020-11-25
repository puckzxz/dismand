package dismand

import (
	"testing"
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

	d := &Disgoman{}
	ptr := d.RegisterDefaults()
	if ptr == nil {
		t.Error("RegisterDefaults returned a nil pointer")
	}

	_, ok := commands["ping"]
	if !ok {
		t.Error("Commands did not contain the default ping command")
	}
}
