package handlers

import (
	"fmt"
	"os"
	"strconv"
)

type FlagContainer struct {
	flags	[]*Flag
}

func NewFlagContainer() *FlagContainer {
	return &FlagContainer{
		flags: make([]*Flag,0),
	}
}

func (c *FlagContainer) Int16(value *int16, flag, shorthand string, defaultValue int16, description string) {
	c.flags = append(c.flags, &Flag{
		Flag: flag,
		Shorthand: shorthand,
		Description: description,
	})

	for i := range os.Args {
		if os.Args[i] == fmt.Sprintf("--%s", flag) {
			iv, err := strconv.ParseInt(os.Args[i+1], 0, 16)
			if err == nil {
				*value = int16(iv)
				return
			}
		}

		if os.Args[i] == fmt.Sprintf("-%s", shorthand) {
			iv, err := strconv.ParseInt(os.Args[i+1], 0, 16)
			if err == nil {
				*value = int16(iv)
				return
			}
		}
	}
	*value = defaultValue
}

func (c *FlagContainer) String(value *string, flag, shorthand string, defaultValue string, description string) {
	c.flags = append(c.flags, &Flag{
		Flag: flag,
		Shorthand: shorthand,
		Description: description,
	})

	for i := range os.Args {
		if os.Args[i] == fmt.Sprintf("--%s", flag) {
			*value = os.Args[i+1]
			return
		}

		if os.Args[i] == fmt.Sprintf("-%s", shorthand) {
			*value = os.Args[i+1]
			return
		}
	}
	*value = defaultValue
}

func (c *FlagContainer) Boolean(value *bool, flag, shorthand string, defaultValue bool, description string) {
	c.flags = append(c.flags, &Flag{
		Flag: flag,
		Shorthand: shorthand,
		Description: description,
	})

	for i := range os.Args {
		if os.Args[i] == fmt.Sprintf("--%s", flag) {
			*value = true
			return
		}

		if os.Args[i] == fmt.Sprintf("-%s", shorthand) {
			*value = true
			return
		}
	}
	*value = defaultValue
}

func (c *FlagContainer) Print() {
	if len(c.flags) == 0 {
		return
	}

	for i := range c.flags {
		c.flags[i].Print()
	}
}
