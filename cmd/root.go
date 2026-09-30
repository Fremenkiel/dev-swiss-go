package cmd

import (
	"fmt"
	"os"

	"github.com/fremenkiel/dev-swiss-go/internal/handlers"
)

var rootCmd = handlers.NewCommand("devswiss", "A combination as everyday dev tools.", nil)

func Excecute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
