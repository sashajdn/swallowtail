package commands

import (
	"fmt"
	"sync"
)

var (
	registry = map[string]*Command{}
	mu       sync.RWMutex
)

func register(id string, command *Command) {
	if _, ok := registry[id]; ok {
		panic(fmt.Sprintf("Cannot register commands with the same ID: %s", id))
	}

	registry[id] = command
}

// List lists all commands registered to satoshi
func List() []*Command {
	mu.RLock()
	defer mu.RUnlock()

	commands := []*Command{}
	for _, c := range registry {
		commands = append(commands, c)
	}

	return commands
}