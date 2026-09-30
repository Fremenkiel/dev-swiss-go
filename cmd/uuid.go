package cmd

import (
	"fmt"
	"os"

	"github.com/fremenkiel/dev-swiss-go/internal/handlers"
	"github.com/fremenkiel/dev-swiss-go/internal/helpers"
	"github.com/google/uuid"
)

var (
	count int16
	version int16
)

var uuidCmd = handlers.NewCommand("uuid", "",
	func(args []string) error {
		return runUuidGen()
	})

func init() {
	rootCmd.AddCommand(uuidCmd)

	// uuidCmd.Flags().Int16VarP(&count, "count", "c", 1, "Amount of generated UUID")
	// uuidCmd.Flags().Int16VarP(&version, "version", "v", 7, "UUID version")
}

func runUuidGen() error {
	for range count {
		var str = "";
		switch version {
			case 6:
			genUuid, err := uuid.NewV6()
			helpers.ExitOnError(err)

			str = genUuid.String()
			break;
			case 7:
			genUuid, err := uuid.NewV7()
			helpers.ExitOnError(err)

			str = genUuid.String()
			break;
		default:
			fmt.Fprintln(os.Stderr, "Invalid version")
			os.Exit(1)
			break;
		}


		fmt.Fprintln(os.Stdout, str)
	}

	return nil
}
