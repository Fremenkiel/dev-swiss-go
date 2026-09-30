package handlers

import (
	"errors"
	"fmt"
	"os"
)

type Command struct {
	Name				string
	Description	string
	Run					func(args []string) error
	parent			*Command
	children		map[string]*Command
}

func NewCommand(name, description string, run func(args []string) error) *Command {
	return &Command{
		Name: name,
		Description: description,
		Run: run,
		children: make(map[string]*Command),
	}
}

func (c *Command) AddCommand(cmds ...*Command) {
	for i := range cmds {
		if cmds[i] == c {
			panic("Command can't be children of it self");
		}

		if cmds[i] == c.parent {
			panic("Command can't be children and parent");
		}

		if _, ok := c.children[cmds[i].Name]; ok == true {
			panic("A child with the same name already added");
		}

		c.children[cmds[i].Name] = cmds[i]
		cmds[i].parent = c
	}
}

func (c *Command) RunCommand(args []string) error {
	if len(args) == 0 {
		return errors.New("No arguments sendt")
	}

	command, ok := c.children[args[0]]
	if ok != true {
		return errors.New("Command do not exist")
	}

	command.Run(args[1:])

	return nil
}

func (c *Command) Execute() error {
	if len(os.Args) < 2 {
		c.printHelper()
		return nil
	}

	if err := c.RunCommand(os.Args[1:]); err != nil {
		c.printHelper()
		return nil
	}

	return nil
}

func (c *Command) printHelper() {
	fmt.Fprintln(os.Stdout, c.Description)
	fmt.Fprintln(os.Stdout)

	fmt.Fprintln(os.Stdout, "Usage:")
	fmt.Fprintf(os.Stdout, "	%s [command]\n", c.Name)
	fmt.Fprintln(os.Stdout)

	fmt.Fprintln(os.Stdout, "Available Commands:")
	for i := range c.children {
		fmt.Fprintf(os.Stdout, "	%s - %s\n", c.children[i].Name, c.children[i].Description)
	}
	fmt.Fprintln(os.Stdout)

	fmt.Fprintln(os.Stdout, "Flags:")
	fmt.Fprintf(os.Stdout, "	-h, --help - help for %s\n", c.Name)
	fmt.Fprintln(os.Stdout)

	fmt.Fprintf(os.Stdout, "Use \"%s [command] --help\" for more information about a command.\n", c.Name)
}
