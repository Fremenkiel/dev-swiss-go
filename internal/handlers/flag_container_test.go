package handlers

import (
	"os"
	"testing"
)

func TestInt16(t *testing.T) {
	tests := []struct {
		name		string
		args		[]string
		expectedValue	int16
	}{
		{ 
			name: "parse_args",
			args: []string{"devswiss", "run_command", "-v", "1"},
			expectedValue: 1,
		},
		{
			name: "parse_args_full",
			args: []string{"devswiss", "run_command", "--value", "1"},
			expectedValue: 1,
		},
		{
			name: "parse_args_invalid",
			args: []string{"devswiss", "run_command", "--v", "1"},
			expectedValue: 0,
		},
		{
			name: "parse_args_full_invalid",
			args: []string{"devswiss", "run_command", "-value", "1"},
			expectedValue: 0,
		},
		{
			name: "parse_args_no_valid",
			args: []string{"devswiss", "run_command", "-x", "1"},
			expectedValue: 0,
		},
		{
			name: "parse_args_multiple",
			args: []string{"devswiss", "run_command", "-x", "1", "-v", "2", "-y", "3"},
			expectedValue: 2,
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			defer func(args []string, stdout *os.File) {
				cleanupEnv(args, stdout)
			}(os.Args, os.Stdout)

			os.Args = tests[i].args

			c := NewCommand("test_command", "", func(args []string) error { return nil })

			var value int16
			c.Flags.Int16(&value, "value", "v", 0, "set test value")

			if value != tests[i].expectedValue {
				t.Errorf("Unable to parse args, expected %d got %d", tests[i].expectedValue, value)
			}
		})
	}
}

func TestBoolean(t *testing.T) {
	tests := []struct {
		name		string
		args		[]string
		expectedValue	bool
	}{
		{ 
			name: "parse_args",
			args: []string{"devswiss", "run_command", "-f"},
			expectedValue: true,
		},
		{
			name: "parse_args_full",
			args: []string{"devswiss", "run_command", "--flag"},
			expectedValue: true,
		},
		{ 
			name: "parse_args_invalid",
			args: []string{"devswiss", "run_command", "--f"},
			expectedValue: false,
		},
		{
			name: "parse_args_full_invalid",
			args: []string{"devswiss", "run_command", "-flag"},
			expectedValue: false,
		},
		{
			name: "parse_args_no_valid",
			args: []string{"devswiss", "run_command", "-x", "1"},
			expectedValue: false,
		},
		{
			name: "parse_args_multiple",
			args: []string{"devswiss", "run_command", "-x", "1", "-f", "-y", "3"},
			expectedValue: true,
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			defer func(args []string, stdout *os.File) {
				cleanupEnv(args, stdout)
			}(os.Args, os.Stdout)

			os.Args = tests[i].args

			c := NewCommand("test_command", "", func(args []string) error { return nil })

			var value bool
			c.Flags.Boolean(&value, "flag", "f", false, "set test value")

			if value != tests[i].expectedValue {
				t.Errorf("Unable to parse args, expected %v got %v", tests[i].expectedValue, value)
			}
		})
	}
}


func TestString(t *testing.T) {
	tests := []struct {
		name		string
		args		[]string
		expectedValue	string
	}{
		{ 
			name: "parse_args",
			args: []string{"devswiss", "run_command", "-v", "test"},
			expectedValue: "test",
		},
		{
			name: "parse_args_full",
			args: []string{"devswiss", "run_command", "--value", "test"},
			expectedValue: "test",
		},
		{ 
			name: "parse_args_invalid",
			args: []string{"devswiss", "run_command", "--t", "test"},
			expectedValue: "",
		},
		{
			name: "parse_args_full_invalid",
			args: []string{"devswiss", "run_command", "-flag", "test"},
			expectedValue: "",
		},
		{
			name: "parse_args_no_valid",
			args: []string{"devswiss", "run_command", "-x", "test"},
			expectedValue: "",
		},
		{
			name: "parse_args_multiple",
			args: []string{"devswiss", "run_command", "-x", "1", "-v", "test", "-y", "3"},
			expectedValue: "test",
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			defer func(args []string, stdout *os.File) {
				cleanupEnv(args, stdout)
			}(os.Args, os.Stdout)

			os.Args = tests[i].args

			c := NewCommand("test_command", "", func(args []string) error { return nil })

			var value string
			c.Flags.String(&value, "value", "v", "", "set test value")

			if value != tests[i].expectedValue {
				t.Errorf("Unable to parse args, expected %v got %v", tests[i].expectedValue, value)
			}
		})
	}
}
