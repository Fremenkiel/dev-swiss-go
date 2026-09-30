package handlers

import (
	"bytes"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestAddCommand(t *testing.T) {
	t.Run("Add command", func(t *testing.T) {
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

	t.Run("Add command as child", func(t *testing.T) {
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

	t.Run("Add parent as child", func(t *testing.T) {
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

	t.Run("Add child two times", func(t *testing.T) {
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
	t.Run("Run command", func(t *testing.T) {
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

		args := []string{
			"run_command",
		}

		if err := c.RunCommand(args); err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("Run unknown command", func(t *testing.T) {
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

		args := []string{
			"unknown_command",
		}

		if err := c.RunCommand(args); err == nil {
			t.Error("Expected error, got none")
		}
	})
}


func TestExcecute(t *testing.T) {
	t.Run("Excecute", func(t *testing.T) {
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

		orgArgs := os.Args
		defer func() { os.Args = orgArgs }()

		orgStout := os.Stdout

		defer func() { os.Stdout = orgStout }()

		r, w, err := os.Pipe()
		if err != nil {
			t.Errorf("Unable to create pipe")
		}
		defer func() { 
			r.Close()
		}()

		var buf bytes.Buffer
		var wg sync.WaitGroup

		wg.Go(func() {
			io.Copy(&buf, r)
		})

		os.Args = []string{"devswiss", "run_command", "data.json"}
		os.Stdout = w

		if err := c.Execute(); err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		w.Close()
		wg.Wait()

		if strings.Contains(buf.String(), "helper") {
			t.Errorf("Error when executing, buffer string: %s", buf.String())
		}
	})

	t.Run("Excecute unknown command", func(t *testing.T) {
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

		originalArgs := os.Args
		defer func() { os.Args = originalArgs }()

		r, w, err := os.Pipe()
		if err != nil {
			t.Errorf("Unable to create pipe")
		}
		defer func() { 
			r.Close()
		}()

		var buf bytes.Buffer
		var wg sync.WaitGroup

		wg.Go(func() {
			io.Copy(&buf, r)
		})

		os.Args = []string{"devswiss", "unknown_command", "data.json"}
		os.Stdout = w

		if err := c.Execute(); err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		w.Close()
		wg.Wait()

		if !strings.Contains(buf.String(), "Usage") {
			t.Errorf("Error when executing, buffer string: %s", buf.String())
		}
	})
}
