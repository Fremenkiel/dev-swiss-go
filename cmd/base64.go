package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var base64EncodeCmd = &cobra.Command{
	Use:		"base64-encode",
	Short:	"",
	Long:		"",
	RunE:		func(cmd *cobra.Command, args []string) error {
		return runBase64Encode(args)
	},
}

var base64DecodeCmd = &cobra.Command{
	Use:		"base64-decode",
	Short:	"",
	Long:		"",
	RunE:		func(cmd *cobra.Command, args []string) error {
		return runBase64Decode(args)
	},
}

func init() {
	rootCmd.AddCommand(base64EncodeCmd)
	rootCmd.AddCommand(base64DecodeCmd)
}

func runBase64Encode(args []string) error {
	if len(args) == 0 || len(args[0]) == 0 {
		fmt.Fprintln(os.Stderr, "invalid input")
		os.Exit(1)
	}

	bStr := base64.RawStdEncoding.EncodeToString([]byte(args[0]))

	fmt.Fprintln(os.Stdout, bStr)

	return nil
}

func runBase64Decode(args []string) error {
	if len(args) == 0 || len(args[0]) == 0 {
		fmt.Fprintln(os.Stderr, "invalid input")
		os.Exit(1)
	}

	str, err := base64.RawStdEncoding.DecodeString(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stdout, string(str))
	
	return nil
}
