package handlers

import (
	"os"
	"strings"
	"testing"
)

func TestAddCommand(t *testing.T) {
	t.Run("add_command", func(t *testing.T) {
		c := &Command{
			Name: "test_command",
			Description: "A command to test on",
			Run: func(args []string) error { return nil },
			parent: nil,
			children: make(map[string]*Command),
		}

		c.AddCommand(
			&Command{
				Name: "child_command_1",
				Description: "The first child command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			},
			&Command{
				Name: "child_command_2",
				Description: "The secound child command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			},
			&Command{
				Name: "child_command_3",
				Description: "The third child command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			},
			)
	})

	t.Run("add_command_as_child", func(t *testing.T) {
		c := &Command{
			Name: "test_command",
			Description: "A command to test on",
			Run: func(args []string) error { return nil },
			parent: nil,
			children: make(map[string]*Command),
		}

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Code did not paniced")
			}
		}()

		c.AddCommand(
			&Command{
				Name: "child_command_1",
				Description: "The first child command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			},
			&Command{
				Name: "child_command_2",
				Description: "The secound child command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			},
			c,
			)
	})

	t.Run("add_parent_as_child", func(t *testing.T) {
		c := &Command{
			Name: "test_command",
			Description: "A command to test on",
			Run: func(args []string) error { return nil },
			parent: &Command{
				Name: "parent_command_1",
				Description: "The parent command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			},
			children: make(map[string]*Command),
		}

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Code did not paniced")
			}
		}()

		c.AddCommand(
			&Command{
				Name: "child_command_1",
				Description: "The first child command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			},
			&Command{
				Name: "child_command_2",
				Description: "The secound child command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			},
			c.parent,
			)
	})

	t.Run("add_child_two_times", func(t *testing.T) {
		c := &Command{
			Name: "test_command",
			Description: "A command to test on",
			Run: func(args []string) error { return nil },
			parent: nil,
			children: map[string]*Command{
				"child_command_1": {
				Name: "child_command_1",
				Description: "The first child command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			}},
		}

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Code did not paniced")
			}
		}()

		c.AddCommand(
			&Command{
				Name: "child_command_2",
				Description: "The secound child command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			},
			&Command{
				Name: "child_command_3",
				Description: "The third child command",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: make(map[string]*Command),
			},
			c.children["child_command_1"],
			)
	})
}

func TestRunCommand(t *testing.T) {
	tests := []struct{
		name				string
		args				[]string
		child				func() *Command
		expectHelp	bool
		expectError	bool
	}{
		{
			name: "run_command",
			args: []string{"devswiss", "run_command"},
			child: func() *Command {
				return &Command{
					Name: "child_command_1",
					Description: "The first child command",
					Run: func(args []string) error { return nil },
					parent: nil,
					children: make(map[string]*Command),
				}
			},
			expectHelp: false,
			expectError: false,
		},
		{
			name: "run_unknown_command", 
			args: []string{"devswiss", "unknown_command"},
			child: func() *Command {
				return &Command{
					Name: "child_command_1",
					Description: "The first child command",
					Run: func(args []string) error { return nil },
					parent: nil,
					children: make(map[string]*Command),
				}
			},
			expectHelp: false,
			expectError: true,
		},
		{
			name: "run_command_help_flag",
			args: []string{"devswiss", "run_command", "-h"},
			child: func() *Command {
				return NewCommand("run_command", "The first child command", func(args []string) error { return nil })
			},
			expectHelp: true,
			expectError: false,
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			defer func(args []string, stdout *os.File) {
				cleanupEnv(args, stdout)
			}(os.Args, os.Stdout)

			r, w, wg, buf := setupEnv(t, tests[i].args)
			defer func() {
				r.Close()
			}()

			c := &Command{
				Name: "test_command",
				Description: "A command to test on",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: map[string]*Command{
					"run_command": tests[i].child(),
				},
			}

			if err := c.RunCommand(os.Args[1:]); err != nil && !tests[i].expectError {
				t.Errorf("Unexpected error: %v", err)
			}

			w.Close()
			wg.Wait()

			if strings.Contains(buf.String(), "Usage") != tests[i].expectHelp {
				t.Errorf("Incorrect stdout output, expected %v, buffer string: %s", tests[i].expectHelp, buf.String())
			}
		})
	}
}

func TestExcecute(t *testing.T) {
	tests := []struct{
		name				string
		args				[]string
		expectHelp	bool
	}{
		{
			name: "excecute",
			args: []string{"devswiss", "run_command"},
			expectHelp: false,
		},
		{
			name: "excecute_unknown_command",
			args: []string{"devswiss", "unknown_command"},
			expectHelp: true,
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			defer func(args []string, stdout *os.File) {
				cleanupEnv(args, stdout)
			}(os.Args, os.Stdout)

			r, w, wg, buf := setupEnv(t, tests[i].args)
			defer func() {
				r.Close()
			}()

			c := &Command{
				Name: "test_command",
				Description: "A command to test on",
				Run: func(args []string) error { return nil },
				parent: nil,
				children: map[string]*Command{
					"run_command": {
						Name: "child_command_1",
						Description: "The first child command",
						Run: func(args []string) error { return nil },
						parent: nil,
						children: make(map[string]*Command),
					},
				},
			}

			if err := c.Execute(); err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			w.Close()
			wg.Wait()

			if strings.Contains(buf.String(), "Usage") != tests[i].expectHelp {
				t.Errorf("Error when executing, buffer string: %s", buf.String())
			}
		})
	}
}
