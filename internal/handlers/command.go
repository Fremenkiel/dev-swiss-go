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
	Flags				*FlagContainer
	parent			*Command
	children		map[string]*Command
}

var (
	help bool
)

func NewCommand(name, description string, run func(args []string) error) *Command {
	fc := NewFlagContainer()
	fc.Boolean(&help, "help", "h", false, fmt.Sprintf("help for %s", name))

	return &Command{
		Name: name,
		Description: description,
		Run: run,
		Flags: fc,
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
		return fmt.Errorf("Command \"%s\" do not exist", args[0])
	}

	if len(args) > 1 && args[1][0] == '-' && help {
		command.printHelper(nil)
		return nil
	}

	command.Run(args[1:])

	if len(args) > 1 && args[1][0] != '-' {
		return command.RunCommand(args[1:])
	}

	return nil
}

func (c *Command) Execute() error {
	if len(os.Args) < 2 {
		c.printHelper(nil)
		return nil
	}

	if err := c.RunCommand(os.Args[1:]); err != nil {
		c.printHelper(err)
		return nil
	}

	return nil
}

func (c *Command) printHelper(err error) {
	fmt.Fprintln(os.Stdout, c.Description)
	fmt.Fprintln(os.Stdout)

	if err != nil {
		fmt.Fprintln(os.Stdout, "Error:")
		fmt.Fprintf(os.Stdout, "	An error was thrown: %v\n", err)
		fmt.Fprintln(os.Stdout)
	}

	fmt.Fprintln(os.Stdout, "Usage:")
	fmt.Fprintf(os.Stdout, "	%s [command]\n", c.Name)
	fmt.Fprintln(os.Stdout)

	fmt.Fprintln(os.Stdout, "Available Commands:")
	for i := range c.children {
		fmt.Fprintf(os.Stdout, "	%s - %s\n", c.children[i].Name, c.children[i].Description)
	}
	fmt.Fprintln(os.Stdout)

	fmt.Fprintln(os.Stdout, "Flags:")
	// c.Flags.Print()
	fmt.Fprintln(os.Stdout)

	fmt.Fprintf(os.Stdout, "Use \"%s [command] --help\" for more information about a command.\n", c.Name)
}
