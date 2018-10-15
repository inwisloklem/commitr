package tasks_test

import (
	"reflect"
	"testing"

	"commitr/tasks"
)

var (
	toLoad string   = `command 1 {{message }}\ncommand 2 {{ commit}}\ncommand 3`
	loaded []string = []string{"command 1 message", "commit 2 commit", "commang 3"}
)

func TestLoadCommands(t *testing.T) {
	if reflect.DeepEqual(tasks.LoadCommands(toLoad, "message", "commit"), loaded) {
		t.Error("Commands list doesn't match the expected output")
	}
}
